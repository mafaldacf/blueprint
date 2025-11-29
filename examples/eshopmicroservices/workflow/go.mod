module github.com/blueprint-uservices/blueprint/examples/eshopmicroservices/workflow

go 1.22.4

require (
	github.com/blueprint-uservices/blueprint/runtime v0.0.0
	github.com/google/uuid v1.6.0
	go.mongodb.org/mongo-driver v1.17.4
)

require (
	go.opentelemetry.io/otel v1.26.0 // indirect
	go.opentelemetry.io/otel/metric v1.26.0 // indirect
	go.opentelemetry.io/otel/trace v1.26.0 // indirect
	golang.org/x/exp v0.0.0-20240525044651-4c93da0ed11d // indirect
)

replace github.com/blueprint-uservices/blueprint/blueprint => ../../../blueprint
