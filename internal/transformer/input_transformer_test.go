package transformer

import (
	"reflect"
	"testing"
)

func TestInputTransformer_Transform(t *testing.T) {
	transformer := NewInputTransformer()

	tests := []struct {
		name  string
		input string
		want  []Coordinate
	}{
		{
			name:  "single coordinate",
			input: "A0",
			want: []Coordinate{
				{RowIndex: 0, ColumnIndex: 0}, // A0 maps to (0,0)
			},
		},
		{
			name:  "multiple coordinates same row",
			input: "A0A1A2",
			want: []Coordinate{
				{RowIndex: 0, ColumnIndex: 0},
				{RowIndex: 0, ColumnIndex: 1},
				{RowIndex: 0, ColumnIndex: 2},
			},
		},
		{
			name:  "coordinates across rows",
			input: "A35B0",
			want: []Coordinate{
				{RowIndex: 0, ColumnIndex: 35},
				{RowIndex: 1, ColumnIndex: 0},
			},
		},
		{
			name:  "invalid positions filtered out",
			input: "A60A40",
			want:  []Coordinate{},
		},
		{
			name:  "empty input",
			input: "",
			want:  []Coordinate{},
		},
		{
			name:  "last valid positions",
			input: "F34F35",
			want: []Coordinate{
				{RowIndex: 5, ColumnIndex: 34},
				{RowIndex: 5, ColumnIndex: 35},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := transformer.Transform(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Transform(%q) = %+v; want %+v",
					tt.input, got, tt.want)
			}
		})
	}
}
