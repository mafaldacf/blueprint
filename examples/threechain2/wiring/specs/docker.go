package specs

import (
	"github.com/blueprint-uservices/blueprint/blueprint/pkg/wiring"
	"github.com/blueprint-uservices/blueprint/examples/threechain2/workflow/threechain2"
	"github.com/blueprint-uservices/blueprint/plugins/cmdbuilder"
	"github.com/blueprint-uservices/blueprint/plugins/gotests"
	"github.com/blueprint-uservices/blueprint/plugins/mongodb"
	"github.com/blueprint-uservices/blueprint/plugins/rabbitmq"
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

	shipment_db := mongodb.Container(spec, "shipment_db")
	stock_db := mongodb.Container(spec, "stock_db")
	order_db := mongodb.Container(spec, "order_db")
	shipment_queue := rabbitmq.Container(spec, "shipment_queue", "shipment_queue")

	allServices = append(allServices, shipment_db)
	allServices = append(allServices, stock_db)
	allServices = append(allServices, order_db)
	allServices = append(allServices, shipment_queue)

	stock_service := workflow.Service[threechain2.StockService](spec, "stock_service", stock_db)
	stock_service_ctr := applyDockerDefaults(spec, stock_service, "stock_service_proc", "stock_service_container")
	containers = append(containers, stock_service_ctr)
	allServices = append(allServices, "stock_service")

	order_service := workflow.Service[threechain2.OrderService](spec, "order_service", stock_service, stock_db, shipment_queue)
	order_service_ctr := applyHTTPDefaults(spec, order_service, "order_service_proc", "order_service_container")
	containers = append(containers, order_service_ctr)
	allServices = append(allServices, "order_service")

	shipment_service := workflow.Service[threechain2.ShipmentService](spec, "shipment_service", order_service, shipment_db, shipment_queue)
	shipment_service_ctr := applyDockerQueueHandlerDefaults(spec, shipment_service, "shipment_service_proc", "shipment_service_container")
	containers = append(containers, shipment_service_ctr)

	tests := gotests.Test(spec, allServices...)
	containers = append(containers, tests)

	return containers, nil
}
