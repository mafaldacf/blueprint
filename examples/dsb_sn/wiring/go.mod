module github.com/blueprint-uservices/blueprint/examples/dsb_sn/wiring

go 1.22.4

require github.com/blueprint-uservices/blueprint/examples/dsb_sn/workflow v0.0.0

require (
	github.com/blueprint-uservices/blueprint/blueprint v0.0.0-20250729202253-a8f505263256
	github.com/blueprint-uservices/blueprint/plugins v0.0.0-20250729202253-a8f505263256
	github.com/blueprint-uservices/blueprint/runtime v0.0.0-20250729202253-a8f505263256 // indirect
)

require (
	github.com/otiai10/copy v1.11.0 // indirect
	golang.org/x/exp v0.0.0-20231110203233-9a3e6036ecaa // indirect
	golang.org/x/mod v0.14.0 // indirect
	golang.org/x/sys v0.14.0 // indirect
	golang.org/x/tools v0.15.0 // indirect
)

replace github.com/blueprint-uservices/blueprint/examples/dsb_sn/workflow => ../workflow
