package view

import (
	"elecsign/internal/grid"
	"elecsign/internal/transformer"
)

const (
	GridViewWidth  = 36 // Set your desired width
	GridViewHeight = 6  // Set your desired height
)

type GridView struct {
	data grid.Grid
}

// NewView creates a new View with a fixed-size Grid
func NewView() *GridView {
	return &GridView{
		data: *grid.NewGrid(GridViewHeight, GridViewWidth),
	}
}

func (v *GridView) Get() grid.Grid {
	return v.data
}

func (v *GridView) IsOn(c transformer.Coordinate) bool {
	on, err := v.data.IsOn(c)
	if err != nil {
		return false
	}

	return on
}

// TurnOn turns on multiple pixels at the given coordinates
func (v *GridView) TurnOn(coords []transformer.Coordinate) {
	for _, coord := range coords {
		_ = v.data.TurnOn(coord)
	}
}
