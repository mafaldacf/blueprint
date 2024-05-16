package specs

import (
	"github.com/blueprint-uservices/blueprint/examples/postnotification/workflow/postnotification"

	"github.com/blueprint-uservices/blueprint/blueprint/pkg/wiring"
	"github.com/blueprint-uservices/blueprint/plugins/cmdbuilder"
	"github.com/blueprint-uservices/blueprint/plugins/goproc"
	"github.com/blueprint-uservices/blueprint/plugins/gotests"
	"github.com/blueprint-uservices/blueprint/plugins/http"
	"github.com/blueprint-uservices/blueprint/plugins/linuxcontainer"
	"github.com/blueprint-uservices/blueprint/plugins/rabbitmq"
	"github.com/blueprint-uservices/blueprint/plugins/redis"
	"github.com/blueprint-uservices/blueprint/plugins/thrift"
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

	post_db := redis.Container(spec, "post_db")
	notification_queue := rabbitmq.Container(spec, "notification_queue", "notification")

	allServices = append(allServices, post_db)
	allServices = append(allServices, notification_queue)

	/* dummy_service := workflow.Service[postnotification.DummyService](spec, "dummy_service")
	dummy_service_ctr := applyDockerDefaults(spec, dummy_service, "dummy_service_proc", "dummy_service_container")
	containers = append(containers, dummy_service_ctr)
	allServices = append(allServices, "dummy_service") */

	storage_service := workflow.Service[postnotification.StorageService](spec, "storage_service", post_db)
	storage_service_ctr := applyDockerDefaults(spec, storage_service, "storage_service_proc", "storage_service_container")
	containers = append(containers, storage_service_ctr)
	allServices = append(allServices, "storage_service")

	notify_service := workflow.Service[postnotification.NotifyService](spec, "notify_service", storage_service, notification_queue)
	notify_service_ctr := applyDockerDefaults(spec, notify_service, "notify_service_proc", "notify_service_container")
	containers = append(containers, notify_service_ctr)

	upload_service := workflow.Service[postnotification.UploadService](spec, "upload_service", storage_service, notification_queue)
	upload_service_ctr := applyHTTPDefaults(spec, upload_service, "upload_service_proc", "upload_service_container")
	containers = append(containers, upload_service_ctr)
	allServices = append(allServices, "upload_service")

	tests := gotests.Test(spec, allServices...)
	containers = append(containers, tests)

	return containers, nil
}

func applyDockerQueueDefaults(spec wiring.WiringSpec, serviceName, procName, ctrName string) string {
	goproc.CreateProcess(spec, procName, serviceName)
	return linuxcontainer.CreateContainer(spec, ctrName, procName)
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
