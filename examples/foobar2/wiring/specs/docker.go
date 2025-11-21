package specs

import (
	"github.com/blueprint-uservices/blueprint/blueprint/pkg/wiring"
	"github.com/blueprint-uservices/blueprint/examples/foobar2/workflow/foobar2"
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

	movie_db := mongodb.Container(spec, "movie_db")
	plot_db := mongodb.Container(spec, "plot_db")
	price_db := mongodb.Container(spec, "price_db")
	route_db := mongodb.Container(spec, "route_db")
	allServices = append(allServices, movie_db)
	allServices = append(allServices, plot_db)
	allServices = append(allServices, price_db)
	allServices = append(allServices, route_db)

	movie_service := workflow.Service[foobar2.MovieService](spec, "movie_service", movie_db)
	movie_service_ctr := applyDockerDefaults(spec, movie_service, "movie_service_proc", "movie_service_container")
	containers = append(containers, movie_service_ctr)
	allServices = append(allServices, "movie_service")

	plot_service := workflow.Service[foobar2.PlotService](spec, "plot_service", plot_db)
	plot_service_ctr := applyDockerDefaults(spec, plot_service, "plot_service_proc", "plot_service_container")
	containers = append(containers, plot_service_ctr)
	allServices = append(allServices, "plot_service")

	price_service := workflow.Service[foobar2.PriceService](spec, "price_service", price_db)
	price_service_ctr := applyDockerDefaults(spec, price_service, "price_service_proc", "price_service_container")
	containers = append(containers, price_service_ctr)
	allServices = append(allServices, "price_service")

	route_service := workflow.Service[foobar2.RouteService](spec, "route_service", route_db)
	route_service_ctr := applyDockerDefaults(spec, route_service, "route_service_proc", "route_service_container")
	containers = append(containers, route_service_ctr)
	allServices = append(allServices, "route_service")

	frontend_service := workflow.Service[foobar2.Frontend](spec, "frontend_service", movie_service, plot_service, price_service, route_service)
	frontend_service_ctr := applyHTTPDefaults(spec, frontend_service, "frontend_service_proc", "frontend_service_container")
	containers = append(containers, frontend_service_ctr)
	allServices = append(allServices, "frontend_service")

	tests := gotests.Test(spec, allServices...)
	containers = append(containers, tests)

	return containers, nil
}
