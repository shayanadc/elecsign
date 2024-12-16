package internal

import (
	"testing"
)

func TestViews_TurnOn(t *testing.T) {
	tests := []struct {
		name     string
		coords   []Coordinate
		expected map[int]byte
	}{
		{
			name: "multiple coordinates",
			coords: []Coordinate{
				{rowIndex: 0, columnIndex: 0}, // First bit in first byte
				{rowIndex: 1, columnIndex: 5}, // Bit 5 in byte 4
				{rowIndex: 5, columnIndex: 7}, // Last bit in byte 26
			},
			expected: map[int]byte{
				0:  0b10000000, // First coordinate
				4:  0b00000000, // Second coordinate
				26: 0b00000000, // Third coordinate
			},
		},
		{
			name: "row coordinates",
			coords: []Coordinate{
				{rowIndex: 0, columnIndex: 0}, // First bit in byte 0
				{rowIndex: 1, columnIndex: 3}, // Bit 3 in byte 4
				{rowIndex: 2, columnIndex: 0}, // First bit in byte 9
				{rowIndex: 3, columnIndex: 3}, // Bit 3 in byte 13
				{rowIndex: 5, columnIndex: 3}, // Bit 3 in byte 22
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
			v := NewView(36, 6)
			v.TurnOn(tt.coords)

			got := v.Get()
			for byteIdx, expected := range tt.expected {
				if got.data[byteIdx] != expected {
					t.Errorf("byte[%d] = %08b, want %08b",
						byteIdx, got.data[byteIdx], expected)
				}
			}
		})
	}
}
