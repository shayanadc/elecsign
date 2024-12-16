package display

import (
	"elecsign/internal/grid"
	"elecsign/internal/view"
)

// Renderer interface for rendering a grid
type Renderer interface {
	Render(grid grid.Grid)
}

// Display interface for managing views
type Display interface {
	Show()
	AddView(v view.View)
	Clear()
}
