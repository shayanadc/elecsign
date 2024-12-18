package transformer

const offset = 6

// CharacterPattern represents the flyweight object
type CharacterPattern struct {
	Coordinates []Coordinate
}

// CharacterPatternFactory manages the flyweight objects
type CharacterPatternFactory struct {
	patterns map[rune]CharacterPattern
}

func NewCharacterPatternFactory() *CharacterPatternFactory {
	factory := &CharacterPatternFactory{
		patterns: make(map[rune]CharacterPattern, offset),
	}
	factory.initializePatterns()
	return factory
}

// Flyweight factory method to initialize patterns
func (f *CharacterPatternFactory) initializePatterns() {
	f.patterns['A'] = CharacterPattern{
		Coordinates: []Coordinate{
			{RowIndex: 0, ColumnIndex: 2}, {RowIndex: 0, ColumnIndex: 3},
			{RowIndex: 1, ColumnIndex: 1}, {RowIndex: 1, ColumnIndex: 4},
			{RowIndex: 2, ColumnIndex: 0}, {RowIndex: 2, ColumnIndex: 1},
			{RowIndex: 2, ColumnIndex: 2}, {RowIndex: 2, ColumnIndex: 3},
			{RowIndex: 2, ColumnIndex: 4}, {RowIndex: 3, ColumnIndex: 0},
			{RowIndex: 3, ColumnIndex: 4}, {RowIndex: 4, ColumnIndex: 0},
			{RowIndex: 4, ColumnIndex: 4}, {RowIndex: 5, ColumnIndex: 0},
			{RowIndex: 5, ColumnIndex: 4},
		},
	}

	f.patterns['B'] = CharacterPattern{
		Coordinates: []Coordinate{
			{RowIndex: 0, ColumnIndex: 0}, {RowIndex: 0, ColumnIndex: 1},
			{RowIndex: 0, ColumnIndex: 2}, {RowIndex: 0, ColumnIndex: 3},
			{RowIndex: 1, ColumnIndex: 0}, {RowIndex: 1, ColumnIndex: 4},
			{RowIndex: 2, ColumnIndex: 0}, {RowIndex: 2, ColumnIndex: 1},
			{RowIndex: 2, ColumnIndex: 2}, {RowIndex: 2, ColumnIndex: 3},
			{RowIndex: 3, ColumnIndex: 0}, {RowIndex: 3, ColumnIndex: 4},
			{RowIndex: 4, ColumnIndex: 0}, {RowIndex: 4, ColumnIndex: 4},
			{RowIndex: 5, ColumnIndex: 0}, {RowIndex: 5, ColumnIndex: 1},
			{RowIndex: 5, ColumnIndex: 2}, {RowIndex: 5, ColumnIndex: 3},
		},
	}

	f.patterns['C'] = CharacterPattern{
		Coordinates: []Coordinate{
			{RowIndex: 0, ColumnIndex: 1}, {RowIndex: 0, ColumnIndex: 2},
			{RowIndex: 0, ColumnIndex: 3}, {RowIndex: 1, ColumnIndex: 0},
			{RowIndex: 1, ColumnIndex: 4}, {RowIndex: 2, ColumnIndex: 0},
			{RowIndex: 3, ColumnIndex: 0}, {RowIndex: 4, ColumnIndex: 0},
			{RowIndex: 4, ColumnIndex: 4}, {RowIndex: 5, ColumnIndex: 1},
			{RowIndex: 5, ColumnIndex: 2}, {RowIndex: 5, ColumnIndex: 3},
		},
	}

	f.patterns['1'] = CharacterPattern{
		Coordinates: []Coordinate{
			{RowIndex: 0, ColumnIndex: 2}, {RowIndex: 1, ColumnIndex: 1},
			{RowIndex: 1, ColumnIndex: 2}, {RowIndex: 2, ColumnIndex: 2},
			{RowIndex: 3, ColumnIndex: 2}, {RowIndex: 4, ColumnIndex: 2},
			{RowIndex: 5, ColumnIndex: 2},
		},
	}

	f.patterns['2'] = CharacterPattern{
		Coordinates: []Coordinate{
			{RowIndex: 0, ColumnIndex: 1}, {RowIndex: 0, ColumnIndex: 2},
			{RowIndex: 0, ColumnIndex: 3}, {RowIndex: 1, ColumnIndex: 4},
			{RowIndex: 2, ColumnIndex: 3}, {RowIndex: 3, ColumnIndex: 2},
			{RowIndex: 4, ColumnIndex: 1}, {RowIndex: 5, ColumnIndex: 1},
			{RowIndex: 5, ColumnIndex: 2}, {RowIndex: 5, ColumnIndex: 3},
			{RowIndex: 5, ColumnIndex: 4},
		},
	}

	f.patterns['3'] = CharacterPattern{
		Coordinates: []Coordinate{
			{RowIndex: 0, ColumnIndex: 1}, {RowIndex: 0, ColumnIndex: 2},
			{RowIndex: 0, ColumnIndex: 3}, {RowIndex: 1, ColumnIndex: 4},
			{RowIndex: 2, ColumnIndex: 2}, {RowIndex: 2, ColumnIndex: 3},
			{RowIndex: 3, ColumnIndex: 4}, {RowIndex: 4, ColumnIndex: 4},
			{RowIndex: 5, ColumnIndex: 1}, {RowIndex: 5, ColumnIndex: 2},
			{RowIndex: 5, ColumnIndex: 3},
		},
	}
}

// CharacterTransformer uses the flyweight factory
type CharacterTransformer struct {
	patternFactory *CharacterPatternFactory
	characterWidth int
}

func NewCharacterTransformer() *CharacterTransformer {
	return &CharacterTransformer{
		patternFactory: NewCharacterPatternFactory(),
		characterWidth: offset,
	}
}

func (t *CharacterTransformer) Transform(input string, startOffset int) []Coordinate {
	allCoordinates := make([]Coordinate, 0, len(input)/2)

	// Process each character in the input string
	for i, char := range input {
		if pattern, exists := t.patternFactory.patterns[char]; exists {
			// Calculate offset for this character
			charOffset := i * t.characterWidth

			// Apply offset to each coordinate
			for _, coord := range pattern.Coordinates {
				newCoord := Coordinate{
					RowIndex:    coord.RowIndex,
					ColumnIndex: coord.ColumnIndex + charOffset + startOffset,
				}
				allCoordinates = append(allCoordinates, newCoord)
			}
		}
	}

	return allCoordinates
}
