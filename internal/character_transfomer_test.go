package internal

import (
	"reflect"
	"testing"
)

func TestCharacterTransformer_Transform(t *testing.T) {
	transformer := NewCharacterTransformer()

	tests := []struct {
		name    string
		input   string
		pattern string
		want    []Coordinate
	}{
		{
			name:    "pattern A",
			input:   "A",
			pattern: "A2A3B1B4C0C1C2C3C4D0D4E0E4F0F4",
			want: []Coordinate{
				{rowIndex: 0, columnIndex: 2},
				{rowIndex: 0, columnIndex: 3},
				{rowIndex: 1, columnIndex: 1},
				{rowIndex: 1, columnIndex: 4},
				{rowIndex: 2, columnIndex: 0},
				{rowIndex: 2, columnIndex: 1},
				{rowIndex: 2, columnIndex: 2},
				{rowIndex: 2, columnIndex: 3},
				{rowIndex: 2, columnIndex: 4},
				{rowIndex: 3, columnIndex: 0},
				{rowIndex: 3, columnIndex: 4},
				{rowIndex: 4, columnIndex: 0},
				{rowIndex: 4, columnIndex: 4},
				{rowIndex: 5, columnIndex: 0},
				{rowIndex: 5, columnIndex: 4},
			},
		},
		{
			name:  "empty input",
			input: "",
			want:  []Coordinate{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := transformer.Transform(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Transform(%q) = %v; want %v",
					tt.input, got, tt.want)
			}
		})
	}
}
