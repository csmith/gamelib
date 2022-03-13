package text

import (
	"bytes"
	"image/png"
	"os"
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sebdah/goldie/v2"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var impact font.Face

func init() {
	b, _ := os.ReadFile("testdata/impact.ttf")
	f, _ := opentype.Parse(b)
	impact, _ = opentype.NewFace(f, &opentype.FaceOptions{
		Size:    48,
		DPI:     72,
		Hinting: font.HintingFull,
	})
}

func TestShadowed(t *testing.T) {
	gold := goldie.New(t)

	out := ebiten.NewImage(280, 75)
	Shadowed(out, "Hello", impact, 10, 60, colornames.White, colornames.Black, 3)
	Shadowed(out, "World", impact, 125, 60, colornames.Blue, colornames.Pink, 10)
	Shadowed(out, "!", impact, 260, 60, colornames.Orange, colornames.Aquamarine, 0)

	writer := &bytes.Buffer{}
	if err := png.Encode(writer, out); err != nil {
		t.Errorf("Unable to encode test image: %v", err)
		return
	}

	gold.Assert(t, "shadowed", writer.Bytes())
}
