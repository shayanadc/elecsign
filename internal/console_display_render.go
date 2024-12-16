package internal

import "fmt"

// ConsoleRenderer struct for rendering grids to the console
type ConsoleRenderer struct{}

// ConsoleDisplay struct to manage views and render them
type ConsoleDisplay struct {
	views    []View
	renderer Renderer
}

// NewConsoleDisplay creates a new ConsoleDisplay
func NewConsoleDisplay(r Renderer) *ConsoleDisplay {
	return &ConsoleDisplay{
		views:    make([]View, 0),
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
func (d *ConsoleDisplay) AddView(v View) {
	d.views = append(d.views, v)
}

// Clear removes all views from the ConsoleDisplay
func (d *ConsoleDisplay) Clear() {
	d.views = nil
}

// Render method to display the grid in a human-readable format
func (c *ConsoleRenderer) Render(grid Grid) {
	for row := 0; row < grid.height; row++ {
		for col := 0; col < grid.width; col++ {
			coord := Coordinate{rowIndex: row, columnIndex: col}
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
