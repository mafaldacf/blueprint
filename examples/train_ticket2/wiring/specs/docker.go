package specs

import (
	"github.com/blueprint-uservices/blueprint/blueprint/pkg/wiring"
	"github.com/blueprint-uservices/blueprint/examples/train_ticket2/workflow/train_ticket2"
	"github.com/blueprint-uservices/blueprint/plugins/cmdbuilder"
	"github.com/blueprint-uservices/blueprint/plugins/gotests"
	"github.com/blueprint-uservices/blueprint/plugins/mongodb"
	"github.com/blueprint-uservices/blueprint/plugins/rabbitmq"
	"github.com/blueprint-uservices/blueprint/plugins/workflow"
)

// A wiring spec that deploys each service into its own Docker container and uses http to communicate between services.
// The user service uses MongoDB instance to store their data.
var Docker = cmdbuilder.SpecOption{
	Name:        "docker",
	Description: "Deploys each service in a separate container with http, and uses mongodb as NoSQL database backends",
	Build:       makeDockerSpec,
}

// Create a basic train ticket wiring spec.
// Returns the names of the nodes to instantiate or an error.
func makeDockerSpec(spec wiring.WiringSpec) ([]string, error) {
	var containers []string
	var allServices []string

	user_db := mongodb.Container(spec, "user_db")
	allServices = append(allServices, user_db)
	user_service := workflow.Service[train_ticket2.UserService](spec, "user_service", user_db)
	user_service_ctr := applyHTTPDefaults(spec, user_service, "user_proc", "user_container")
	containers = append(containers, user_service_ctr)
	allServices = append(allServices, user_service)

	contacts_db := mongodb.Container(spec, "contacts_db")
	allServices = append(allServices, contacts_db)
	contacts_service := workflow.Service[train_ticket2.ContactsService](spec, "contacts_service", contacts_db)
	contacts_service_ctr := applyHTTPDefaults(spec, contacts_service, "contacts_proc", "contacts_container")
	containers = append(containers, contacts_service_ctr)
	allServices = append(allServices, contacts_service)

	price_db := mongodb.Container(spec, "price_db")
	allServices = append(allServices, price_db)
	price_service := workflow.Service[train_ticket2.PriceService](spec, "price_service", price_db)
	price_service_ctr := applyHTTPDefaults(spec, price_service, "price_proc", "price_container")
	containers = append(containers, price_service_ctr)
	allServices = append(allServices, price_service)

	station_db := mongodb.Container(spec, "station_db")
	allServices = append(allServices, station_db)
	station_service := workflow.Service[train_ticket2.StationService](spec, "station_service", station_db)
	station_service_ctr := applyHTTPDefaults(spec, station_service, "station_proc", "station_container")
	containers = append(containers, station_service_ctr)
	allServices = append(allServices, station_service)

	news_service := workflow.Service[train_ticket2.NewsService](spec, "news_service")
	news_service_ctr := applyHTTPDefaults(spec, news_service, "news_proc", "news_container")
	containers = append(containers, news_service_ctr)
	allServices = append(allServices, news_service)

	assurance_db := mongodb.Container(spec, "assurance_db")
	allServices = append(allServices, assurance_db)
	assurance_service := workflow.Service[train_ticket2.AssuranceService](spec, "assurance_service", assurance_db)
	assurance_service_ctr := applyHTTPDefaults(spec, assurance_service, "assurance_proc", "assurance_container")
	containers = append(containers, assurance_service_ctr)
	allServices = append(allServices, assurance_service)

	config_db := mongodb.Container(spec, "config_db")
	allServices = append(allServices, config_db)
	config_service := workflow.Service[train_ticket2.ConfigService](spec, "config_service", config_db)
	config_service_ctr := applyHTTPDefaults(spec, config_service, "config_proc", "config_container")
	containers = append(containers, config_service_ctr)
	allServices = append(allServices, config_service)

	consignprice_db := mongodb.Container(spec, "consignprice_db")
	allServices = append(allServices, consignprice_db)
	consignprice_service := workflow.Service[train_ticket2.ConsignPriceService](spec, "consignprice_service", consignprice_db)
	consignprice_service_ctr := applyHTTPDefaults(spec, consignprice_service, "consignprice_proc", "consignprice_container")
	containers = append(containers, consignprice_service_ctr)
	allServices = append(allServices, consignprice_service)

	trainfood_db := mongodb.Container(spec, "trainfood_db")
	allServices = append(allServices, trainfood_db)
	trainfood_service := workflow.Service[train_ticket2.TrainFoodService](spec, "trainfood_service", trainfood_db)
	trainfood_service_ctr := applyHTTPDefaults(spec, trainfood_service, "trainfood_proc", "trainfood_container")
	containers = append(containers, trainfood_service_ctr)
	allServices = append(allServices, trainfood_service)

	train_db := mongodb.Container(spec, "train_db")
	allServices = append(allServices, train_db)
	train_service := workflow.Service[train_ticket2.TrainService](spec, "train_service", train_db)
	train_service_ctr := applyHTTPDefaults(spec, train_service, "train_proc", "train_container")
	containers = append(containers, train_service_ctr)
	allServices = append(allServices, train_service)

	route_db := mongodb.Container(spec, "route_db")
	allServices = append(allServices, route_db)
	route_service := workflow.Service[train_ticket2.RouteService](spec, "route_service", route_db)
	route_service_ctr := applyHTTPDefaults(spec, route_service, "route_proc", "route_container")
	containers = append(containers, route_service_ctr)
	allServices = append(allServices, route_service)

	stationfood_db := mongodb.Container(spec, "stationfood_db")
	allServices = append(allServices, stationfood_db)
	stationfood_service := workflow.Service[train_ticket2.StationFoodService](spec, "stationfood_service", stationfood_db)
	stationfood_service_ctr := applyHTTPDefaults(spec, stationfood_service, "stationfood_proc", "stationfood_container")
	containers = append(containers, stationfood_service_ctr)
	allServices = append(allServices, stationfood_service)

	delivery_queue := rabbitmq.Container(spec, "delivery_q", "delivery_q")
	delivery_db := mongodb.Container(spec, "delivery_db")
	allServices = append(allServices, delivery_queue)
	allServices = append(allServices, delivery_db)
	delivery_service := workflow.Service[train_ticket2.DeliveryService](spec, "delivery_service", delivery_queue, delivery_db)
	delivery_service_ctr := applyDockerQueueHandlerDefaults(spec, delivery_service, "delivery_service_proc", "delivery_service_container")
	containers = append(containers, delivery_service_ctr)

	payment_db := mongodb.Container(spec, "payment_db")
	money_db := mongodb.Container(spec, "money_db")
	allServices = append(allServices, payment_db)
	allServices = append(allServices, money_db)
	payment_service := workflow.Service[train_ticket2.PaymentService](spec, "payment_service", payment_db, money_db)
	payment_service_ctr := applyHTTPDefaults(spec, payment_service, "payment_proc", "payment_container")
	containers = append(containers, payment_service_ctr)
	allServices = append(allServices, payment_service)
	
	tests := gotests.Test(spec, allServices...)
	containers = append(containers, tests)

	return containers, nil
}
