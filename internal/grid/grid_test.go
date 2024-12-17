package grid

import (
	"testing"

	"elecsign/internal/transformer"
)

func TestGrid_TurnOn(t *testing.T) {
	tests := []struct {
		name        string
		coord       transformer.Coordinate
		wantErr     bool
		expectedBit byte
	}{
		{
			name:        "valid coordinate middle of display",
			coord:       transformer.Coordinate{RowIndex: 2, ColumnIndex: 3},
			wantErr:     false,
			expectedBit: 0b00010000, // bit 4 set in the appropriate byte
		},
		{
			name:        "valid coordinate first position",
			coord:       transformer.Coordinate{RowIndex: 0, ColumnIndex: 0},
			wantErr:     false,
			expectedBit: 0b10000000, // leftmost bit set
		},
		{
			name:    "invalid coordinate negative row",
			coord:   transformer.Coordinate{RowIndex: -1, ColumnIndex: 0},
			wantErr: true,
		},
		{
			name:    "invalid coordinate negative column",
			coord:   transformer.Coordinate{RowIndex: 0, ColumnIndex: -1},
			wantErr: true,
		},
		{
			name:    "invalid coordinate row too large",
			coord:   transformer.Coordinate{RowIndex: 6, ColumnIndex: 0},
			wantErr: true,
		},
		{
			name:    "invalid coordinate column too large",
			coord:   transformer.Coordinate{RowIndex: 0, ColumnIndex: 36},
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
				position := (tt.coord.RowIndex * g.Width) + tt.coord.ColumnIndex
				byteIndex := position / 8

				if g.Data[byteIndex] != tt.expectedBit {
					t.Errorf("TurnOn() got byte = %08b, want %08b", g.Data[byteIndex], tt.expectedBit)
				}
			}
		})
	}
}

func TestGrid_IsOn(t *testing.T) {
	tests := []struct {
		name    string
		setup   func(*Grid)
		coord   transformer.Coordinate
		want    bool
		wantErr bool
	}{
		{
			name: "pixel is on",
			setup: func(g *Grid) {
				_ = g.TurnOn(transformer.Coordinate{RowIndex: 2, ColumnIndex: 3})
			},
			coord:   transformer.Coordinate{RowIndex: 2, ColumnIndex: 3},
			want:    true,
			wantErr: false,
		},
		{
			name:    "pixel is off",
			setup:   func(g *Grid) {},
			coord:   transformer.Coordinate{RowIndex: 2, ColumnIndex: 3},
			want:    false,
			wantErr: false,
		},
		{
			name:    "invalid coordinate negative row",
			setup:   func(g *Grid) {},
			coord:   transformer.Coordinate{RowIndex: -1, ColumnIndex: 0},
			want:    false,
			wantErr: true,
		},
		{
			name:    "invalid coordinate column too large",
			setup:   func(g *Grid) {},
			coord:   transformer.Coordinate{RowIndex: 0, ColumnIndex: 36},
			want:    false,
			wantErr: true,
		},
		{
			name: "multiple pixels on",
			setup: func(g *Grid) {
				_ = g.TurnOn(transformer.Coordinate{RowIndex: 0, ColumnIndex: 0})
				_ = g.TurnOn(transformer.Coordinate{RowIndex: 0, ColumnIndex: 1})
			},
			coord:   transformer.Coordinate{RowIndex: 0, ColumnIndex: 1},
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
