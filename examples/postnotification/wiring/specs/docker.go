package specs

import (
	"github.com/blueprint-uservices/blueprint/examples/postnotification/workflow/postnotification"

	"github.com/blueprint-uservices/blueprint/blueprint/pkg/wiring"
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

	post_db := mongodb.Container(spec, "post_db")
	post_cache := redis.Container(spec, "post_cache")
	notification_queue := rabbitmq.Container(spec, "notification_queue", "notification")

	allServices = append(allServices, post_db)
	allServices = append(allServices, post_cache)
	allServices = append(allServices, notification_queue)
	
	storage_service := workflow.Service[postnotification.StorageService](spec, "storage_service", post_cache, post_db)
	storage_service_ctr := applyDockerDefaults(spec, storage_service, "storage_service_proc", "storage_service_container")
	containers = append(containers, storage_service_ctr)
	allServices = append(allServices, "storage_service")

	notify_service := workflow.Service[postnotification.NotifyService](spec, "notify_service", storage_service, notification_queue)
	notify_service_ctr := applyDockerQueueHandlerDefaults(spec, notify_service, "notify_service_proc", "notify_service_container")
	containers = append(containers, notify_service_ctr)
	/* allServices = append(allServices, "notify_service") */

	upload_service := workflow.Service[postnotification.UploadService](spec, "upload_service", storage_service, /* notify_service, */ notification_queue)
	upload_service_ctr := applyHTTPDefaults(spec, upload_service, "upload_service_proc", "upload_service_container")
	containers = append(containers, upload_service_ctr)
	allServices = append(allServices, "upload_service")

	tests := gotests.Test(spec, allServices...)
	containers = append(containers, tests)

	return containers, nil
}
