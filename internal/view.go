package internal

type View interface {
	IsOn(c Coordinate) bool
	Get() Grid
	TurnOn([]Coordinate)
}

type GridView struct {
	data Grid
}

// NewView creates a new View with a fixed-size Grid
func NewView(width, height int) *GridView {
	return &GridView{
		data: *NewGrid(height, width),
	}
}

func (v *GridView) Get() Grid {
	return v.data
}

func (v *GridView) IsOn(c Coordinate) bool {
	on, err := v.data.IsOn(c)
	if err != nil {
		return false
	}

	return on
}

// TurnOn turns on multiple pixels at the given coordinates
func (v *GridView) TurnOn(coords []Coordinate) {
	for _, coord := range coords {
		_ = v.data.TurnOn(coord)
	}
}
