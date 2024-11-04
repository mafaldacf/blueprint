// Package main provides an application for compiling different
// wiring specs for shopping_simple application.
//
// To display options and usage, invoke:
//
//	go run main.go -h
package main

import (
	"github.com/blueprint-uservices/blueprint/plugins/cmdbuilder"

	"github.com/blueprint-uservices/blueprint/examples/shopping_simple/wiring/specs"
)

func main() {
	// Build a supported wiring spec
	name := "shopping_simple"
	cmdbuilder.MakeAndExecute(
		name,
		specs.Docker,
	)
}
