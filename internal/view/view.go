package view

import (
	"elecsign/internal/grid"
	"elecsign/internal/transformer"
)

type View interface {
	IsOn(transformer.Coordinate) bool
	Get() grid.Grid
	TurnOn([]transformer.Coordinate)
}
