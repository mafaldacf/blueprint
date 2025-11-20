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
)

const (
	CALL_DEPTH = 4
	OUT_DEGREE = 4 // aka fanout
)

// Example:
// CALL_DEPTH = 4, OUT_DEGREE = 3 => NUM_SERVICES = 1 + 3 + 9 + 27 + 81 = 121
// CALL_DEPTH = 4, OUT_DEGREE = 4 => NUM_SERVICES = 1 + 4 + 16 + 64 + 256 = 341
var NUM_SERVICES int

//go:embed templates/main.go.template
var mainTemplate string

//go:embed templates/wiring.go.template
var wiringTemplate string

//go:embed templates/service_entry_with_next.go.template
var worflowServiceEntryWithNextTemplate string

//go:embed templates/service_with_next.go.template
var worflowServiceWithNextTemplate string

//go:embed templates/service_terminal.go.template
var workflowServiceTerminalTemplate string

func main() {
	NUM_SERVICES = computeNumberOfServices()
	GenWorkflowV1()
	GenWiringV1()
}

func computeNumberOfServices() int {
	total := 0
	pow := 1
	for i := 0; i <= CALL_DEPTH; i++ {
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
	levels := make([][]int, CALL_DEPTH+1)

	nextID := 1
	// root
	levels[0] = []int{nextID}
	nextID++

	// allocate IDs for each subsequent level
	for d := 1; d <= CALL_DEPTH; d++ {
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
			}
		}
	}

	// assign Next for all nodes except leaves (depth = MaxDepth)
	for d := 0; d < CALL_DEPTH; d++ {
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

// -------------------------------------------------------------------
// Wiring generation
// -------------------------------------------------------------------

type ServiceSpec struct {
	N       int
	Next    []int
	HasNext bool
	HTTP    bool
	Comment string
}

type DockerSpecData struct {
	ServiceCount      int
	ServicesPkgImport string
	Services          []ServiceSpec
}

func GenWiringV1() {
	outputDir := "../../wiring/specs"
	outputFile := "docker.go"
	servicesPkgImport := "github.com/blueprint-uservices/blueprint/examples/large_scale_app/workflow/large_scale_app"

	if err := os.MkdirAll(outputDir, 0o755); err != nil {
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
		ServicesPkgImport: servicesPkgImport,
		Services:          services,
	}

	tmpl := template.Must(template.New("docker").Parse(wiringTemplate))

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		panic(err)
	}

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		panic(err)
	}

	path := filepath.Join(outputDir, outputFile)
	if err := os.WriteFile(path, formatted, 0o644); err != nil {
		panic(err)
	}

	fmt.Printf("Generated %s\n", path)
}

// -------------------------------------------------------------------
// Workflow/service files generation
// -------------------------------------------------------------------

type serviceData struct {
	N    int
	Next []int
}

type mainData struct {
	N    int
}

func GenWorkflowV1() {
	outputDir := "../../workflow/large_scale_app"

	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		panic(fmt.Errorf("failed to create directory %s: %w", outputDir, err))
	}

	graph := generateCallGraph()
	filename_main := filepath.Join(outputDir, "main.go")

	for _, g := range graph {
		filename := filepath.Join(outputDir, fmt.Sprintf("service%d.go", g.ID))
		data := serviceData{N: g.ID, Next: g.Next}

		code, err := GenWorkflowHelperV1(data)
		if err != nil {
			panic(err)
		}
		if err := os.WriteFile(filename, []byte(code), 0o644); err != nil {
			panic(err)
		}
	}

	g := graph[len(graph)-1]
	data := mainData{N: g.ID}
	code, err := GenMain(data)
	if err != nil {
		panic(err)
	}
	if err := os.WriteFile(filename_main, []byte(code), 0o644); err != nil {
		panic(err)
	}

	fmt.Printf("Generated %d files: service1.go ... service%d.go\n", len(graph), NUM_SERVICES)
}

func GenWorkflowHelperV1(data serviceData) (string, error) {
	serviceEntryWithNext := template.Must(template.New("entry_with_next").Parse(worflowServiceEntryWithNextTemplate))
	serviceWithNext := template.Must(template.New("with_next").Parse(worflowServiceWithNextTemplate))
	serviceTerminal := template.Must(template.New("terminal").Parse(workflowServiceTerminalTemplate))

	buf := &bytes.Buffer{}
	var err error

	switch {
	case data.N == 1:
		err = serviceEntryWithNext.Execute(buf, data)
	case len(data.Next) > 0:
		err = serviceWithNext.Execute(buf, data)
	default:
		err = serviceTerminal.Execute(buf, data)
	}

	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func GenMain(data mainData) (string, error) {
	main := template.Must(template.New("main").Parse(mainTemplate))
	buf := &bytes.Buffer{}
	err := main.Execute(buf, data)

	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
