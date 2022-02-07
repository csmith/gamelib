package structure

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// SceneGame is an Ebiten Game implementation that defers all drawing and updating logic to a "Scene".
//
// The Scene's Update() method returns the next Scene to be rendered, allowing the game to transition between
// different views based on the current state.
type SceneGame struct {
	Scene Scene
}

func (s *SceneGame) Update() (err error) {
	s.Scene, err = s.Scene.Update()
	return
}

func (s *SceneGame) Draw(screen *ebiten.Image) {
	s.Scene.Draw(screen)
}

func (s *SceneGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

// Check we've satisfied the Game interface.
var _ ebiten.Game = (*SceneGame)(nil)
