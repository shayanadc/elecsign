package display

import (
	"elecsign/internal/grid"
	"elecsign/internal/transformer"
	"elecsign/internal/view"
	"fmt"
)

type Display interface {
	Show()
	AddView(view.View)
	Clear()
}

// Renderer interface for rendering a grid
type Renderer interface {
	Render(grid grid.Grid)
}

// ConsoleRenderer struct for rendering grids to the console
type ConsoleRenderer struct{}

// NewConsoleRenderer creates a new ConsoleRenderer
func NewConsoleRenderer() *ConsoleRenderer {
	return &ConsoleRenderer{}
}

// Render method to display the grid in a human-readable format
func (c *ConsoleRenderer) Render(grid grid.Grid) {
	for row := 0; row < grid.Height; row++ {
		for col := 0; col < grid.Width; col++ {
			coord := transformer.Coordinate{RowIndex: row, ColumnIndex: col}
			on, _ := grid.IsOn(coord)
			if on {
				fmt.Print("*") // On pixel
			} else {
				fmt.Print(" ") // Off pixel
			}
		}
		fmt.Println() // New line after each row
	}
}

// ConsoleDisplay struct to manage views and render them
type ConsoleDisplay struct {
	views    []view.View
	renderer Renderer
}

// NewConsoleDisplay creates a new ConsoleDisplay
func NewConsoleDisplay(r Renderer) *ConsoleDisplay {
	return &ConsoleDisplay{
		views:    make([]view.View, 0),
		renderer: r,
	}
}

// Show renders all views in the ConsoleDisplay
func (d *ConsoleDisplay) Show() {
	for _, view := range d.views {
		d.renderer.Render(view.Get())
	}
}

// AddView adds a new view to the ConsoleDisplay
func (d *ConsoleDisplay) AddView(v view.View) {
	d.views = append(d.views, v)
}

// Clear removes all views from the ConsoleDisplay
func (d *ConsoleDisplay) Clear() {
	d.views = nil
}
