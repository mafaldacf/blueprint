package specs

import (
	"github.com/blueprint-uservices/blueprint/blueprint/pkg/wiring"
	"github.com/blueprint-uservices/blueprint/examples/threechain/workflow/threechain"
	"github.com/blueprint-uservices/blueprint/plugins/cmdbuilder"
	"github.com/blueprint-uservices/blueprint/plugins/gotests"
	"github.com/blueprint-uservices/blueprint/plugins/mongodb"
	"github.com/blueprint-uservices/blueprint/plugins/rabbitmq"
	"github.com/blueprint-uservices/blueprint/plugins/redis"
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

	post_db := mongodb.Container(spec, "post_nosql")
	post_cache := redis.Container(spec, "post_cache")
	notification_queue := rabbitmq.Container(spec, "notif_queue", "notification")

	allServices = append(allServices, post_db)
	allServices = append(allServices, post_cache)
	allServices = append(allServices, notification_queue)

	order_service := workflow.Service[threechain.OrderService](spec, "order_service", post_cache, post_db)
	order_service_ctr := applyDockerDefaults(spec, order_service, "order_service_proc", "order_service_container")
	containers = append(containers, order_service_ctr)
	allServices = append(allServices, "order_service")

	stock_service := workflow.Service[threechain.StockService](spec, "stock_service", order_service, notification_queue)
	stock_service_ctr := applyDockerQueueHandlerDefaults(spec, stock_service, "stock_service_proc", "stock_service_container")
	containers = append(containers, stock_service_ctr)
	/* allServices = append(allServices, "stock_service") */

	cart_service := workflow.Service[threechain.CartService](spec, "cart_service", order_service /* stock_service, */, notification_queue)
	cart_service_ctr := applyHTTPDefaults(spec, cart_service, "cart_service_proc", "cart_service_container")
	containers = append(containers, cart_service_ctr)
	allServices = append(allServices, "cart_service")

	tests := gotests.Test(spec, allServices...)
	containers = append(containers, tests)

	return containers, nil
}
