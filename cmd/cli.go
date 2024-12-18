package cmd

import (
	"bufio"
	"elecsign/internal/display"
	"fmt"
	"os"
	"strings"
)

type CLI struct {
	display display.Display
	scanner *bufio.Scanner
}

func NewCLI(display display.Display) *CLI {
	return &CLI{
		display: display,
		scanner: bufio.NewScanner(os.Stdin),
	}
}

func (c *CLI) Run() {
	c.printHelp()

	handler := NewCommandHandler(c.display)

	for {
		fmt.Print("> ")
		if !c.scanner.Scan() {
			break
		}

		if err := c.executeCommand(c.scanner.Text(), handler); err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
}

func (c *CLI) executeCommand(input string, handler *CommandHandler) error {
	args := strings.Fields(input)
	if len(args) == 0 {
		return nil
	}

	switch args[0] {
	case "add":
		return handler.HandleAdd(args[1:])
	case "show":
		handler.HandleShow()
		return nil
	case "clear":
		handler.HandleClear()
		return nil
	case "exit":
		fmt.Println("Goodbye!")
		return nil
	default:
		return fmt.Errorf("unknown command: %s", args[0])
	}
}

func (c *CLI) printHelp() {
	fmt.Println("Electronic Sign CLI")
	fmt.Println("Commands:")
	fmt.Println("  add <type> <text> - Add a new view (type: pixel or character)")
	fmt.Println("  show            - Display all views")
	fmt.Println("  clear           - Clear all views")
	fmt.Println("  exit            - Exit the program")
}
