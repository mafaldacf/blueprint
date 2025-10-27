package main

import (
	"github.com/blueprint-uservices/blueprint/plugins/cmdbuilder"

	"github.com/blueprint-uservices/blueprint/examples/large_scale_app/wiring/specs"
)

func main() {
	name := "LargeScaleApp"
	cmdbuilder.MakeAndExecute(
		name,
		specs.Docker,
	)
}
