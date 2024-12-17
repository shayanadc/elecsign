package grid

import (
	"elecsign/internal/transformer"
	"fmt"
)

const (
	byteSize     = 8
	gridDataSize = 27 // 27 bytes = (6 rows * 36 columns) / 8 bits rounded up
)

type Grid struct {
	Data          [gridDataSize]byte
	Width, Height int
}

func NewGrid(height, width int) *Grid {
	return &Grid{
		Height: height,
		Width:  width,
		Data:   [gridDataSize]byte{},
	}
}

func (g *Grid) isValidCoordinate(coord transformer.Coordinate) bool {
	return coord.RowIndex >= 0 && coord.RowIndex < g.Height &&
		coord.ColumnIndex >= 0 && coord.ColumnIndex < g.Width
}

func (g *Grid) TurnOn(coord transformer.Coordinate) error {
	if !g.isValidCoordinate(coord) {
		return fmt.Errorf("invalid coordinate: row=%d, col=%d", coord.RowIndex, coord.ColumnIndex)
	}

	position := g.calculatePosition(coord)
	byteIndex := position / byteSize
	bitIndex := position % byteSize

	// Calculate bit position from right to left (7 to 0)
	// For bit position 3, we want 00001000
	g.Data[byteIndex] |= (1 << ((byteSize - 1) - bitIndex))

	return nil
}

func (g *Grid) calculatePosition(coord transformer.Coordinate) int {
	// Calculate the absolute position in the grid
	return (coord.RowIndex * g.Width) + coord.ColumnIndex
}

func (g *Grid) IsOn(coord transformer.Coordinate) (bool, error) {
	if !g.isValidCoordinate(coord) {
		return false, fmt.Errorf("invalid coordinate: row=%d, col=%d", coord.RowIndex, coord.ColumnIndex)
	}

	position := g.calculatePosition(coord)
	byteIndex := position / byteSize
	bitIndex := position % byteSize

	return g.Data[byteIndex]&(1<<((byteSize-1)-bitIndex)) != 0, nil
}
