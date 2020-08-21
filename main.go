package main

import (
	"github.com/kpfaulkner/slacker/pkg/slacker"
)

// Main file for the example slack app.

func main() {

	app := slacker.NewSlacker()

	app.SetupUI()
	app.Run()
}
