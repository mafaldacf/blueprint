module github.com/blueprint-uservices/blueprint/examples/digota/tests

go 1.21

toolchain go1.22.1

require github.com/blueprint-uservices/blueprint/runtime v0.0.0

replace github.com/blueprint-uservices/blueprint/runtime => ../../../runtime

require (
	github.com/blueprint-uservices/blueprint/examples/digota/workflow v0.0.0
	github.com/stretchr/testify v1.9.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.mongodb.org/mongo-driver v1.15.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
