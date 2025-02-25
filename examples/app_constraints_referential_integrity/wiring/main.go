// Package main provides an application for compiling different
// wiring specs for PostNotification application.
//
// To display options and usage, invoke:
//
//	go run main.go -h
package main

import (
	"github.com/blueprint-uservices/blueprint/plugins/cmdbuilder"

	"github.com/blueprint-uservices/blueprint/examples/app_constraints_referential_integrity/wiring/specs"
)

func main() {
	// Build a supported wiring spec
	name := "App_Constraints_Referential_Integrity"
	cmdbuilder.MakeAndExecute(
		name,
		specs.Docker,
	)
}
