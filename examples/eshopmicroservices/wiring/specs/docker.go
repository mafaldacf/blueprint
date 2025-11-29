package specs

import (
	"github.com/blueprint-uservices/blueprint/blueprint/pkg/wiring"
	"github.com/blueprint-uservices/blueprint/examples/eshopmicroservices/workflow/basket"
	"github.com/blueprint-uservices/blueprint/examples/eshopmicroservices/workflow/catalog"
	"github.com/blueprint-uservices/blueprint/examples/eshopmicroservices/workflow/discount"
	"github.com/blueprint-uservices/blueprint/examples/eshopmicroservices/workflow/order"
	"github.com/blueprint-uservices/blueprint/examples/eshopmicroservices/workflow/web"
	"github.com/blueprint-uservices/blueprint/plugins/cmdbuilder"
	"github.com/blueprint-uservices/blueprint/plugins/gotests"
	"github.com/blueprint-uservices/blueprint/plugins/mongodb"
	"github.com/blueprint-uservices/blueprint/plugins/rabbitmq"
	"github.com/blueprint-uservices/blueprint/plugins/workflow"
)

var Docker = cmdbuilder.SpecOption{
	Name:  "docker",
	Build: makeDockerSpec,
}

func makeDockerSpec(spec wiring.WiringSpec) ([]string, error) {
	var containers []string
	var allServices []string

	discount_db := mongodb.Container(spec, "discount_db")
	allServices = append(allServices, discount_db)
	catalog_db := mongodb.Container(spec, "catalog_db")
	allServices = append(allServices, catalog_db)
	basket_db := mongodb.Container(spec, "basket_db")
	allServices = append(allServices, basket_db)
	order_db := mongodb.Container(spec, "order_db")
	allServices = append(allServices, order_db)
	order_queue := rabbitmq.Container(spec, "order_queue", "order_queue")
	allServices = append(allServices, order_queue)

	catalog_service := workflow.Service[catalog.CatalogService](spec, "catalog_service", catalog_db)
	catalog_service_ctr := applyHTTPDefaults(spec, catalog_service, "catalog_service_proc", "catalog_service_container")
	containers = append(containers, catalog_service_ctr)
	allServices = append(allServices, "catalog_service")

	discount_service := workflow.Service[discount.DiscountService](spec, "discount_service", discount_db)
	discount_service_ctr := applyHTTPDefaults(spec, discount_service, "discount_service_proc", "discount_service_container")
	containers = append(containers, discount_service_ctr)
	allServices = append(allServices, "discount_service")

	basket_service := workflow.Service[basket.BasketService](spec, "basket_service", basket_db, order_queue, discount_service)
	basket_service_ctr := applyHTTPDefaults(spec, basket_service, "basket_service_proc", "basket_service_container")
	containers = append(containers, basket_service_ctr)
	allServices = append(allServices, "basket_service")

	order_service := workflow.Service[order.OrderService](spec, "order_service", order_db, order_queue)
	order_service_ctr := applyHTTPDefaults(spec, order_service, "order_service_proc", "order_service_container")
	containers = append(containers, order_service_ctr)
	allServices = append(allServices, "order_service")

	web_app := workflow.Service[web.WebApp](spec, "web_app", basket_service, catalog_service, discount_service, order_service)
	web_app_ctr := applyHTTPDefaults(spec, web_app, "web_app_proc", "web_app_container")
	containers = append(containers, web_app_ctr)
	allServices = append(allServices, "web_app")

	tests := gotests.Test(spec, allServices...)
	containers = append(containers, tests)

	return containers, nil
}
