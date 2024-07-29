package specs

import (
	"github.com/blueprint-uservices/blueprint/blueprint/pkg/wiring"
	"github.com/blueprint-uservices/blueprint/examples/shopping_app/workflow/shopping_app"
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
	cart_db := mongodb.Container(spec, "cart_db")
	product_db := mongodb.Container(spec, "product_db")
	analytics_db := mongodb.Container(spec, "analytics_db")
	billing_db := mongodb.Container(spec, "billing_db")
	order_db := mongodb.Container(spec, "order_db")
	shipment_queue := rabbitmq.Container(spec, "shipment_queue", "shipment_queue")
	analytics_queue := rabbitmq.Container(spec, "analytics_queue", "analytics_queue")

	allServices = append(allServices, shipment_db)
	allServices = append(allServices, stock_db)
	allServices = append(allServices, cart_db)
	allServices = append(allServices, product_db)
	allServices = append(allServices, analytics_db)
	allServices = append(allServices, billing_db)
	allServices = append(allServices, order_db)
	allServices = append(allServices, shipment_queue)
	allServices = append(allServices, analytics_queue)

	stock_service := workflow.Service[shopping_app.StockService](spec, "stock_service", stock_db)
	stock_service_ctr := applyDockerDefaults(spec, stock_service, "stock_service_proc", "stock_service_container")
	containers = append(containers, stock_service_ctr)
	allServices = append(allServices, "stock_service")

	billing_service := workflow.Service[shopping_app.BillingService](spec, "billing_service", billing_db)
	billing_service_ctr := applyDockerDefaults(spec, billing_service, "billing_service_proc", "billing_service_container")
	containers = append(containers, billing_service_ctr)
	allServices = append(allServices, "billing_service")

	analytics_service := workflow.Service[shopping_app.AnalyticsService](spec, "analytics_service", analytics_db, analytics_queue)
	analytics_service_ctr := applyDockerQueueHandlerDefaults(spec, analytics_service, "analytics_service_proc", "analytics_service_container")
	containers = append(containers, analytics_service_ctr)

	product_service := workflow.Service[shopping_app.ProductService](spec, "product_service", product_db)
	product_service_ctr := applyDockerDefaults(spec, product_service, "product_service_proc", "product_service_container")
	containers = append(containers, product_service_ctr)
	allServices = append(allServices, "product_service")

	order_service := workflow.Service[shopping_app.OrderService](spec, "order_service", stock_service, billing_service, product_service, order_db, shipment_queue, analytics_queue)
	order_service_ctr := applyDockerDefaults(spec, order_service, "order_service_proc", "order_service_container")
	containers = append(containers, order_service_ctr)
	allServices = append(allServices, "order_service")

	cart_service := workflow.Service[shopping_app.CartService](spec, "cart_service", order_service, product_service, cart_db)
	cart_service_ctr := applyDockerDefaults(spec, cart_service, "cart_service_proc", "cart_service_container")
	containers = append(containers, cart_service_ctr)
	allServices = append(allServices, "cart_service")

	shipment_service := workflow.Service[shopping_app.ShipmentService](spec, "shipment_service", order_service, shipment_db, shipment_queue)
	shipment_service_ctr := applyDockerDefaults(spec, shipment_service, "shipment_service_proc", "shipment_service_container")
	containers = append(containers, shipment_service_ctr)

	frontend := workflow.Service[shopping_app.Frontend](spec, "frontend", order_service, cart_service)
	frontend_ctr := applyHTTPDefaults(spec, frontend, "frontend_proc", "frontend_container")
	containers = append(containers, frontend_ctr)
	allServices = append(allServices, "frontend")

	tests := gotests.Test(spec, allServices...)
	containers = append(containers, tests)

	return containers, nil
}
