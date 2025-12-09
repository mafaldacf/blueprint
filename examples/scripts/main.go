package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"sort"
	"text/template"

	"gopkg.in/yaml.v2"
)

type AppConfig struct {
	AppName     string `yaml:"app"`
	Description string `yaml:"description"`
	CallDepth   int    `yaml:"call_depth"`
	RpcDepth    int    `yaml:"rpc_depth"`
	Fanout   int    `yaml:"fanout"`
	CallGraphs int    	   `yaml:"call_graphs"`
	ReqsStateful  int    `yaml:"reqs_stateful"`
	NumServices int    `yaml:"num_services"` // NEW: optional override
}

var cfgs []AppConfig

var APPNAME, APP_BASE_DIR, WORKFLOW_DIR, SERVICES_PKG_IMPORT string
var RPC_DEPTH, OUT_DEGREE, NUM_SERVICES, NUM_ENTRYPOINTS, NUM_DB_ACCESSES int

//go:embed templates/.gitignore.template
var gitIgnoreTemplate string

//go:embed templates/wiring.specs.docker.go.template
var wiringTemplate string

//go:embed templates/wiring.go.mod.template
var wiringGoModTemplate string

//go:embed templates/workflow.go.mod.template
var workflowGoModTemplate string

//go:embed templates/workflow.app.service.go.template
var workflowServiceTemplate string

func loadConfigs(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(data, &cfgs)
	if err != nil {
		panic(err)
	}
}

func main() {
	loadConfigs("config.yaml")

	for _, cfg := range cfgs {
		fmt.Printf(
			"==== Generating app: %s (call_depth=%d, rpc_depth=%d, fanout=%d, call_graphs=%d) ====\n",
			cfg.AppName, cfg.CallDepth, cfg.RpcDepth, cfg.Fanout, cfg.CallGraphs,
		)

		APPNAME = cfg.AppName

		if cfg.RpcDepth > 0 {
			RPC_DEPTH = cfg.RpcDepth
		} else {
			RPC_DEPTH = cfg.CallDepth - 1
		}

		OUT_DEGREE = cfg.Fanout

		if cfg.CallGraphs <= 0 {
			NUM_ENTRYPOINTS = 1
		} else {
			NUM_ENTRYPOINTS = cfg.CallGraphs
		}

		APP_BASE_DIR = fmt.Sprintf("../%s", APPNAME)
		WORKFLOW_DIR = filepath.Join(APP_BASE_DIR, fmt.Sprintf("workflow/%s", APPNAME))
		SERVICES_PKG_IMPORT = fmt.Sprintf("github.com/blueprint-uservices/blueprint/examples/%s/workflow/%s", APPNAME, APPNAME)

		// Decide NUM_SERVICES: either from config or theoretical maximum
		maxPossible := computeNumberOfServices() // full k-ary tree size
		if cfg.NumServices > 0 {
			if cfg.NumServices > maxPossible {
				panic(fmt.Errorf(
					"num_services=%d exceeds maximum possible (%d) for rpc_depth=%d, fanout=%d",
					cfg.NumServices, maxPossible, RPC_DEPTH, OUT_DEGREE,
				))
			}
			NUM_SERVICES = cfg.NumServices
		} else {
			NUM_SERVICES = maxPossible
		}

		if err := os.RemoveAll(APP_BASE_DIR); err != nil {
			panic(err)
		}

		if cfg.ReqsStateful <= 0 || cfg.ReqsStateful > NUM_SERVICES {
			NUM_DB_ACCESSES = NUM_SERVICES
		} else {
			NUM_DB_ACCESSES = cfg.ReqsStateful
		}

		if err := os.MkdirAll(APP_BASE_DIR, 0o755); err != nil {
			panic(err)
		}

		GenBaseDirFiles()
		GenWorkflow()
		GenWiring()
	}
}

// computeNumberOfServices returns the full k-ary tree size for the given
// RPC_DEPTH and OUT_DEGREE, i.e., the theoretical maximum number of services.
func computeNumberOfServices() int {
	total := 0
	pow := 1
	for i := 0; i <= RPC_DEPTH; i++ {
		total += pow
		pow *= OUT_DEGREE
	}
	return total
}

// -------------------------------------------------------------------
// Call graph generation
// -------------------------------------------------------------------

type GraphService struct {
	ID    int
	Next  []int
	Depth int
	HasDB bool
}

// generateCallGraph builds a k-ary tree up to RPC_DEPTH with fan-out OUT_DEGREE,
// but truncated to exactly NUM_SERVICES nodes. It guarantees IDs 1..NUM_SERVICES.
func generateCallGraph() []GraphService {
	if NUM_SERVICES <= 0 {
		panic("NUM_SERVICES must be > 0")
	}

	// Initialize services 1..NUM_SERVICES
	services := make([]GraphService, NUM_SERVICES+1) // 1-based
	for i := 1; i <= NUM_SERVICES; i++ {
		services[i] = GraphService{
			ID:    i,
			Next:  nil,
			Depth: -1, // uninitialized
			HasDB: false,
		}
	}

	// BFS-style construction of a k-ary tree up to RPC_DEPTH
	type qItem struct {
		id    int
		depth int
	}

	// Root
	services[1].Depth = 0
	queue := []qItem{{id: 1, depth: 0}}
	nextID := 2

	for len(queue) > 0 && nextID <= NUM_SERVICES {
		item := queue[0]
		queue = queue[1:]

		if item.depth >= RPC_DEPTH {
			// do not add children beyond RPC_DEPTH
			continue
		}

		parent := services[item.id]

		for c := 0; c < OUT_DEGREE && nextID <= NUM_SERVICES; c++ {
			childID := nextID
			nextID++

			parent.Next = append(parent.Next, childID)

			child := services[childID]
			child.Depth = item.depth + 1
			services[childID] = child

			queue = append(queue, qItem{id: childID, depth: item.depth + 1})
		}

		services[item.id] = parent
	}

	// Compute levels from Depth to assign DBs
	maxDepth := 0
	for i := 1; i <= NUM_SERVICES; i++ {
		if services[i].Depth < 0 {
			// In practice this should not happen if NUM_SERVICES <= maxPossible,
			// but guard anyway: treat as depth 0.
			services[i].Depth = 0
		}
		if services[i].Depth > maxDepth {
			maxDepth = services[i].Depth
		}
	}

	levels := make([][]int, maxDepth+1)
	for i := 1; i <= NUM_SERVICES; i++ {
		d := services[i].Depth
		levels[d] = append(levels[d], i)
	}

	// ----------------- DB assignment -----------------

	remaining := NUM_DB_ACCESSES

	// Root (service 1) must always have a DB
	if remaining > 0 {
		root := services[1]
		if !root.HasDB {
			root.HasDB = true
			services[1] = root
			remaining--
		}
	}

	// Assign DB accesses deepest-first until remaining == 0
	for d := maxDepth; d >= 0 && remaining > 0; d-- {
		for _, id := range levels[d] {
			if remaining == 0 {
				break
			}
			svc := services[id]
			if !svc.HasDB {
				svc.HasDB = true
				services[id] = svc
				remaining--
			}
		}
	}

	// Flatten 1..NUM_SERVICES
	res := make([]GraphService, 0, NUM_SERVICES)
	for id := 1; id <= NUM_SERVICES; id++ {
		res = append(res, services[id])
	}

	if len(res) != NUM_SERVICES {
		panic(fmt.Sprintf("NUM_SERVICES mismatch: expected=%d, generated=%d", NUM_SERVICES, len(res)))
	}

	return res
}

func GenBaseDirFiles() {
	// generate .gitignore
	tmpl := template.Must(template.New("gitignore").Parse(gitIgnoreTemplate))
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, nil); err != nil {
		panic(err)
	}
	path := filepath.Join(APP_BASE_DIR, ".gitignore")
	if err := os.WriteFile(path, buf.Bytes(), 0o644); err != nil {
		panic(err)
	}
	fmt.Printf("Generated %s\n", path)
}

// -------------------------------------------------------------------
// Wiring generation
// -------------------------------------------------------------------

type ServiceSpec struct {
	N       int
	Next    []int
	HasNext bool
	HTTP    bool
	HasDB   bool
	Comment string
	PkgName string
}

type DockerSpecData struct {
	ServiceCount      int
	ServicesPkgImport string
	Services          []ServiceSpec
	PkgName           string
}

func GenWiring() {
	if err := os.MkdirAll(filepath.Join(APP_BASE_DIR, "wiring"), 0o755); err != nil {
		panic(err)
	}
	if err := os.MkdirAll(filepath.Join(APP_BASE_DIR, "wiring/specs"), 0o755); err != nil {
		panic(err)
	}

	graph := generateCallGraph()

	// Sort by depth DESC so leaves come first (helpful if your template
	// assumes children are constructed before parents).
	sort.Slice(graph, func(i, j int) bool {
		return graph[i].Depth > graph[j].Depth
	})

	var services []ServiceSpec
	for _, g := range graph {
		spec := ServiceSpec{
			N:       g.ID,
			Next:    g.Next,
			HasNext: len(g.Next) > 0,
			HTTP:    g.ID == 1, // service 1 is entry / HTTP
			HasDB:   g.HasDB,
			PkgName: APPNAME,
		}

		if len(g.Next) == 0 {
			spec.Comment = "(terminal)"
		} else if g.ID == 1 {
			spec.Comment = "(entry)"
		} else {
			spec.Comment = fmt.Sprintf("depends on %v", g.Next)
		}

		services = append(services, spec)
	}

	data := DockerSpecData{
		ServiceCount:      len(services),
		ServicesPkgImport: SERVICES_PKG_IMPORT,
		Services:          services,
		PkgName:           APPNAME,
	}

	// generate wiring.go
	tmpl := template.Must(template.New("docker").Parse(wiringTemplate))
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		panic(err)
	}
	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		panic(err)
	}
	path := filepath.Join(APP_BASE_DIR, "wiring/specs/docker.go")
	if err := os.WriteFile(path, formatted, 0o644); err != nil {
		panic(err)
	}
	fmt.Printf("Generated %s\n", path)

	// generate go.mod
	tmpl2 := template.Must(template.New("gomod").Parse(wiringGoModTemplate))
	var buf2 bytes.Buffer
	if err := tmpl2.Execute(&buf2, data); err != nil {
		panic(err)
	}
	path2 := filepath.Join(APP_BASE_DIR, "wiring/go.mod")
	if err := os.WriteFile(path2, buf2.Bytes(), 0o644); err != nil {
		panic(err)
	}
	fmt.Printf("Generated %s\n", path2)
}

// -------------------------------------------------------------------
// Workflow/service files generation
// -------------------------------------------------------------------

type serviceData struct {
	N       int
	Next    []int
	PkgName string
	Methods []int
	HasDB   bool
}

type mainData struct {
	N       int
	PkgName string
}

func makeMethods(n int) []int {
	ms := make([]int, n)
	for i := range ms {
		ms[i] = i + 1 // methods indices: 1..n
	}
	return ms
}

func GenWorkflow() {
	if err := os.MkdirAll(WORKFLOW_DIR, os.ModePerm); err != nil {
		panic(fmt.Errorf("failed to create directory %s: %w", WORKFLOW_DIR, err))
	}

	graph := generateCallGraph()

	for _, g := range graph {
		filename := filepath.Join(WORKFLOW_DIR, fmt.Sprintf("service%d.go", g.ID))
		data := serviceData{
			N:       g.ID,
			Next:    g.Next,
			PkgName: APPNAME,
			Methods: makeMethods(NUM_ENTRYPOINTS),
			HasDB:   g.HasDB,
		}

		code, err := GenWorkflowServices(data)
		if err != nil {
			panic(err)
		}
		if err := os.WriteFile(filename, []byte(code), 0o644); err != nil {
			panic(err)
		}
	}

	g := graph[len(graph)-1]
	data := mainData{N: g.ID, PkgName: APPNAME}

	fmt.Printf("Generated %d files: service1.go ... service%d.go\n", len(graph), NUM_SERVICES)

	// generate go.mod
	tmpl2 := template.Must(template.New("gomod").Parse(workflowGoModTemplate))
	var buf2 bytes.Buffer
	if err := tmpl2.Execute(&buf2, data); err != nil {
		panic(err)
	}
	path2 := filepath.Join(APP_BASE_DIR, "workflow/go.mod")
	if err := os.WriteFile(path2, buf2.Bytes(), 0o644); err != nil {
		panic(err)
	}
	fmt.Printf("Generated %s\n", path2)
}

func GenWorkflowServices(data serviceData) (string, error) {
	var workflowServiceTmpl = template.Must(
		template.New("service").Parse(workflowServiceTemplate),
	)
	buf := &bytes.Buffer{}
	if err := workflowServiceTmpl.Execute(buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}
