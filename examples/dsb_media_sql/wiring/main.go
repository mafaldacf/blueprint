// Package main provides an application for compiling different
// wiring specs for DeathStarBench SocialNetwork application.
//
// To display options and usage, invoke:
//
//	go run main.go -h
package main

import (
	"github.com/blueprint-uservices/blueprint/examples/mediamicroservices_sql/wiring/specs"
	"github.com/blueprint-uservices/blueprint/plugins/cmdbuilder"
)

func main() {
	name := "MediaMicroservicesSQL"
	cmdbuilder.MakeAndExecute(
		name,
		specs.Docker,
	)
}
