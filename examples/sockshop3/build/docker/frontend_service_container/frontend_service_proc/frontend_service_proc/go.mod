module blueprint/goproc/frontend_service_proc

go 1.24.0

replace github.com/blueprint-uservices/blueprint/examples/sockshop3/workflow => ../workflow

replace github.com/blueprint-uservices/blueprint/runtime => ../runtime

require (
	github.com/apache/thrift v0.22.0
	github.com/blueprint-uservices/blueprint/examples/sockshop3/workflow v0.0.0-00010101000000-000000000000
	github.com/blueprint-uservices/blueprint/runtime v0.0.0-20240405152959-f078915d2306
	github.com/gorilla/mux v1.8.1
	golang.org/x/exp v0.0.0-20250813145105-42675adae3e6
)

require (
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/google/uuid v1.6.0 // indirect
	go.mongodb.org/mongo-driver v1.15.0 // indirect
	go.opentelemetry.io/otel v1.26.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdoutmetric v1.26.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.26.0 // indirect
	go.opentelemetry.io/otel/metric v1.26.0 // indirect
	go.opentelemetry.io/otel/sdk v1.26.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.26.0 // indirect
	go.opentelemetry.io/otel/trace v1.26.0 // indirect
	golang.org/x/sys v0.19.0 // indirect
)
