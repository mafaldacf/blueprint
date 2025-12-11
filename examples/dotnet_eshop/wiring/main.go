// Package main provides an application for compiling different
// wiring specs for PostNotification application.
//
// To display options and usage, invoke:
//
//	go run main.go -h
package main

import (
	"github.com/blueprint-uservices/blueprint/plugins/cmdbuilder"

	"github.com/blueprint-uservices/blueprint/examples/dotnet_eshop/wiring/specs"
)

func main() {
	// Build a supported wiring spec
	name := "dotnet_eshop"
	cmdbuilder.MakeAndExecute(
		name,
		specs.Docker,
	)
}
