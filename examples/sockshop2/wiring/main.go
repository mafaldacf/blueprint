// An application for compiling the SockShop application.
// Provides a number of different wiring specs for compiling
// the application in different configurations.
//
// To display options and usage, invoke:
//
//	go run main.go -h
package main

import (
	"github.com/blueprint-uservices/blueprint/plugins/cmdbuilder"

	"github.com/blueprint-uservices/blueprint/examples/sockshop2/wiring/specs"
)

func main() {

	// Build a supported wiring spec
	name := "SockShop"
	cmdbuilder.MakeAndExecute(
		name,
		specs.Docker,
	)
}
