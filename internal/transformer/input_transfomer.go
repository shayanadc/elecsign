package transformer

import (
	"fmt"
	"strconv"
)

const (
	minRow        = 'A'
	maxRow        = 'F'
	endWidth      = '/'
	maxColumn     = 36
	minSignLength = 2
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
	coordinates := make([]Coordinate, 0, len(input)/2)
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

	return coordinates
}

func (t *InputTransformer) parseCoordinate(input string, offset int) (Coordinate, error) {
	if len(input) < minSignLength {
		return Coordinate{}, fmt.Errorf("invalid input length: %d", len(input))
	}

	row := rune(input[0])
	if !t.isValidRow(row) {
		return Coordinate{}, fmt.Errorf("invalid row character: %c", row)
	}

	col, err := strconv.Atoi(input[1:])
	if err != nil {
		return Coordinate{}, fmt.Errorf("invalid column number: %s", input[1:])
	}

	return Coordinate{
		RowIndex:    int(row - t.minRow),
		ColumnIndex: col + offset,
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
