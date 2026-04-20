package main

import (
	"github.com/blueprint-uservices/blueprint/plugins/cmdbuilder"

	"github.com/blueprint-uservices/blueprint/examples/digota/wiring/specs"
)

func main() {
	name := "digota"
	cmdbuilder.MakeAndExecute(
		name,
		specs.Docker,
	)
}
