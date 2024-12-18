// cmd/handler.go
package cmd

import (
	"elecsign/internal/display"
	"elecsign/internal/transformer"
	"elecsign/internal/view"
	"errors"
	"fmt"
	"strings"
)

type CommandHandler struct {
	display display.Display
}

func NewCommandHandler(display display.Display) *CommandHandler {
	return &CommandHandler{
		display: display,
	}
}

func (h *CommandHandler) HandleAdd(args []string) error {
	if len(args) < 2 {
		return errors.New("usage: add <pixel, character> <text>")
	}

	inputType := args[0]
	transformerInstance, err := transformer.NewTransformerFromInput(inputType)
	if err != nil {
		return fmt.Errorf("invalid transformer type: %w", err)
	}

	text := strings.Join(args[1:], " ")
	coordinates := transformerInstance.Transform(text, 0)

	view := view.NewView()
	view.TurnOn(coordinates)
	h.display.AddView(view)

	return nil
}

func (h *CommandHandler) HandleShow() {
	h.display.Show()
}

func (h *CommandHandler) HandleClear() {
	h.display.Clear()
}
