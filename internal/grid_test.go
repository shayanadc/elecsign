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

func TestGrid_IsOn(t *testing.T) {
	tests := []struct {
		name    string
		setup   func(*Grid)
		coord   Coordinate
		want    bool
		wantErr bool
	}{
		{
			name: "pixel is on",
			setup: func(g *Grid) {
				g.TurnOn(Coordinate{rowIndex: 2, columnIndex: 3})
			},
			coord:   Coordinate{rowIndex: 2, columnIndex: 3},
			want:    true,
			wantErr: false,
		},
		{
			name:    "pixel is off",
			setup:   func(g *Grid) {},
			coord:   Coordinate{rowIndex: 2, columnIndex: 3},
			want:    false,
			wantErr: false,
		},
		{
			name:    "invalid coordinate negative row",
			setup:   func(g *Grid) {},
			coord:   Coordinate{rowIndex: -1, columnIndex: 0},
			want:    false,
			wantErr: true,
		},
		{
			name:    "invalid coordinate column too large",
			setup:   func(g *Grid) {},
			coord:   Coordinate{rowIndex: 0, columnIndex: 36},
			want:    false,
			wantErr: true,
		},
		{
			name: "multiple pixels on",
			setup: func(g *Grid) {
				g.TurnOn(Coordinate{rowIndex: 0, columnIndex: 0})
				g.TurnOn(Coordinate{rowIndex: 0, columnIndex: 1})
			},
			coord:   Coordinate{rowIndex: 0, columnIndex: 1},
			want:    true,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGrid(6, 36)
			tt.setup(g)

			got, err := g.IsOn(tt.coord)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsOn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsOn() = %v, want %v", got, tt.want)
			}
		})
	}
}
