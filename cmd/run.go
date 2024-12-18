package cmd

import "elecsign/internal/display"

func RunCLI(display display.Display) {
	cli := NewCLI(display)
	cli.Run()
}
