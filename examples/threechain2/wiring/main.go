// Package main provides an application for compiling different
// wiring specs for threechain2 application.
//
// To display options and usage, invoke:
//
//	go run main.go -h
package main

import (
	"github.com/blueprint-uservices/blueprint/plugins/cmdbuilder"

	"github.com/blueprint-uservices/blueprint/examples/threechain2/wiring/specs"
)

func main() {
	// Build a supported wiring spec
	name := "threechain2"
	cmdbuilder.MakeAndExecute(
		name,
		specs.Docker,
	)
}
