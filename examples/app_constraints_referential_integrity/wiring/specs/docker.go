package specs

import (
	"github.com/blueprint-uservices/blueprint/blueprint/pkg/wiring"
	"github.com/blueprint-uservices/blueprint/examples/app_constraints_referential_integrity/workflow/app_constraints_referential_integrity"
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
	
	users_db := mongodb.Container(spec, "users_db")
	accounts_db := mongodb.Container(spec, "accounts_db")

	allServices = append(allServices, users_db)
	allServices = append(allServices, accounts_db)

	user_service := workflow.Service[app_constraints_referential_integrity.UserService](spec, "user_service", users_db)
	user_service_ctr := applyDockerDefaults(spec, user_service, "user_service_proc", "user_service_container")
	containers = append(containers, user_service_ctr)
	allServices = append(allServices, "user_service")

	account_service := workflow.Service[app_constraints_referential_integrity.AccountService](spec, "account_service", user_service, accounts_db)
	account_service_ctr := applyDockerDefaults(spec, account_service, "account_service_proc", "account_service_container")
	containers = append(containers, account_service_ctr)
	allServices = append(allServices, "account_service")

	frontend := workflow.Service[app_constraints_referential_integrity.Frontend](spec, "frontend", account_service, user_service)
	frontend_ctr := applyHTTPDefaults(spec, frontend, "frontend_proc", "frontend_container")
	containers = append(containers, frontend_ctr)
	allServices = append(allServices, "frontend")

	tests := gotests.Test(spec, allServices...)
	containers = append(containers, tests)

	return containers, nil
}
