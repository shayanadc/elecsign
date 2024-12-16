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
