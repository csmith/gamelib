package math

import (
	"bytes"
	"github.com/sebdah/goldie/v2"
	"golang.org/x/image/colornames"
	"image"
	"image/png"
	"math"
	"testing"
)

func TestEasingFunctionsVisual(t *testing.T) {
	gold := goldie.New(t)

	funcs := map[string]func(float64) float64{
		"linear":          func(p float64) float64 { return Lerp(0, 1, p) },
		"ease-in":         EaseIn,
		"ease-out":        EaseOut,
		"ease-in-out":     EaseInOut,
		"ease-out-bounce": EaseOutBounce,
	}

	render := func(f func(float64) float64) ([]byte, error) {
		im := image.NewRGBA(image.Rect(0, 0, 140, 140))
		for x := 0; x < 140; x++ {
			for y := 0; y < 140; y++ {
				if x == 20 || x == 120 || y == 20 || y == 120 {
					im.Set(x, y, colornames.Gray)
				} else {
					im.Set(x, y, colornames.White)
				}
			}
		}

		for x := -0.2; x < 1.2; x += 0.001 {
			y := int(math.Round(120 - 100*f(x)))
			im.Set(int(20+100*x), y, colornames.Black)
		}
		writer := &bytes.Buffer{}
		if err := png.Encode(writer, im); err != nil {
			return nil, err
		}
		return writer.Bytes(), nil
	}

	for f := range funcs {
		t.Run(f, func(t *testing.T) {
			actual, err := render(funcs[f])
			if err != nil {
				t.Errorf("failed to render image: %v", err)
			}
			gold.Assert(t, f, actual)
		})
	}
}
