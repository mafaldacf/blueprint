package specs

import (
	"github.com/blueprint-uservices/blueprint/blueprint/pkg/wiring"
	"github.com/blueprint-uservices/blueprint/examples/digota/workflow/digota"
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

	skus_db := mongodb.Container(spec, "skus_db")
	orders_db := mongodb.Container(spec, "orders_db")
	payments_db := mongodb.Container(spec, "payments_db")
	products_db := mongodb.Container(spec, "products_db")
	allServices = append(allServices, skus_db)
	allServices = append(allServices, orders_db)
	allServices = append(allServices, payments_db)
	allServices = append(allServices, products_db)

	sku_service := workflow.Service[digota.SkuService](spec, "sku_service", skus_db)
	sku_service_ctr := applyDockerDefaults(spec, sku_service, "sku_service_proc", "sku_service_container")
	containers = append(containers, sku_service_ctr)
	allServices = append(allServices, "sku_service")

	order_service := workflow.Service[digota.OrderService](spec, "order_service", sku_service, orders_db)
	order_service_ctr := applyHTTPDefaults(spec, order_service, "order_service_proc", "order_service_container")
	containers = append(containers, order_service_ctr)
	allServices = append(allServices, "order_service")

	payment_service := workflow.Service[digota.PaymentService](spec, "payment_service", payments_db)
	payment_service_ctr := applyDockerDefaults(spec, payment_service, "payment_service_proc", "payment_service_container")
	containers = append(containers, payment_service_ctr)
	allServices = append(allServices, "payment_service")

	product_service := workflow.Service[digota.ProductService](spec, "product_service", products_db)
	product_service_ctr := applyHTTPDefaults(spec, product_service, "product_service_proc", "product_service_container")
	containers = append(containers, product_service_ctr)
	allServices = append(allServices, "product_service")

	tests := gotests.Test(spec, allServices...)
	containers = append(containers, tests)

	return containers, nil
}
