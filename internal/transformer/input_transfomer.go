package transformer

import (
	"fmt"
	"strconv"
)

const (
	minRow    = 'A'
	maxRow    = 'F'
	maxColumn = 36
	endWidth  = '/'
)

type Coordinate struct {
	RowIndex, ColumnIndex int
}
type InputTransformer struct {
	minRow    rune
	maxRow    rune
	maxColumn uint8
}

func NewInputTransformer() *InputTransformer {
	return &InputTransformer{
		minRow:    minRow,
		maxRow:    maxRow,
		maxColumn: maxColumn,
	}
}

func (t *InputTransformer) Transform(input string, offset int) []Coordinate {
	if input == "" {
		return []Coordinate{}
	}

	coordinates := make([]Coordinate, 0)
	currentStart := 0

	// Add sentinel character
	input = input + string(endWidth)

	for i := 1; i < len(input); i++ {
		currentChar := rune(input[i])

		if (currentChar >= t.minRow && currentChar <= t.maxRow) || i == len(input)-1 {
			if segment := input[currentStart:i]; segment != "" {
				if coord, err := t.parseCoordinate(segment, offset); err == nil && t.isValidPosition(coord) {
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

func (t *InputTransformer) parseCoordinate(input string, offset int) (Coordinate, error) {
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
	shiftedColumn := col + offset

	return Coordinate{
		RowIndex:    int(row - t.minRow),
		ColumnIndex: shiftedColumn,
	}, nil
}

// Validation methods
func (t *InputTransformer) isValidPosition(coord Coordinate) bool {
	return t.isValidRow(rune(coord.RowIndex+int(minRow))) && t.isValidColumn(coord.ColumnIndex)
}

func (t *InputTransformer) isValidRow(char rune) bool {
	return char >= t.minRow && char <= t.maxRow
}

func (t *InputTransformer) isValidColumn(digit int) bool {
	return digit >= 0 && digit < int(t.maxColumn)
}
