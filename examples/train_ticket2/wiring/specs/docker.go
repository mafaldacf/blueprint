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

	// ------
	// DOCKER
	// ------

	user_db := mongodb.Container(spec, "user_db")
	allServices = append(allServices, user_db)
	user_service := workflow.Service[train_ticket2.UserService](spec, "user_service", user_db)
	user_service_ctr := applyDockerDefaults(spec, user_service, "user_proc", "user_container")
	containers = append(containers, user_service_ctr)
	allServices = append(allServices, user_service)

	contacts_db := mongodb.Container(spec, "contacts_db")
	allServices = append(allServices, contacts_db)
	contacts_service := workflow.Service[train_ticket2.ContactsService](spec, "contacts_service", contacts_db)
	contacts_service_ctr := applyDockerDefaults(spec, contacts_service, "contacts_proc", "contacts_container")
	containers = append(containers, contacts_service_ctr)
	allServices = append(allServices, contacts_service)

	price_db := mongodb.Container(spec, "price_db")
	allServices = append(allServices, price_db)
	price_service := workflow.Service[train_ticket2.PriceService](spec, "price_service", price_db)
	price_service_ctr := applyDockerDefaults(spec, price_service, "price_proc", "price_container")
	containers = append(containers, price_service_ctr)
	allServices = append(allServices, price_service)

	station_db := mongodb.Container(spec, "station_db")
	allServices = append(allServices, station_db)
	station_service := workflow.Service[train_ticket2.StationService](spec, "station_service", station_db)
	station_service_ctr := applyDockerDefaults(spec, station_service, "station_proc", "station_container")
	containers = append(containers, station_service_ctr)
	allServices = append(allServices, station_service)

	news_service := workflow.Service[train_ticket2.NewsService](spec, "news_service")
	news_service_ctr := applyDockerDefaults(spec, news_service, "news_proc", "news_container")
	containers = append(containers, news_service_ctr)
	allServices = append(allServices, news_service)

	assurance_db := mongodb.Container(spec, "assurance_db")
	allServices = append(allServices, assurance_db)
	assurance_service := workflow.Service[train_ticket2.AssuranceService](spec, "assurance_service", assurance_db)
	assurance_service_ctr := applyDockerDefaults(spec, assurance_service, "assurance_proc", "assurance_container")
	containers = append(containers, assurance_service_ctr)
	allServices = append(allServices, assurance_service)

	config_db := mongodb.Container(spec, "config_db")
	allServices = append(allServices, config_db)
	config_service := workflow.Service[train_ticket2.ConfigService](spec, "config_service", config_db)
	config_service_ctr := applyDockerDefaults(spec, config_service, "config_proc", "config_container")
	containers = append(containers, config_service_ctr)
	allServices = append(allServices, config_service)

	consignprice_db := mongodb.Container(spec, "consignprice_db")
	allServices = append(allServices, consignprice_db)
	consignprice_service := workflow.Service[train_ticket2.ConsignPriceService](spec, "consignprice_service", consignprice_db)
	consignprice_service_ctr := applyDockerDefaults(spec, consignprice_service, "consignprice_proc", "consignprice_container")
	containers = append(containers, consignprice_service_ctr)
	allServices = append(allServices, consignprice_service)

	trainfood_db := mongodb.Container(spec, "trainfood_db")
	allServices = append(allServices, trainfood_db)
	trainfood_service := workflow.Service[train_ticket2.TrainFoodService](spec, "trainfood_service", trainfood_db)
	trainfood_service_ctr := applyDockerDefaults(spec, trainfood_service, "trainfood_proc", "trainfood_container")
	containers = append(containers, trainfood_service_ctr)
	allServices = append(allServices, trainfood_service)

	train_db := mongodb.Container(spec, "train_db")
	allServices = append(allServices, train_db)
	train_service := workflow.Service[train_ticket2.TrainService](spec, "train_service", train_db)
	train_service_ctr := applyDockerDefaults(spec, train_service, "train_proc", "train_container")
	containers = append(containers, train_service_ctr)
	allServices = append(allServices, train_service)

	route_db := mongodb.Container(spec, "route_db")
	allServices = append(allServices, route_db)
	route_service := workflow.Service[train_ticket2.RouteService](spec, "route_service", route_db)
	route_service_ctr := applyDockerDefaults(spec, route_service, "route_proc", "route_container")
	containers = append(containers, route_service_ctr)
	allServices = append(allServices, route_service)

	stationfood_db := mongodb.Container(spec, "stationfood_db")
	allServices = append(allServices, stationfood_db)
	stationfood_service := workflow.Service[train_ticket2.StationFoodService](spec, "stationfood_service", stationfood_db)
	stationfood_service_ctr := applyDockerDefaults(spec, stationfood_service, "stationfood_proc", "stationfood_container")
	containers = append(containers, stationfood_service_ctr)
	allServices = append(allServices, stationfood_service)

	delivery_queue := rabbitmq.Container(spec, "delivery_queue", "delivery_queue")
	delivery_db := mongodb.Container(spec, "delivery_db")
	allServices = append(allServices, delivery_queue)
	allServices = append(allServices, delivery_db)
	delivery_service := workflow.Service[train_ticket2.DeliveryService](spec, "delivery_service", delivery_queue, delivery_db)
	delivery_service_ctr := applyDockerDefaults(spec, delivery_service, "delivery_service_proc", "delivery_service_container")
	containers = append(containers, delivery_service_ctr)

	payment_db := mongodb.Container(spec, "payment_db")
	money_db := mongodb.Container(spec, "money_db")
	allServices = append(allServices, payment_db)
	allServices = append(allServices, money_db)
	payment_service := workflow.Service[train_ticket2.PaymentService](spec, "payment_service", payment_db, money_db)
	payment_service_ctr := applyDockerDefaults(spec, payment_service, "payment_proc", "payment_container")
	containers = append(containers, payment_service_ctr)
	allServices = append(allServices, payment_service)

	consign_db := mongodb.Container(spec, "consign_db")
	allServices = append(allServices, consign_db)
	consign_service := workflow.Service[train_ticket2.ConsignService](spec, "consign_price_service", consignprice_service, consign_db)
	consign_service_ctr := applyDockerDefaults(spec, consign_service, "consign_service_proc", "consign_service_container")
	containers = append(containers, consign_service_ctr)
	allServices = append(allServices, consign_service)

	order_db := mongodb.Container(spec, "order_db")
	allServices = append(allServices, order_db)
	order_service := workflow.Service[train_ticket2.OrderService](spec, "order_service", order_db)
	order_service_ctr := applyDockerDefaults(spec, order_service, "order_service_proc", "order_service_container")
	containers = append(containers, order_service_ctr)
	allServices = append(allServices, order_service)

	seat_service := workflow.Service[train_ticket2.SeatService](spec, "seat_service", order_service, config_service)
	seat_service_ctr := applyDockerDefaults(spec, seat_service, "seat_service_proc", "seat_service_container")
	containers = append(containers, seat_service_ctr)
	allServices = append(allServices, seat_service)

	food_db := mongodb.Container(spec, "food_db")
	allServices = append(allServices, food_db)
	food_service := workflow.Service[train_ticket2.FoodService](spec, "food_service", food_db, delivery_queue, trainfood_service, "travel_service.client", stationfood_service)
	food_service_ctr := applyDockerDefaults(spec, food_service, "food_service_proc", "food_service_container")
	containers = append(containers, food_service_ctr)
	allServices = append(allServices, food_service)

	inside_payment_db := mongodb.Container(spec, "inside_payment_db")
	allServices = append(allServices, inside_payment_db)
	inside_payment_service := workflow.Service[train_ticket2.InsidePaymentService](spec, "inside_payment_service", inside_payment_db, payment_service, order_service)
	inside_payment_service_ctr := applyDockerDefaults(spec, inside_payment_service, "inside_payment_service_proc", "inside_payment_service_container")
	containers = append(containers, inside_payment_service_ctr)
	allServices = append(allServices, inside_payment_service)

	execute_service := workflow.Service[train_ticket2.ExecuteService](spec, "execute_service", order_service)
	execute_service_ctr := applyDockerDefaults(spec, execute_service, "execute_service_proc", "execute_service_container")
	containers = append(containers, execute_service_ctr)
	allServices = append(allServices, execute_service)

	basic_service := workflow.Service[train_ticket2.BasicService](spec, "basic_service", station_service, train_service, route_service, price_service)
	basic_service_ctr := applyDockerDefaults(spec, basic_service, "basic_service_proc", "basic_service_container")
	containers = append(containers, basic_service_ctr)
	allServices = append(allServices, basic_service)

	travel_db := mongodb.Container(spec, "travel_db")
	allServices = append(allServices, order_db)
	travel_service := workflow.Service[train_ticket2.TravelService](spec, "travel_service", basic_service, seat_service, route_service, train_service, travel_db)
	travel_service_ctr := applyDockerDefaults(spec, travel_service, "travel_service_proc", "travel_service_container")
	containers = append(containers, travel_service_ctr)
	allServices = append(allServices, travel_service)

	email_queue := rabbitmq.Container(spec, "email_queue", "email_queue")
	allServices = append(allServices, email_queue)
	preserve_service := workflow.Service[train_ticket2.PreserveService](spec, "preserve_service",
		assurance_service, basic_service, consign_service, contacts_service, food_service, order_service,
		seat_service, station_service, travel_service, user_service, email_queue)
	preserve_service_ctr := applyDockerDefaults(spec, preserve_service, "preserve_service_proc", "preserve_service_container")
	containers = append(containers, preserve_service_ctr)
	allServices = append(allServices, preserve_service)

	cancel_service := workflow.Service[train_ticket2.CancelService](spec, "cancel_service", order_service, user_service, inside_payment_service, email_queue)
	cancel_service_ctr := applyDockerDefaults(spec, cancel_service, "cancel_service_proc", "cancel_service_container")
	containers = append(containers, cancel_service_ctr)
	allServices = append(allServices, cancel_service)

	rebook_service := workflow.Service[train_ticket2.RebookService](spec, "rebook_service", seat_service, travel_service, order_service, train_service, route_service, inside_payment_service)
	rebook_service_ctr := applyDockerDefaults(spec, rebook_service, "rebook_service_proc", "rebook_service_container")
	containers = append(containers, rebook_service_ctr)
	allServices = append(allServices, rebook_service)

	// ----
	// HTTP
	// ----

	admin_basic_info_service := workflow.Service[train_ticket2.AdminBasicInfoService](spec, "admin_basic_info_service", station_service, train_service, config_service, contacts_service, price_service)
	admin_basic_info_service_ctr := applyHTTPDefaults(spec, admin_basic_info_service, "admin_basic_info_service_proc", "admin_basic_info_service_container")
	containers = append(containers, admin_basic_info_service_ctr)
	allServices = append(allServices, admin_basic_info_service)

	admin_order_service := workflow.Service[train_ticket2.AdminOrderService](spec, "admin_order_service", order_service)
	admin_order_service_ctr := applyHTTPDefaults(spec, admin_order_service, "admin_order_service_proc", "admin_order_service_container")
	containers = append(containers, admin_order_service_ctr)
	allServices = append(allServices, admin_order_service)

	admin_route_service := workflow.Service[train_ticket2.AdminRouteService](spec, "admin_route_service", route_service)
	admin_route_service_ctr := applyHTTPDefaults(spec, admin_route_service, "admin_route_service_proc", "admin_route_service_container")
	containers = append(containers, admin_route_service_ctr)
	allServices = append(allServices, admin_route_service)

	admin_travel_service := workflow.Service[train_ticket2.AdminTravelService](spec, "admin_travel_service", travel_service)
	admin_travel_service_ctr := applyHTTPDefaults(spec, admin_travel_service, "admin_travel_service_proc", "admin_travel_service_container")
	containers = append(containers, admin_travel_service_ctr)
	allServices = append(allServices, admin_travel_service)

	admin_user_service := workflow.Service[train_ticket2.AdminUserService](spec, "admin_user_service", user_service)
	admin_user_service_ctr := applyHTTPDefaults(spec, admin_user_service, "admin_user_service_proc", "admin_user_service_container")
	containers = append(containers, admin_user_service_ctr)
	allServices = append(allServices, admin_user_service)

	dashboard := workflow.Service[train_ticket2.Dashboard](spec, "dashboard",
		basic_service, config_service, contacts_service, payment_service, preserve_service, execute_service, cancel_service, rebook_service, price_service,
		assurance_service, order_service, food_service, consign_service, delivery_service)
	dashboard_ctr := applyHTTPDefaults(spec, dashboard, "dashboard_proc", "dashboard_container")
	containers = append(containers, dashboard_ctr)
	allServices = append(allServices, dashboard)

	tests := gotests.Test(spec, allServices...)
	containers = append(containers, tests)

	return containers, nil
}
