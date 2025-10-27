package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"text/template"
)

const NUM_SERVICES = 500

//go:embed templates/wiring.go.template
var wiringTemplate string
//go:embed templates/service_entry_with_next.go.template
var worflowServiceEntryWithNextTemplate string
//go:embed templates/service_with_next.go.template
var worflowServiceWithNextTemplate string
//go:embed templates/service_terminal.go.template
var workflowServiceTerminalTemplate string

func main() {
	GenWorkflowV1()
	GenWiringV1()
}


type ServiceSpec struct {
	N       int
	Next    int
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

	// ordered list: terminal N, middles N-1..2, entry 1
	var services []ServiceSpec

	// terminal
	services = append(services, ServiceSpec{
		N:       NUM_SERVICES,
		HasNext: false,
		HTTP:    false,
		Comment: "(terminal)",
	})

	// middles: N-1..2 (each depends on Next = i+1)
	for i := NUM_SERVICES - 1; i >= 2; i-- {
		services = append(services, ServiceSpec{
			N:       i,
			Next:    i + 1,
			HasNext: true,
			HTTP:    false,
			Comment: fmt.Sprintf("depends on service%d", i+1),
		})
	}

	// entry (1) — depends on 2 and is HTTP
	services = append(services, ServiceSpec{
		N:       1,
		Next:    2,
		HasNext: true,
		HTTP:    true,
		Comment: "(entry; depends on service2)",
	})

	data := DockerSpecData{
		ServiceCount:      NUM_SERVICES,
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

func GenWorkflowV1() {
	outputDir   := "../../workflow/large_scale_app"

	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		panic(fmt.Errorf("failed to create directory %s: %w", outputDir, err))
	}

	for i := 1; i <= NUM_SERVICES; i++ {
		filename := filepath.Join(outputDir, fmt.Sprintf("service%d.go", i))
		code, err := GenWorkflowHelperV1(i)
		if err != nil {
			panic(err)
		}
		if err := os.WriteFile(filename, []byte(code), 0o644); err != nil {
			panic(err)
		}
	}
	fmt.Printf("Generated %d files: service1.go ... service%d.go\n", NUM_SERVICES, NUM_SERVICES)
}


type svcData struct {
	N    int
	Next int
}


func GenWorkflowHelperV1(n int) (string, error) {
	serviceEntryWithNext := template.Must(template.New("entry_with_next").Parse(worflowServiceEntryWithNextTemplate))
	serviceWithNext := template.Must(template.New("with_next").Parse(worflowServiceWithNextTemplate))
	serviceTerminal := template.Must(template.New("terminal").Parse(workflowServiceTerminalTemplate))

	buf := &bytes.Buffer{}
	data := svcData{N: n, Next: n + 1}
	var err error
	if n == 1 {
		err = serviceEntryWithNext.Execute(buf, data)
	} else if n < NUM_SERVICES {
		err = serviceWithNext.Execute(buf, data)
	} else {
		err = serviceTerminal.Execute(buf, data)
	}
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
