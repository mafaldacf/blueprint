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
	OutDegree   int    `yaml:"out_degree"`
	Entrypoints int    `yaml:"entrypoints"`
	DbAccesses  int    `yaml:"db_accesses"`
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
		fmt.Printf("==== Generating app: %s (call_depth=%d, rpc_depth=%d, out_degree=%d, entrypoints=%d) ====\n", cfg.AppName, cfg.CallDepth, cfg.RpcDepth, cfg.OutDegree, cfg.Entrypoints)
		APPNAME = cfg.AppName
		if cfg.RpcDepth > 0 {
			RPC_DEPTH = cfg.RpcDepth
		} else {
			RPC_DEPTH = cfg.CallDepth - 1
		}
		OUT_DEGREE = cfg.OutDegree
		if cfg.Entrypoints <= 0 {
			NUM_ENTRYPOINTS = 1
		} else {
			NUM_ENTRYPOINTS = cfg.Entrypoints
		}
		APP_BASE_DIR = fmt.Sprintf("../%s", APPNAME)
		WORKFLOW_DIR = filepath.Join(APP_BASE_DIR, fmt.Sprintf("workflow/%s", APPNAME))
		SERVICES_PKG_IMPORT = fmt.Sprintf("github.com/blueprint-uservices/blueprint/examples/%s/workflow/%s", APPNAME, APPNAME)

		NUM_SERVICES = computeNumberOfServices()
		if err := os.RemoveAll(APP_BASE_DIR); err != nil {
			panic(err)
		}
		if cfg.DbAccesses <= 0 || cfg.DbAccesses > NUM_SERVICES {
			NUM_DB_ACCESSES = NUM_SERVICES
		} else {
			NUM_DB_ACCESSES = cfg.DbAccesses
		}
		if err := os.MkdirAll(APP_BASE_DIR, 0o755); err != nil {
			panic(err)
		}
		GenBaseDirFiles()
		GenWorkflow()
		GenWiring()
	}
}

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

func intPow(a, b int) int {
	res := 1
	for i := 0; i < b; i++ {
		res *= a
	}
	return res
}

func generateCallGraph() []GraphService {
	// levels[d] = list of IDs at depth d
	levels := make([][]int, RPC_DEPTH+1)

	nextID := 1
	// root
	levels[0] = []int{nextID}
	nextID++

	// allocate IDs for each subsequent level
	for d := 1; d <= RPC_DEPTH; d++ {
		levelSize := intPow(OUT_DEGREE, d)
		level := make([]int, levelSize)
		for i := 0; i < levelSize; i++ {
			level[i] = nextID
			nextID++
		}
		levels[d] = level
	}

	total := nextID - 1
	services := make([]GraphService, total+1) // 1-based, index 0 unused

	for d, lvl := range levels {
		for _, id := range lvl {
			services[id] = GraphService{
				ID:    id,
				Depth: d,
				Next:  nil,
				HasDB: false,
			}
		}
	}

	// assign Next for all nodes except leaves (depth = MaxDepth)
	for d := 0; d < RPC_DEPTH; d++ {
		parents := levels[d]
		children := levels[d+1]
		childIdx := 0

		for _, p := range parents {
			nexts := make([]int, 0, OUT_DEGREE)
			for i := 0; i < OUT_DEGREE; i++ {
				if childIdx >= len(children) {
					panic("not enough children allocated for given fanout")
				}
				nexts = append(nexts, children[childIdx])
				childIdx++
			}
			svc := services[p]
			svc.Next = nexts
			services[p] = svc
		}
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

	// assign DB accesses: deepest levels first (leaves → root) until remaining == 0
	for d := RPC_DEPTH; d >= 0 && remaining > 0; d-- {
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

	// flatten 1..total
	res := make([]GraphService, 0, total)
	for id := 1; id <= total; id++ {
		res = append(res, services[id])
	}

	if len(res) != NUM_SERVICES {
		panic(fmt.Sprintf("NUM_SERVICES mismatch: const=%d, generated=%d", NUM_SERVICES, len(res)))
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
