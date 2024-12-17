package cmd

import (
	"bufio"
	"elecsign/internal/display"
	"elecsign/internal/transformer"
	"elecsign/internal/view"
	"errors"
	"fmt"
	"os"
	"strings"
)

func RunCLI(display display.Display) {

	fmt.Println("Electronic Sign CLI")
	fmt.Println("Commands:")
	fmt.Println("  add <type> <text> - Add a new view (type: pixel or character)")
	fmt.Println("  show            - Display all views")
	fmt.Println("  clear           - Clear all views")
	fmt.Println("  exit            - Exit the program")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		input := scanner.Text()
		args := strings.Fields(input)
		if len(args) == 0 {
			continue
		}

		command := args[0]
		switch command {
		case "add":
			handleAddCommand(args[1:], display)
		case "show":
			handleShowCommand(display)
		case "clear":
			handleClearCommand(display)
		case "exit":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Unknown command. Available commands: add, show, clear, exit")
		}
	}
}

func handleAddCommand(args []string, display display.Display) {
	if len(args) < 2 {
		fmt.Println("Usage: add <pixel, character> <text>")
		return
	}

	inputType := args[0]
	transformerInstance, err := transformer.NewTransformerFromInput(inputType)
	if errors.Is(err, transformer.ErrInvalidTransformerType) {
		fmt.Println("Invalid transformer type provided")
		return
	}

	text := strings.Join(args[1:], " ")

	coordinates := transformerInstance.Transform(text, 0)

	view := view.NewView()
	view.TurnOn(coordinates)
	display.AddView(view)
	fmt.Printf("View added with %s transformer\n", inputType)
}

func handleShowCommand(display display.Display) {
	fmt.Println("Displaying all views:")
	display.Show()
	display.Clear()
	fmt.Println("All views cleared")
}

func handleClearCommand(display display.Display) {
	display.Clear()
	fmt.Println("All views cleared")
}
