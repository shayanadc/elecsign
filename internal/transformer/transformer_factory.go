package transformer

import "fmt"

var (
	ErrInvalidTransformerType = fmt.Errorf("invalid transformer type")
)

type Transformer interface {
	Transform(input string, offset int) []Coordinate
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
		return nil, ErrInvalidTransformerType
	}
}

// NewTransformerFromInput creates a new transformer based on the input type
func NewTransformerFromInput(inputType string) (Transformer, error) {
	transformerType := TransformerType(inputType)
	if transformerType != PixelType && transformerType != CharacterType {
		return nil, ErrInvalidTransformerType
	}
	return NewTransformer(transformerType)
}
