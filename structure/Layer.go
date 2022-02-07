package structure

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Layer represents a visually distinct layer of the UI that can be rendered independently of the rest.
type Layer interface {
	// Draw is called every frame to render the game to screen.
	Draw(screen *ebiten.Image)
}

// Layers is a collection of Layers, which can be rendered in order.
type Layers []Layer

// Draw invokes the Draw method of each Layer in the collection in order.
func (l Layers) Draw(screen *ebiten.Image) {
	for i := range l {
		l[i].Draw(screen)
	}
}
