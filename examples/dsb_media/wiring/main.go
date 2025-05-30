// Package main provides an application for compiling different
// wiring specs for DeathStarBench SocialNetwork application.
//
// To display options and usage, invoke:
//
//	go run main.go -h
package main

import (
	"github.com/blueprint-uservices/blueprint/plugins/cmdbuilder"

	"github.com/blueprint-uservices/blueprint/examples/dsb_media/wiring/specs"
)

func main() {
	name := "MediaMicroservices"
	cmdbuilder.MakeAndExecute(
		name,
		specs.Docker,
	)
}
