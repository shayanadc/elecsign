package view

import (
	"elecsign/internal/grid"
	"elecsign/internal/transformer"
)

type View interface {
	IsOn(c transformer.Coordinate) bool
	Get() grid.Grid
	TurnOn([]transformer.Coordinate)
}

type GridView struct {
	data grid.Grid
}

// NewView creates a new View with a fixed-size Grid
func NewView(width, height int) *GridView {
	return &GridView{
		data: *grid.NewGrid(height, width),
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
