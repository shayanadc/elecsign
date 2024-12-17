package main

import (
	"elecsign/cmd"
	"elecsign/internal/display"
)

func main() {
	renderer := display.NewConsoleRenderer()
	display := display.NewConsoleDisplay(renderer)

	cmd.RunCLI(display)
}
