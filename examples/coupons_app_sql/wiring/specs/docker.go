package specs

import (
	"github.com/blueprint-uservices/blueprint/blueprint/pkg/wiring"
	"github.com/blueprint-uservices/blueprint/examples/coupons_app_sql/workflow/coupons_app_sql"
	"github.com/blueprint-uservices/blueprint/plugins/cmdbuilder"
	"github.com/blueprint-uservices/blueprint/plugins/gotests"
	"github.com/blueprint-uservices/blueprint/plugins/mysql"
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
	
	students_db := mysql.Container(spec, "students_db")
	coupons_db := mysql.Container(spec, "coupons_db")

	allServices = append(allServices, students_db)
	allServices = append(allServices, coupons_db)

	student_service := workflow.Service[coupons_app_sql.StudentService](spec, "student_service", students_db)
	student_service_ctr := applyDockerDefaults(spec, student_service, "student_service_proc", "student_service_container")
	containers = append(containers, student_service_ctr)
	allServices = append(allServices, "student_service")

	coupon_service := workflow.Service[coupons_app_sql.CouponService](spec, "coupon_service", coupons_db)
	coupon_service_ctr := applyDockerDefaults(spec, coupon_service, "coupon_service_proc", "coupon_service_container")
	containers = append(containers, coupon_service_ctr)
	allServices = append(allServices, "coupon_service")

	frontend := workflow.Service[coupons_app_sql.Frontend](spec, "frontend", student_service, coupon_service)
	frontend_ctr := applyHTTPDefaults(spec, frontend, "frontend_proc", "frontend_container")
	containers = append(containers, frontend_ctr)
	allServices = append(allServices, "frontend")

	tests := gotests.Test(spec, allServices...)
	containers = append(containers, tests)

	return containers, nil
}
