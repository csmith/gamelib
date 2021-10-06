package sprites

import (
	"bytes"
	"image/png"
	"os"
	"path/filepath"
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sebdah/goldie/v2"
)

func TestTextRenderer_Render_CaseInsensitive(t *testing.T) {
	gold := goldie.New(t)

	f, err := os.Open(filepath.Join("testdata", "text-singlecase.png"))
	if err != nil {
		t.Errorf("Unable to open test data: %v", err)
		return
	}

	im, err := png.Decode(f)
	if err != nil {
		t.Errorf("Unable to decode test data: %v", err)
		return
	}

	sheet := NewSheet(ebiten.NewImageFromImage(im), 3, 5)
	renderer := NewTextRenderer(sheet, "aBcD†", false)

	out := ebiten.NewImage(33, 14)
	renderer.Render(out, 1, 1, "aAbBcCdD\n†† ?! ††")

	writer := &bytes.Buffer{}
	if err := png.Encode(writer, out); err != nil {
		t.Errorf("Unable to encode test image: %v", err)
		return
	}

	gold.Assert(t, "textrender-case-insensitive", writer.Bytes())
}

func TestTextRenderer_Render_CaseSensitive(t *testing.T) {
	gold := goldie.New(t)

	f, err := os.Open(filepath.Join("testdata", "text-mixedcase.png"))
	if err != nil {
		t.Errorf("Unable to open test data: %v", err)
		return
	}

	im, err := png.Decode(f)
	if err != nil {
		t.Errorf("Unable to decode test data: %v", err)
		return
	}

	sheet := NewSheet(ebiten.NewImageFromImage(im), 3, 5)
	renderer := NewTextRenderer(sheet, "AaBb†", true)

	out := ebiten.NewImage(33, 14)
	renderer.Render(out, 1, 1, "aAbBaabb\n†† ?! ††")

	writer := &bytes.Buffer{}
	if err := png.Encode(writer, out); err != nil {
		t.Errorf("Unable to encode test image: %v", err)
		return
	}

	gold.Assert(t, "textrender-case-sensitive", writer.Bytes())
}

func TestTextRenderer_RenderWrapped(t *testing.T) {
	gold := goldie.New(t)

	f, err := os.Open(filepath.Join("testdata", "text-singlecase.png"))
	if err != nil {
		t.Errorf("Unable to open test data: %v", err)
		return
	}

	im, err := png.Decode(f)
	if err != nil {
		t.Errorf("Unable to decode test data: %v", err)
		return
	}

	sheet := NewSheet(ebiten.NewImageFromImage(im), 3, 5)
	renderer := NewTextRenderer(sheet, "aBcD†", false)

	out := ebiten.NewImage(33, 126)
	renderer.RenderWrapped(out, 1, 1, 32, "aAbBaabb\n\n†† ?! ††\n\naa bb aa bb aa bbbb cc dd aaaa\n\n†††††††††††abcd\n\n††††††††††† abcd\n\nabcd †††††††††††")

	writer := &bytes.Buffer{}
	if err := png.Encode(writer, out); err != nil {
		t.Errorf("Unable to encode test image: %v", err)
		return
	}

	gold.Assert(t, "textrender-wrapped", writer.Bytes())
}
