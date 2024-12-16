package internal

import "testing"

func TestGrid_TurnOn(t *testing.T) {
	tests := []struct {
		name        string
		coord       Coordinate
		wantErr     bool
		expectedBit byte
	}{
		{
			name:        "valid coordinate middle of display",
			coord:       Coordinate{rowIndex: 2, columnIndex: 3},
			wantErr:     false,
			expectedBit: 0b00010000, // bit 4 set in the appropriate byte
		},
		{
			name:        "valid coordinate first position",
			coord:       Coordinate{rowIndex: 0, columnIndex: 0},
			wantErr:     false,
			expectedBit: 0b10000000, // leftmost bit set
		},
		{
			name:    "invalid coordinate negative row",
			coord:   Coordinate{rowIndex: -1, columnIndex: 0},
			wantErr: true,
		},
		{
			name:    "invalid coordinate negative column",
			coord:   Coordinate{rowIndex: 0, columnIndex: -1},
			wantErr: true,
		},
		{
			name:    "invalid coordinate row too large",
			coord:   Coordinate{rowIndex: 6, columnIndex: 0},
			wantErr: true,
		},
		{
			name:    "invalid coordinate column too large",
			coord:   Coordinate{rowIndex: 0, columnIndex: 36},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGrid(6, 36)
			err := g.TurnOn(tt.coord)

			if (err != nil) != tt.wantErr {
				t.Errorf("TurnOn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				// Calculate which byte should contain our bit
				position := (tt.coord.rowIndex * g.width) + tt.coord.columnIndex
				byteIndex := position / 8

				// Verify the bit was set correctly
				if g.data[byteIndex] != tt.expectedBit {
					t.Errorf("TurnOn() got byte = %08b, want %08b", g.data[byteIndex], tt.expectedBit)
				}
			}
		})
	}
}
