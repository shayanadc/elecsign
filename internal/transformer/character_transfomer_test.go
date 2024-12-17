package transformer

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
				{RowIndex: 0, ColumnIndex: 2},
				{RowIndex: 0, ColumnIndex: 3},
				{RowIndex: 1, ColumnIndex: 1},
				{RowIndex: 1, ColumnIndex: 4},
				{RowIndex: 2, ColumnIndex: 0},
				{RowIndex: 2, ColumnIndex: 1},
				{RowIndex: 2, ColumnIndex: 2},
				{RowIndex: 2, ColumnIndex: 3},
				{RowIndex: 2, ColumnIndex: 4},
				{RowIndex: 3, ColumnIndex: 0},
				{RowIndex: 3, ColumnIndex: 4},
				{RowIndex: 4, ColumnIndex: 0},
				{RowIndex: 4, ColumnIndex: 4},
				{RowIndex: 5, ColumnIndex: 0},
				{RowIndex: 5, ColumnIndex: 4},
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
			got := transformer.Transform(tt.input, 0)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Transform(%q) = %v; want %v",
					tt.input, got, tt.want)
			}
		})
	}
}
