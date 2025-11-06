package specs

import (
	"github.com/blueprint-uservices/blueprint/blueprint/pkg/wiring"
	"github.com/blueprint-uservices/blueprint/examples/foobar/workflow/foobar"
	"github.com/blueprint-uservices/blueprint/plugins/cmdbuilder"
	"github.com/blueprint-uservices/blueprint/plugins/gotests"
	"github.com/blueprint-uservices/blueprint/plugins/mongodb"
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

	foo_db := mongodb.Container(spec, "foo_db")
	bar_db := mongodb.Container(spec, "bar_db")
	allServices = append(allServices, foo_db)
	allServices = append(allServices, bar_db)

	bar_service := workflow.Service[foobar.BarService](spec, "bar_service", bar_db)
	bar_service_ctr := applyDockerDefaults(spec, bar_service, "bar_service_proc", "bar_service_container")
	containers = append(containers, bar_service_ctr)
	allServices = append(allServices, "bar_service")

	foo_service := workflow.Service[foobar.FooService](spec, "foo_service", bar_service, foo_db)
	foo_service_ctr := applyDockerDefaults(spec, foo_service, "foo_service_proc", "foo_service_container")
	containers = append(containers, foo_service_ctr)
	allServices = append(allServices, "foo_service")

	frontend_service := workflow.Service[foobar.Frontend](spec, "frontend_service", foo_service, bar_service)
	frontend_service_ctr := applyHTTPDefaults(spec, frontend_service, "frontend_service_proc", "frontend_service_container")
	containers = append(containers, frontend_service_ctr)
	allServices = append(allServices, "frontend_service")

	tests := gotests.Test(spec, allServices...)
	containers = append(containers, tests)

	return containers, nil
}
