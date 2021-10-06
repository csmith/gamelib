package sprites

import (
	"bytes"
	"errors"
	"fmt"
	"image/png"
	"os"
	"path/filepath"
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sebdah/goldie/v2"
)

var errDone = errors.New("done")

type TestGame struct {
	m *testing.M
}

func (t *TestGame) Update() error {
	t.m.Run()
	return errDone
}

func (t *TestGame) Draw(screen *ebiten.Image) {
	// Ignore
}

func (t *TestGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1, 1
}

func TestMain(m *testing.M) {
	if err := ebiten.RunGame(&TestGame{m}); err != errDone {
		panic(err)
	}
}

func TestSheet_Sprite(t *testing.T) {
	gold := goldie.New(t)

	files := []string{"grid", "horizontal", "vertical"}
	for i := range files {
		t.Run(files[i], func(t *testing.T) {
			f, err := os.Open(filepath.Join("testdata", fmt.Sprintf("sheet-%s.png", files[i])))
			if err != nil {
				t.Errorf("Unable to open test data: %v", err)
				return
			}

			im, err := png.Decode(f)
			if err != nil {
				t.Errorf("Unable to decode test data: %v", err)
				return
			}

			sheet := NewSheet(ebiten.NewImageFromImage(im), 16, 32)

			// Draw a new image with one of each sprite in a certain place
			out := ebiten.NewImage(64, 64)
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(5, 6)
			out.DrawImage(sheet.Sprite(0), op)
			op.GeoM.Translate(10, 20)
			out.DrawImage(sheet.Sprite(1), op)
			op.GeoM.Translate(10, -20)
			out.DrawImage(sheet.Sprite(2), op)
			op.GeoM.Translate(10, 20)
			out.DrawImage(sheet.Sprite(3), op)
			op.GeoM.Translate(10, -20)
			out.DrawImage(sheet.Sprite(4), op)

			writer := &bytes.Buffer{}
			if err := png.Encode(writer, out); err != nil {
				t.Errorf("Unable to encode test image: %v", err)
				return
			}

			gold.Assert(t, "sheet", writer.Bytes())
		})
	}
}

