package internal

import (
	"fmt"
	"strconv"
)

type Coordinate struct {
	x, y int
}
type InputTransformer struct {
	minRow    rune
	maxRow    rune
	maxColumn int
}

func NewInputTransformer() *InputTransformer {
	return &InputTransformer{
		minRow:    'A',
		maxRow:    'F',
		maxColumn: 36,
	}
}

func (t *InputTransformer) Transform(input string) []Coordinate {
	if input == "" {
		return []Coordinate{} // This works for empty input
	}

	coordinates := make([]Coordinate, 0) // Initialize with zero length
	currentStart := 0

	// Add sentinel character
	input = input + "X"

	for i := 1; i < len(input); i++ {
		currentChar := rune(input[i])

		if (currentChar >= t.minRow && currentChar <= t.maxRow) || i == len(input)-1 {
			if segment := input[currentStart:i]; segment != "" {
				if coord, err := t.parseCoordinate(segment); err == nil && t.isValidPosition(coord) {
					coordinates = append(coordinates, coord)
				}
			}
			currentStart = i
		}
	}

	// If no valid coordinates were found, return empty slice
	if len(coordinates) == 0 {
		return []Coordinate{}
	}

	return coordinates
}

func (t *InputTransformer) parseCoordinate(input string) (Coordinate, error) {
	if len(input) < 2 {
		return Coordinate{}, fmt.Errorf("invalid input length: %d", len(input))
	}

	row := rune(input[0])
	if !t.isValidRow(row) {
		return Coordinate{}, fmt.Errorf("invalid row character: %c", row)
	}

	colStr := input[1:]
	col, err := strconv.Atoi(colStr)
	if err != nil {
		return Coordinate{}, fmt.Errorf("invalid column number: %s", colStr)
	}

	return Coordinate{
		x: int(row - t.minRow),
		y: col,
	}, nil
}

// Validation methods
func (t *InputTransformer) isValidPosition(coord Coordinate) bool {
	return t.isValidRow(rune(coord.x+int('A'))) && t.isValidColumn(coord.y)
}

func (t *InputTransformer) isValidRow(char rune) bool {
	return char >= t.minRow && char <= t.maxRow
}

func (t *InputTransformer) isValidColumn(digit int) bool {
	return digit >= 0 && digit < t.maxColumn
}
