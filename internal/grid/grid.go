package grid

import (
	"errors"
	"fmt"

	"elecsign/internal/transformer"
)

var ErrInvalidCoordinate = errors.New("invalid coordinate")

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
func (g *Grid) IsOn(coord transformer.Coordinate) (bool, error) {
	if !g.isValidCoordinate(coord) {
		return false, fmt.Errorf("%w: row=%d, col=%d",
			ErrInvalidCoordinate,
			coord.RowIndex,
			coord.ColumnIndex)
	}

	byteIndex, bitIndex := g.toBitPosition(coord)
	return g.Data[byteIndex]&(1<<((byteSize-1)-bitIndex)) != 0, nil
}

func (g *Grid) TurnOn(coord transformer.Coordinate) error {
	if !g.isValidCoordinate(coord) {
		return fmt.Errorf("%w: row=%d, col=%d",
			ErrInvalidCoordinate,
			coord.RowIndex,
			coord.ColumnIndex)
	}

	byteIdx, bitIdx := g.toBitPosition(coord)

	// Calculate bit position from right to left (7 to 0)
	// For bit position 3, we want 00010000
	g.Data[byteIdx] |= (1 << ((byteSize - 1) - bitIdx))

	return nil
}

func (g *Grid) absolutePosition(coord transformer.Coordinate) int {
	// Calculate the absolute position in the grid
	return (coord.RowIndex * g.Width) + coord.ColumnIndex
}

func (g *Grid) toBitPosition(coord transformer.Coordinate) (byteIndex, bitIndex int) {
	position := g.absolutePosition(coord)
	byteIndex = position / byteSize
	bitIndex = position % byteSize
	return
}

func (g *Grid) isValidCoordinate(coord transformer.Coordinate) bool {
	return coord.RowIndex >= 0 && coord.RowIndex < g.Height &&
		coord.ColumnIndex >= 0 && coord.ColumnIndex < g.Width
}
