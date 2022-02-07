package structure

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Scene describes a single "scene" or screen in the game, such as a particular menu or level.
type Scene interface {
	// Update is called every frame to update the state of the game.
	Update() (Scene, error)

	// Draw is called every frame to render the game to screen.
	Draw(screen *ebiten.Image)
}
