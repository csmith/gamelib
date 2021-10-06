package sprites

import (
	"math"
	"strings"
	"unicode"

	"github.com/hajimehoshi/ebiten/v2"
)

// NewTextRenderer creates a new TextRenderer that will use sprites from the given sheets corresponding to the
// set of characters.  If caseSensitive is set to false then case is ignored when searching for a matching sprite.
func NewTextRenderer(sheet *Sheet, chars string, caseSensitive bool) *TextRenderer {
	if !caseSensitive {
		chars = strings.ToLower(chars)
	}

	width, height := sheet.Size()
	return &TextRenderer{
		sheet:         sheet,
		chars:         chars,
		caseSensitive: caseSensitive,
		lineHeight:    float64(height + 2),
		letterWidth:   float64(width + 1),
		op:            &ebiten.DrawImageOptions{},
	}
}

// TextRenderer provides functions for rendering blocks of text from a sprite-sheet based font.
type TextRenderer struct {
	sheet         *Sheet
	chars         string
	caseSensitive bool
	lineHeight    float64
	letterWidth   float64
	op            *ebiten.DrawImageOptions
}

// Render renders the given text line by line starting at the given co-ordinates.
func (t *TextRenderer) Render(image *ebiten.Image, x, y float64, text string) {
	op := t.op
	op.GeoM.Reset()
	op.GeoM.Translate(x, y)
	line := 0.0
	runes := []rune(text)
	for i := range runes {
		if runes[i] == '\n' {
			line++
			op.GeoM.Reset()
			op.GeoM.Translate(x, y+line*t.lineHeight)
			continue
		}
		image.DrawImage(t.spriteFor(runes[i]), op)
		op.GeoM.Translate(t.letterWidth, 0)
	}
}

// RenderWrapped renders the given text onto the image, starting at the given co-ordinates. Wrapping is done at the
// nearest space character; if a single word is longer than the maximum width then it will be cut mid-word.
func (t *TextRenderer) RenderWrapped(image *ebiten.Image, x, y float64, width float64, text string) {
	maxLineLength := int(math.Floor(width / t.letterWidth))
	var lines [][]rune
	sourceLines := strings.Split(text, "\n")
	for i := range sourceLines {
		line := []rune(sourceLines[i])
		for len(line) > maxLineLength {
			space := -1
			for j := maxLineLength; j > 0; j-- {
				if line[j] == ' ' {
					space = j
					break
				}
			}
			if space == -1 {
				lines = append(lines, line[:maxLineLength])
				line = line[maxLineLength:]
			} else {
				lines = append(lines, line[:space])
				line = line[space+1:]
			}
		}

		if len(line) > 0 || len(sourceLines[i]) == 0 {
			lines = append(lines, line)
		}
	}

	op := t.op
	op.GeoM.Reset()
	op.GeoM.Translate(x, y)
	line := 0.0
	for l := range lines {
		for i := range lines[l] {
			image.DrawImage(t.spriteFor(lines[l][i]), op)
			op.GeoM.Translate(t.letterWidth, 0)
		}
		line++
		op.GeoM.Reset()
		op.GeoM.Translate(x, y+line*t.lineHeight)
	}
}

func (t *TextRenderer) spriteFor(rune rune) *ebiten.Image {
	if !t.caseSensitive {
		rune = unicode.ToLower(rune)
	}
	return t.sheet.Sprite(strings.IndexRune(t.chars, rune))
}
