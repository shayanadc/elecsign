package internal

// Renderer interface for rendering a grid
type Renderer interface {
	Render(grid Grid)
}

// Display interface for managing views
type Display interface {
	Show()
	AddView(v View)
	Clear()
}
