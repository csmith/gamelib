package text

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	etext "github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

// Shadowed renders the given text with a shadow around the outside.
// Note that per Ebiten, the origin point is a 'dot' (period) position not the top-left corner as you might expect.
func Shadowed(dst *ebiten.Image, text string, face font.Face, x, y float64, clr color.Color, shadowClr color.Color, shadowSize float64) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)

	ShadowedWithOptions(dst, text, face, op, clr, shadowClr, shadowSize)
}

// ShadowedWithOptions renders the given text with a shadow around the outside.
func ShadowedWithOptions(dst *ebiten.Image, text string, face font.Face, op *ebiten.DrawImageOptions, clr color.Color, shadowClr color.Color, shadowSize float64) {
	cr, cg, cb, ca := clr.RGBA()
	sr, sg, sb, sa := shadowClr.RGBA()

	if shadowSize > 0 {
		op.ColorM.Scale(float64(sr)/float64(sa), float64(sg)/float64(sa), float64(sb)/float64(sa), float64(sa)/0xffff)

		for angle := 0.0; angle < (2 * math.Pi); angle += 0.35 {
			dx := math.Sin(angle) * shadowSize
			dy := math.Cos(angle) * shadowSize
			op.GeoM.Translate(dx, dy)
			etext.DrawWithOptions(dst, text, face, op)
			op.GeoM.Translate(-dx, -dy)
		}

		op.ColorM.Reset()
	}

	op.ColorM.Scale(float64(cr)/float64(ca), float64(cg)/float64(ca), float64(cb)/float64(ca), float64(ca)/0xffff)
	etext.DrawWithOptions(dst, text, face, op)
}
