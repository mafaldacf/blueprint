module github.com/blueprint-uservices/blueprint/examples/digota/workflow

go 1.24.0

require (
	github.com/Rhymond/go-money v1.0.15
	github.com/blueprint-uservices/blueprint/runtime v0.0.0-20250815220819-bcd4a51069cb
	go.mongodb.org/mongo-driver v1.17.6
)

require (
	github.com/pkg/errors v0.9.1 // indirect
	go.opentelemetry.io/otel v1.32.0 // indirect
	go.opentelemetry.io/otel/metric v1.32.0 // indirect
	go.opentelemetry.io/otel/trace v1.26.0 // indirect
	golang.org/x/exp v0.0.0-20240416160154-fe59bbe5cc7f // indirect
)

replace go.opentelemetry.io/otel => go.opentelemetry.io/otel v1.26.0

replace go.opentelemetry.io/otel/sdk => go.opentelemetry.io/otel/sdk v1.26.0

replace go.opentelemetry.io/otel/metric => go.opentelemetry.io/otel/metric v1.26.0
