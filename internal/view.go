package internal

type ViewManager interface {
	IsOn(c Coordinate) bool
	Get() Grid
	TurnOn([]Coordinate)
}

type View struct {
	data Grid
}

// NewView creates a new View with a fixed-size Grid
func NewView(width, height int) *View {
	return &View{
		data: *NewGrid(height, width),
	}
}

func (v *View) Get() Grid {
	return v.data
}

func (v *View) IsOn(c Coordinate) bool {
	on, err := v.data.IsOn(c)
	if err != nil {
		return false
	}

	return on
}

// TurnOn turns on multiple pixels at the given coordinates
func (v *View) TurnOn(coords []Coordinate) {
	for _, coord := range coords {
		_ = v.data.TurnOn(coord)
	}
}
