module github.com/blueprint-uservices/blueprint/examples/leaf/workflow

go 1.22

toolchain go1.22.1

require (
	github.com/blueprint-uservices/blueprint/runtime v0.0.0
	go.mongodb.org/mongo-driver v1.17.6
	go.opentelemetry.io/otel/metric v1.32.0
)

require (
	github.com/go-logr/logr v1.4.3 // indirect
	github.com/stretchr/testify v1.10.0 // indirect
	go.opentelemetry.io/otel v1.32.0 // indirect
	go.opentelemetry.io/otel/trace v1.32.0 // indirect
	golang.org/x/exp v0.0.0-20240525044651-4c93da0ed11d // indirect
)

replace github.com/blueprint-uservices/blueprint/runtime => ../../../runtime
