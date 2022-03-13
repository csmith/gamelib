package text

import (
	"errors"
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
)

var errDone = errors.New("done")

type testGame struct {
	m *testing.M
}

func (t *testGame) Update() error {
	t.m.Run()
	return errDone
}

func (t *testGame) Draw(screen *ebiten.Image) {
	// Ignore
}

func (t *testGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1, 1
}

func TestMain(m *testing.M) {
	if err := ebiten.RunGame(&testGame{m}); err != errDone {
		panic(err)
	}
}
