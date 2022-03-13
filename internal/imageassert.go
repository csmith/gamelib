package internal

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	"image/png"
	"os"
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sebdah/goldie/v2"
)

func AssertImageMatches(t *testing.T, image *ebiten.Image, goldenName string, maxDifference float64) {
	writer := &bytes.Buffer{}
	if err := png.Encode(writer, image); err != nil {
		t.Errorf("Unable to encode test image: %v", err)
		return
	}

	gold := goldie.New(t)

	// Goldie defines a flag but has no way to expose it, so munge it out of the default flag set :/
	if flag.CommandLine.Lookup("update").Value.String() == "true" {
		err := gold.Update(t, goldenName, writer.Bytes())
		if err != nil {
			t.Errorf("Failed to update golden image: %v", err)
			return
		}
		return
	}

	f, err := os.Open(gold.GoldenFileName(t, goldenName))
	if err != nil {
		t.Errorf("Failed to read golden image: %v", err)
		return
	}
	defer f.Close()

	golden, err := png.Decode(f)
	if err != nil {
		t.Errorf("Failed to decode golden image: %v", err)
		return
	}

	_ = os.MkdirAll("../testoutput/", 0755)
	_ = os.WriteFile(fmt.Sprintf("../testoutput/%s.png", t.Name()), writer.Bytes(), 0644)

	compare(t, golden, image, maxDifference)
}

func compare(t *testing.T, src, dst image.Image, difference float64) {
	srcBounds := src.Bounds()
	dstBounds := dst.Bounds()

	if srcBounds.Min.X != dstBounds.Min.X || srcBounds.Min.Y != dstBounds.Min.Y || srcBounds.Max.X != dstBounds.Max.X || srcBounds.Max.Y != dstBounds.Max.Y {
		t.Errorf("Image bounds do not match (%v != %v)", srcBounds, dstBounds)
		return
	}

	var differentPixels float64
	for y := srcBounds.Min.Y; y < srcBounds.Max.Y; y++ {
		for x := srcBounds.Min.X; x < srcBounds.Max.X; x++ {
			r1, g1, b1, a1 := src.At(x, y).RGBA()
			r2, g2, b2, a2 := dst.At(x, y).RGBA()
			if r1 != r2 || g1 != g2 || b1 != b2 || a1 != a2 {
				differentPixels++
			}
		}
	}

	diffPercent := differentPixels / float64(srcBounds.Max.X*srcBounds.Max.Y)
	if diffPercent > difference {
		t.Errorf("Image differs by %v%% (%v pixels different)", diffPercent*100, differentPixels)
	}
}
