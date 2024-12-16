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
