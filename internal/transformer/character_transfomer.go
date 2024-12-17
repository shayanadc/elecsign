package transformer

const offset = 6

type CharacterTransformer struct {
	InputTransformer
	characterPatterns map[rune]string
	characterWidth    int
}

func NewCharacterTransformer() *CharacterTransformer {
	return &CharacterTransformer{
		InputTransformer: *NewInputTransformer(),
		characterPatterns: map[rune]string{
			'A': "A2A3B1B4C0C1C2C3C4D0D4E0E4F0F4",
			'B': "A0A1A2A3B0B4C0C1C2C3D0D4E0E4F0F1F2F3",
			'C': "A1A2A3B0B4C0D0E0E4F1F2F3",
			'1': "A2B1B2C2D2E2F2",
			'2': "A1A2A3B4C3D2E1F1F2F3F4",
			'3': "A1A2A3B4C2C3D4E4F1F2F3",
		},
		characterWidth: offset,
	}
}

func (t *CharacterTransformer) Transform(input string, offset int) []Coordinate {
	if input == "" {
		return []Coordinate{}
	}

	var allCoordinates []Coordinate

	// Process each character in the input string
	for i, char := range input {
		if pattern, exists := t.characterPatterns[char]; exists {
			// Calculate offset for this character
			offset := i * t.characterWidth

			// Transform the pattern using the base InputTransformer
			baseCoordinates := t.InputTransformer.Transform(pattern, offset)
			allCoordinates = append(allCoordinates, baseCoordinates...)
		}
	}

	return allCoordinates
}
