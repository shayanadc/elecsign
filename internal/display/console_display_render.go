package display

import (
	"fmt"
	"os"
	"strings"

	"elecsign/internal/transformer"
	"elecsign/internal/view"
)

const initialViewCapacity = 10

type Display interface {
	Show()
	AddView(view.View)
	Clear()
}

// Renderer interface for rendering a grid
type Renderer interface {
	Render(view view.View)
}

// ConsoleRenderer struct for rendering grids to the console
type ConsoleRenderer struct{}

// NewConsoleRenderer creates a new ConsoleRenderer
func NewConsoleRenderer() *ConsoleRenderer {
	return &ConsoleRenderer{}
}

// Render method to display the grid in a human-readable format
func (c *ConsoleRenderer) Render(view view.View) {
	var output strings.Builder
	width, height := view.Dimennsions()
	output.Grow(height * (width + 1))

	coord := transformer.Coordinate{}
	for row := 0; row < height; row++ {
		coord.RowIndex = row
		for col := 0; col < width; col++ {
			coord.ColumnIndex = col
			if on := view.IsOn(coord); on {
				output.WriteByte('*')
			} else {
				output.WriteByte(' ')
			}
		}
		output.WriteByte('\n')
	}

	// Avoid string allocation from String()
	fmt.Fprint(os.Stdout, output.String())
}

// ConsoleDisplay struct to manage views and render them
type ConsoleDisplay struct {
	views    []view.View
	renderer Renderer
}

// NewConsoleDisplay creates a new ConsoleDisplay
func NewConsoleDisplay(r Renderer) *ConsoleDisplay {
	return &ConsoleDisplay{
		views:    make([]view.View, 0, initialViewCapacity),
		renderer: r,
	}
}

// Show renders all views in the ConsoleDisplay
func (d *ConsoleDisplay) Show() {
	for _, view := range d.views {
		d.renderer.Render(view)
	}
}

// AddView adds a new view to the ConsoleDisplay
func (d *ConsoleDisplay) AddView(v view.View) {
	d.views = append(d.views, v)
}

// Clear removes all views from the ConsoleDisplay
func (d *ConsoleDisplay) Clear() {
	d.views = d.views[:0]
}
