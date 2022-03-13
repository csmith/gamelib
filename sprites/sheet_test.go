package sprites

import (
	"bytes"
	"fmt"
	"image/png"
	"os"
	"path/filepath"
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sebdah/goldie/v2"
)

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

			_ = os.MkdirAll("../testoutput/", 0755)
			_ = os.WriteFile(fmt.Sprintf("../testoutput/sheet-%s.png", files[i]), writer.Bytes(), 0644)

			gold.Assert(t, "sheet", writer.Bytes())
		})
	}
}
