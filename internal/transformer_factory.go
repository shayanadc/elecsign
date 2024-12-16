package internal

import "fmt"

type Transformer interface {
	Transform(input string) []Coordinate
}

// Define the TransformerType type
type TransformerType string

// Define the available transformer types
const (
	PixelType     TransformerType = "pixel"
	CharacterType TransformerType = "character"
)

// TransformerFactory struct
type TransformerFactory struct{}

// NewTransformer function to create a new transformer based on the type
func NewTransformer(transformType TransformerType) (Transformer, error) {
	switch transformType {
	case PixelType:
		return NewInputTransformer(), nil
	case CharacterType:
		return NewCharacterTransformer(), nil
	default:
		return nil, fmt.Errorf("unknown transformer type: %s", transformType)
	}
}
