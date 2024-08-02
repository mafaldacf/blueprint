module github.com/blueprint-uservices/blueprint/examples/sockshop2/workflow

go 1.22

toolchain go1.22.5

require (
	github.com/blueprint-uservices/blueprint/runtime v0.0.0
	github.com/google/uuid v1.6.0
	go.mongodb.org/mongo-driver v1.15.0
)

require (
	go.opentelemetry.io/otel v1.26.0 // indirect
	go.opentelemetry.io/otel/metric v1.26.0 // indirect
	go.opentelemetry.io/otel/trace v1.26.0 // indirect
	golang.org/x/exp v0.0.0-20240416160154-fe59bbe5cc7f // indirect
)

replace github.com/blueprint-uservices/blueprint/runtime => ../../../runtime
