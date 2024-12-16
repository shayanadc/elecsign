package internal

import (
	"reflect"
	"testing"
)

func TestViews_Get(t *testing.T) {
	tests := []struct {
		name     string
		setup    func(*View)
		wantGrid Grid
	}{
		{
			name:  "empty grid",
			setup: func(v *View) {},
			wantGrid: Grid{
				data:   [27]byte{},
				width:  36,
				height: 6,
			},
		},
		{
			name: "grid with one pixel on",
			setup: func(v *View) {
				v.data.TurnOn(Coordinate{rowIndex: 0, columnIndex: 0})
			},
			wantGrid: Grid{
				data: func() [27]byte {
					var data [27]byte
					data[0] = 0b10000000 // First bit set in first byte
					return data
				}(),
				width:  36,
				height: 6,
			},
		},
		{
			name: "grid with multiple pixels on",
			setup: func(v *View) {
				v.data.TurnOn(Coordinate{rowIndex: 0, columnIndex: 0})
				v.data.TurnOn(Coordinate{rowIndex: 0, columnIndex: 1})
			},
			wantGrid: Grid{
				data: func() [27]byte {
					var data [27]byte
					data[0] = 0b11000000 // First two bits set in first byte
					return data
				}(),
				width:  36,
				height: 6,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &View{
				data: *NewGrid(6, 36),
			}
			tt.setup(v)

			got := v.Get()
			if !reflect.DeepEqual(got, tt.wantGrid) {
				t.Errorf("Get() = %v, want %v", got, tt.wantGrid)
			}
		})
	}
}
