module github.com/blueprint-uservices/blueprint/examples/dsb_hotel/workload

go 1.22

toolchain go1.22.2

replace github.com/blueprint-uservices/blueprint/runtime => ../../../runtime

replace github.com/blueprint-uservices/blueprint/examples/dsb_hotel/workflow => ../workflow

require github.com/blueprint-uservices/blueprint/examples/dsb_hotel/workflow v0.0.0

require github.com/blueprint-uservices/blueprint/runtime v0.0.0

require (
	github.com/hailocab/go-geoindex v0.0.0-20160127134810-64631bfe9711 // indirect
	golang.org/x/exp v0.0.0-20240416160154-fe59bbe5cc7f // indirect
)
