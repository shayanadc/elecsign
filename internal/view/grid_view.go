package view

import (
	"elecsign/internal/grid"
	"elecsign/internal/transformer"
)

var (
	GridViewWidth  = 36
	GridViewHeight = 6
)

type View interface {
	IsOn(transformer.Coordinate) bool
	TurnOn([]transformer.Coordinate)
	Dimennsions() (int, int)
}

type GridView struct {
	data grid.Grid
}

// NewView creates a new View with a fixed-size Grid
func NewView() *GridView {
	return &GridView{
		data: *grid.NewGrid(GridViewHeight, GridViewWidth),
	}
}

func (v *GridView) Dimennsions() (int, int) {
	return GridViewWidth, GridViewHeight
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
		// Ignoring error as invalid coordinates need to be ignored (fault tolerance)
		_ = v.data.TurnOn(coord)
	}
}
