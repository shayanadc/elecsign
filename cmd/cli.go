package cmd

import (
	"bufio"
	"elecsign/internal/display"
	"elecsign/internal/transformer"
	"elecsign/internal/view"
	"fmt"
	"os"
	"strings"
)

func RunCLI() {
	renderer := &display.ConsoleRenderer{}
	display := display.NewConsoleDisplay(renderer)

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

func handleAddCommand(args []string, display *display.ConsoleDisplay) {
	if len(args) < 2 {
		fmt.Println("Usage: add <type> <text>")
		fmt.Println("Types: pixel, character")
		return
	}

	inputType := args[0]
	if inputType != "pixel" && inputType != "character" {
		fmt.Println("Invalid type. Use 'pixel' or 'character'")
		return
	}

	text := strings.Join(args[1:], " ")
	transformerType := transformer.TransformerType(inputType)

	// Call NewTransformer directly from the transformer package
	transformerInstance, err := transformer.NewTransformer(transformerType) // This line is incorrect
	if err != nil {
		fmt.Printf("Error creating transformer: %v\n", err)
		return
	}

	view := view.NewView(36, 6)
	coordinates := transformerInstance.Transform(text)
	view.TurnOn(coordinates)
	display.AddView(view)
	fmt.Printf("View added with %s transformer\n", transformerType)
}

func handleShowCommand(display *display.ConsoleDisplay) {
	fmt.Println("Displaying all views:")
	fmt.Println(strings.Repeat("-", 36))
	display.Show()
	fmt.Println(strings.Repeat("-", 36))
	display.Clear()
}

func handleClearCommand(display *display.ConsoleDisplay) {
	display.Clear()
	fmt.Println("All views cleared")
}
