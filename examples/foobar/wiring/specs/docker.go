package specs

import (
	"github.com/blueprint-uservices/blueprint/examples/foobar/workflow/foobar"
	"github.com/blueprint-uservices/blueprint/examples/foobar/workflow/foobar/bar"
	"github.com/blueprint-uservices/blueprint/examples/foobar/workflow/foobar/foo"

	"github.com/blueprint-uservices/blueprint/blueprint/pkg/wiring"
	"github.com/blueprint-uservices/blueprint/plugins/cmdbuilder"
	"github.com/blueprint-uservices/blueprint/plugins/gotests"
	"github.com/blueprint-uservices/blueprint/plugins/workflow"
)

var Docker = cmdbuilder.SpecOption{
	Name:        "docker",
	Description: "Deploys each service in a separate container with thrift, and uses mongodb as NoSQL database backends.",
	Build:       makeDockerSpec,
}

// Create a basic social network wiring spec.
// Returns the names of the nodes to instantiate or an error.
func makeDockerSpec(spec wiring.WiringSpec) ([]string, error) {
	var containers []string
	var allServices []string
	
	foo_service := workflow.Service[foo.FooService](spec, "foo_service")
	foo_service_ctr := applyDockerDefaults(spec, foo_service, "foo_service_proc", "foo_service_container")
	containers = append(containers, foo_service_ctr)
	allServices = append(allServices, "foo_service")

	bar_service := workflow.Service[bar.BarService](spec, "bar_service")
	bar_service_ctr := applyDockerDefaults(spec, bar_service, "bar_service_proc", "bar_service_container")
	containers = append(containers, bar_service_ctr)
	allServices = append(allServices, "bar_service")

	frontend_service := workflow.Service[foobar.FrontendService](spec, "frontend_service", foo_service, bar_service)
	frontend_service_ctr := applyHTTPDefaults(spec, frontend_service, "frontend_service_proc", "frontend_service_container")
	containers = append(containers, frontend_service_ctr)
	allServices = append(allServices, "frontend_service")

	tests := gotests.Test(spec, allServices...)
	containers = append(containers, tests)

	return containers, nil
}
