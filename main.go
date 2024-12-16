package main

import (
	"bufio"
	"elecsign/internal"
	"fmt"
	"os"
	"strings"
)

func main() {
	renderer := &internal.ConsoleRenderer{}
	display := internal.NewConsoleDisplay(renderer)

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
			if len(args) < 3 {
				fmt.Println("Usage: add <type> <text>")
				fmt.Println("Types: pixel, character")
				continue
			}

			transformerType := args[1]
			if transformerType != "pixel" && transformerType != "character" {
				fmt.Println("Invalid type. Use 'pixel' or 'character'")
				continue
			}

			text := strings.Join(args[2:], " ")
			// Convert string to internal.TransformerType
			transformer := internal.TransformerType(transformerType)

			// Create the transformer and handle multiple return values
			transformerInstance, err := internal.NewTransformer(transformer)
			if err != nil {
				fmt.Printf("Error creating transformer: %v\n", err)
				continue
			}

			view := internal.NewView(36, 6)
			coordinates := transformerInstance.Transform(text)
			view.TurnOn(coordinates)
			display.AddView(view)
			fmt.Printf("View added with %s transformer\n", transformerType)

		case "show":
			fmt.Println("Displaying all views:")
			fmt.Println(strings.Repeat("-", 36))
			display.Show()
			fmt.Println(strings.Repeat("-", 36))
			display.Clear()

		case "clear":
			display.Clear()
			fmt.Println("All views cleared")

		case "exit":
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Unknown command. Available commands: add, show, clear, exit")
		}
	}
}
