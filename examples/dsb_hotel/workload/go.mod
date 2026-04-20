module github.com/blueprint-uservices/blueprint/examples/dsb_hotel/workload

go 1.24.0

replace github.com/blueprint-uservices/blueprint/runtime => ../../../runtime

replace github.com/blueprint-uservices/blueprint/examples/dsb_hotel/workflow => ../workflow

require github.com/blueprint-uservices/blueprint/examples/dsb_hotel/workflow v0.0.0

require github.com/blueprint-uservices/blueprint/runtime v0.0.0 // indirect

require (
	github.com/go-logr/logr v1.4.3 // indirect
	github.com/stretchr/testify v1.10.0 // indirect
	go.mongodb.org/mongo-driver v1.17.6 // indirect
	go.opentelemetry.io/otel v1.32.0 // indirect
	go.opentelemetry.io/otel/metric v1.32.0 // indirect
	go.opentelemetry.io/otel/trace v1.32.0 // indirect
	golang.org/x/exp v0.0.0-20240525044651-4c93da0ed11d // indirect
)
