package specs

import (
	"github.com/blueprint-uservices/blueprint/blueprint/pkg/wiring"
	"github.com/blueprint-uservices/blueprint/examples/app_constraints_specialization/workflow/app_constraints_specialization"
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
	
	freelancers_db := mongodb.Container(spec, "freelancers_db")
	fulltimes_db := mongodb.Container(spec, "fulltimes_db")
	interns_db := mongodb.Container(spec, "interns_db")
	employees_db := mongodb.Container(spec, "employees_db")

	allServices = append(allServices, freelancers_db)
	allServices = append(allServices, fulltimes_db)
	allServices = append(allServices, interns_db)
	allServices = append(allServices, employees_db)

	freelancer_service := workflow.Service[app_constraints_specialization.FreelancerService](spec, "freelancer_service", freelancers_db)
	freelancer_service_ctr := applyDockerDefaults(spec, freelancer_service, "freelancer_service_proc", "freelancer_service_container")
	containers = append(containers, freelancer_service_ctr)
	allServices = append(allServices, "freelancer_service")

	fulltime_service := workflow.Service[app_constraints_specialization.FulltimeService](spec, "fulltime_service", fulltimes_db)
	fulltime_service_ctr := applyDockerDefaults(spec, fulltime_service, "fulltime_service_proc", "fulltime_service_container")
	containers = append(containers, fulltime_service_ctr)
	allServices = append(allServices, "fulltime_service")

	intern_service := workflow.Service[app_constraints_specialization.InternService](spec, "intern_service", interns_db)
	intern_service_ctr := applyDockerDefaults(spec, intern_service, "intern_service_proc", "intern_service_container")
	containers = append(containers, intern_service_ctr)
	allServices = append(allServices, "intern_service")

	employee_service := workflow.Service[app_constraints_specialization.EmployeeService](spec, "employee_service", freelancer_service, fulltime_service, intern_service, employees_db)
	employee_service_ctr := applyDockerDefaults(spec, employee_service, "employee_service_proc", "employee_service_container")
	containers = append(containers, employee_service_ctr)
	allServices = append(allServices, "employee_service")

	frontend := workflow.Service[app_constraints_specialization.Frontend](spec, "frontend", employee_service, freelancer_service, fulltime_service, intern_service)
	frontend_ctr := applyHTTPDefaults(spec, frontend, "frontend_proc", "frontend_container")
	containers = append(containers, frontend_ctr)
	allServices = append(allServices, "frontend")

	tests := gotests.Test(spec, allServices...)
	containers = append(containers, tests)

	return containers, nil
}
