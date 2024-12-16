package view

import (
	"elecsign/internal/transformer"
	"testing"
)

func TestViews_TurnOn(t *testing.T) {
	tests := []struct {
		name     string
		coords   []transformer.Coordinate
		expected map[int]byte
	}{
		{
			name: "multiple coordinates",
			coords: []transformer.Coordinate{
				{RowIndex: 0, ColumnIndex: 0}, // First bit in first byte
				{RowIndex: 1, ColumnIndex: 5}, // Bit 5 in byte 4
				{RowIndex: 5, ColumnIndex: 7}, // Last bit in byte 26
			},
			expected: map[int]byte{
				0:  0b10000000, // First coordinate
				4:  0b00000000, // Second coordinate
				26: 0b00000000, // Third coordinate
			},
		},
		{
			name: "row coordinates",
			coords: []transformer.Coordinate{
				{RowIndex: 0, ColumnIndex: 0}, // First bit in byte 0
				{RowIndex: 1, ColumnIndex: 3}, // Bit 3 in byte 4
				{RowIndex: 2, ColumnIndex: 0}, // First bit in byte 9
				{RowIndex: 3, ColumnIndex: 3}, // Bit 3 in byte 13
				{RowIndex: 5, ColumnIndex: 3}, // Bit 3 in byte 22
			},
			expected: map[int]byte{
				0:  0b10000000,
				4:  0b00000001,
				9:  0b10000000,
				13: 0b00000001,
				22: 0b00000001,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NewView()
			v.TurnOn(tt.coords)

			got := v.Get()
			for byteIdx, expected := range tt.expected {
				if got.Data[byteIdx] != expected {
					t.Errorf("byte[%d] = %08b, want %08b",
						byteIdx, got.Data[byteIdx], expected)
				}
			}
		})
	}
}
