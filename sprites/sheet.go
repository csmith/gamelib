package sprites

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

// NewSheet creates a new sprite sheet from the given image, with fixed width and height sprites.
func NewSheet(source *ebiten.Image, width, height int) *Sheet {
	return &Sheet{
		im:     source,
		width:  width,
		height: height,
		cols:   source.Bounds().Dx() / width,
	}
}

// Sheet contains the data and meta-data for a sprite sheet.
type Sheet struct {
	im     *ebiten.Image
	width  int
	height int
	cols   int
}

// Sprite retrieves the sprite at the given index within the sheet.
func (s *Sheet) Sprite(index int) *ebiten.Image {
	x := (index % s.cols) * s.width
	y := (index / s.cols) * s.height
	return s.im.SubImage(image.Rect(x, y, x+s.width, y+s.height)).(*ebiten.Image)
}

// Size returns the width and height of a single sprite in this sheet, in pixels.
func (s *Sheet) Size() (width, height int) {
	return s.width, s.height
}
