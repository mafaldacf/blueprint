// Package main provides an application for compiling different
// wiring specs for PostNotification application.
//
// To display options and usage, invoke:
//
//	go run main.go -h
package main

import (
	"github.com/blueprint-uservices/blueprint/examples/foobar/wiring/specs"
	"github.com/blueprint-uservices/blueprint/plugins/cmdbuilder"
)

func main() {
	// Build a supported wiring spec
	name := "FooBar"
	cmdbuilder.MakeAndExecute(
		name,
		specs.Docker,
	)
}
