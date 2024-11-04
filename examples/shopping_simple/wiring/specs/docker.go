package specs

import (
	"github.com/blueprint-uservices/blueprint/blueprint/pkg/wiring"
	"github.com/blueprint-uservices/blueprint/examples/shopping_simple/workflow/shopping_simple"
	"github.com/blueprint-uservices/blueprint/plugins/cmdbuilder"
	"github.com/blueprint-uservices/blueprint/plugins/gotests"
	"github.com/blueprint-uservices/blueprint/plugins/mongodb"
	"github.com/blueprint-uservices/blueprint/plugins/rabbitmq"
	"github.com/blueprint-uservices/blueprint/plugins/workflow"
	"github.com/blueprint-uservices/blueprint/plugins/goproc"
	"github.com/blueprint-uservices/blueprint/plugins/http"
	"github.com/blueprint-uservices/blueprint/plugins/linuxcontainer"
	"github.com/blueprint-uservices/blueprint/plugins/thrift"
)

var Docker = cmdbuilder.SpecOption{
	Name:        "docker",
	Description: "Deploys each service in a separate container with thrift, and uses mongodb as NoSQL database backends.",
	Build:       makeDockerSpec,
}

func applyDockerDefaults(spec wiring.WiringSpec, serviceName, procName, ctrName string) string {
	thrift.Deploy(spec, serviceName)
	goproc.CreateProcess(spec, procName, serviceName)
	return linuxcontainer.CreateContainer(spec, ctrName, procName)
}

func applyHTTPDefaults(spec wiring.WiringSpec, serviceName, procName, ctrName string) string {
	http.Deploy(spec, serviceName)
	goproc.CreateProcess(spec, procName, serviceName)
	return linuxcontainer.CreateContainer(spec, ctrName, procName)
}

// Create a basic social network wiring spec.
// Returns the names of the nodes to instantiate or an error.
func makeDockerSpec(spec wiring.WiringSpec) ([]string, error) {
	var containers []string
	var allServices []string

	cart_db := mongodb.Container(spec, "cart_db")
	product_db := mongodb.Container(spec, "product_db")
	product_queue := rabbitmq.Container(spec, "product_queue", "product_queue")

	allServices = append(allServices, cart_db)
	allServices = append(allServices, product_db)
	allServices = append(allServices, product_queue)

	product_service := workflow.Service[shopping_simple.ProductService](spec, "product_service", product_db, product_queue)
	product_service_ctr := applyDockerDefaults(spec, product_service, "product_service_proc", "product_service_container")
	containers = append(containers, product_service_ctr)
	allServices = append(allServices, "product_service")

	cart_service := workflow.Service[shopping_simple.CartService](spec, "cart_service", product_service, cart_db, product_queue)
	cart_service_ctr := applyDockerDefaults(spec, cart_service, "cart_service_proc", "cart_service_container")
	containers = append(containers, cart_service_ctr)
	allServices = append(allServices, "cart_service")

	frontend := workflow.Service[shopping_simple.Frontend](spec, "frontend", product_service, cart_service)
	frontend_ctr := applyHTTPDefaults(spec, frontend, "frontend_proc", "frontend_container")
	containers = append(containers, frontend_ctr)
	allServices = append(allServices, "frontend")

	tests := gotests.Test(spec, allServices...)
	containers = append(containers, tests)

	return containers, nil
}
