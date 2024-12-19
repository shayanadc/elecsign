package view

import (
	"elecsign/internal/transformer"
	"testing"
)

func TestViews_TurnOn(t *testing.T) {
	tests := []struct {
		name        string
		coords      []transformer.Coordinate
		checkPoints map[transformer.Coordinate]bool
	}{
		{
			name: "multiple coordinates",
			coords: []transformer.Coordinate{
				{RowIndex: 0, ColumnIndex: 0},
				{RowIndex: 1, ColumnIndex: 5},
				{RowIndex: 5, ColumnIndex: 7},
			},
			checkPoints: map[transformer.Coordinate]bool{
				{RowIndex: 0, ColumnIndex: 0}: true,
				{RowIndex: 1, ColumnIndex: 5}: true,
				{RowIndex: 5, ColumnIndex: 7}: true,
				{RowIndex: 0, ColumnIndex: 1}: false, // verify off pixel
				{RowIndex: 2, ColumnIndex: 2}: false, // verify off pixel
			},
		},
		{
			name: "row coordinates",
			coords: []transformer.Coordinate{
				{RowIndex: 0, ColumnIndex: 0},
				{RowIndex: 1, ColumnIndex: 3},
				{RowIndex: 2, ColumnIndex: 0},
				{RowIndex: 3, ColumnIndex: 3},
				{RowIndex: 5, ColumnIndex: 3},
			},
			checkPoints: map[transformer.Coordinate]bool{
				{RowIndex: 0, ColumnIndex: 0}: true,
				{RowIndex: 1, ColumnIndex: 3}: true,
				{RowIndex: 2, ColumnIndex: 0}: true,
				{RowIndex: 3, ColumnIndex: 3}: true,
				{RowIndex: 5, ColumnIndex: 3}: true,
				{RowIndex: 0, ColumnIndex: 1}: false, // verify off pixel
				{RowIndex: 1, ColumnIndex: 4}: false, // verify off pixel
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NewView()
			v.TurnOn(tt.coords)

			// Check each coordinate's state
			for coord, expectedState := range tt.checkPoints {
				if got := v.IsOn(coord); got != expectedState {
					t.Errorf("coordinate (%d,%d) = %v, want %v",
						coord.RowIndex, coord.ColumnIndex, got, expectedState)
				}
			}
		})
	}
}
