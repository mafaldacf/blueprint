package specs

import (
	"github.com/blueprint-uservices/blueprint/blueprint/pkg/wiring"
	"github.com/blueprint-uservices/blueprint/examples/dsb_media_sql/workflow/mediamicroservices_sql"
	"github.com/blueprint-uservices/blueprint/plugins/cmdbuilder"
	"github.com/blueprint-uservices/blueprint/plugins/goproc"
	"github.com/blueprint-uservices/blueprint/plugins/http"
	"github.com/blueprint-uservices/blueprint/plugins/linuxcontainer"
	"github.com/blueprint-uservices/blueprint/plugins/mysql"
	"github.com/blueprint-uservices/blueprint/plugins/thrift"
	"github.com/blueprint-uservices/blueprint/plugins/workflow"
)

var Docker = cmdbuilder.SpecOption{
	Name:        "docker",
	Description: "Deploys each service in a separate container with thrift, and uses mongodb as NoSQL database backends.",
	Build:       makeDockerSpec,
}

func makeDockerSpec(spec wiring.WiringSpec) ([]string, error) {
	var containers []string

	movieid_db := mysql.Container(spec, "movieid_db")
	movieinfo_db := mysql.Container(spec, "movieinfo_db")
	castinfo_db := mysql.Container(spec, "castinfo_db")
	plot_db := mysql.Container(spec, "plot_db")

	movieid_service := workflow.Service[mediamicroservices_sql.MovieIdService](spec, "movieid_service", movieid_db)
	movieid_ctr := applyDockerDefaults(spec, movieid_service, "movieid_proc", "movieid_container")
	containers = append(containers, movieid_ctr)

	movieinfo_service := workflow.Service[mediamicroservices_sql.MovieInfoService](spec, "movieinfo_service", movieinfo_db)
	movieinfo_ctr := applyDockerDefaults(spec, movieinfo_service, "movieinfo_proc", "movieinfo_container")
	containers = append(containers, movieinfo_ctr)

	castinfo_service := workflow.Service[mediamicroservices_sql.CastInfoService](spec, "castinfo_service", castinfo_db)
	castinfo_ctr := applyDockerDefaults(spec, castinfo_service, "castinfo_proc", "castinfo_container")
	containers = append(containers, castinfo_ctr)

	plot_service := workflow.Service[mediamicroservices_sql.PlotService](spec, "plot_service", plot_db)
	plot_ctr := applyDockerDefaults(spec, plot_service, "plot_ctr", "plot_container")
	containers = append(containers, plot_ctr)

	api_service := workflow.Service[mediamicroservices_sql.APIService](spec, "api_service", movieid_service, movieinfo_service, castinfo_service, plot_service)
	api_ctr := applyHTTPDefaults(spec, api_service, "api_proc", "api_container")
	containers = append(containers, api_ctr)

	return containers, nil
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
