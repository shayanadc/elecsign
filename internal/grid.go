package internal

import "fmt"

const (
	byteSize = 8
)

type Grid struct {
	data          [27]byte // 27 bytes = (6 rows * 36 columns) / 8 bits rounded up
	width, height int
}

func NewGrid(height, width int) *Grid {
	return &Grid{
		height: height,
		width:  width,
		data:   [27]byte{},
	}
}

func (g *Grid) isValidCoordinate(coord Coordinate) bool {
	return coord.rowIndex >= 0 && coord.rowIndex < g.height &&
		coord.columnIndex >= 0 && coord.columnIndex < g.width
}

func (g *Grid) TurnOn(coord Coordinate) error {
	if !g.isValidCoordinate(coord) {
		return fmt.Errorf("invalid coordinate: row=%d, col=%d", coord.rowIndex, coord.columnIndex)
	}

	position := g.calculatePosition(coord)
	byteIndex := position / byteSize
	bitIndex := position % byteSize

	// Calculate bit position from right to left (7 to 0)
	// For bit position 3, we want 00001000
	g.data[byteIndex] |= (1 << (7 - bitIndex))

	return nil
}

func (g *Grid) calculatePosition(coord Coordinate) int {
	// Calculate the absolute position in the grid
	return (coord.rowIndex * g.width) + coord.columnIndex
}

func (g *Grid) IsOn(coord Coordinate) (bool, error) {
	if !g.isValidCoordinate(coord) {
		return false, fmt.Errorf("invalid coordinate: row=%d, col=%d", coord.rowIndex, coord.columnIndex)
	}

	position := g.calculatePosition(coord)
	byteIndex := position / byteSize
	bitIndex := position % byteSize

	return g.data[byteIndex]&(1<<((byteSize-1)-bitIndex)) != 0, nil
}
