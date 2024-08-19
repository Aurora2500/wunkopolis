package ui

import (
	"log"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Alignment int

const (
	Left Alignment = iota
	Right
	Centered
	Justified
)

type Text struct {
	UIBase
	Content string
	Font    rl.Font
	Size    float32
	lines   []string
}

func (t *Text) Layout(area Area) {
	font := t.Font
	if !rl.IsFontReady(font) {
		font = rl.GetFontDefault()
	}
	size := t.Size
	if size == 0. {
		size = float32(font.BaseSize)
	}
	spacing := size / float32(font.BaseSize)

	widthLimit := area.Width
	spaceWidth := rl.MeasureTextEx(t.Font, " ", size, spacing).X

	for _, oLine := range strings.Split(t.Content, "\n") {
		realLine := ""
		var widthCurr float32

		for _, word := range strings.Split(oLine, " ") {
			wordSize := rl.MeasureTextEx(t.Font, word, size, spacing)
			if wordSize.X > widthLimit {
				// TODO: handle this properly, maybe
				log.Fatalln("Encountered word longer than allowed line width")
			}
			extraWidth := wordSize.X
			if len(realLine) > 0 {
				extraWidth += spaceWidth
			}
			if widthCurr+extraWidth > widthLimit {
				t.lines = append(t.lines, realLine)
				realLine = word
				widthCurr = wordSize.X
				continue
			}
			if len(realLine) > 0 {
				realLine += " "
			}
			realLine += word
			widthCurr += extraWidth
		}
		if len(realLine) > 0 {
			t.lines = append(t.lines, realLine)
		}
	}

	totalHeight := float32(len(t.lines)) * size
	if totalHeight > area.Height {
		log.Fatalln("Text height greater than allowed height")
	}
	t.RealSize = area
}

func (t *Text) Draw(ctx *Context) {
	font := t.Font
	if !rl.IsFontReady(font) {
		font = rl.GetFontDefault()
	}
	size := t.Size
	if size == 0. {
		size = float32(font.BaseSize)
	}
	spacing := size / float32(font.BaseSize)
	rl.DrawTextEx(font, strings.Join(t.lines, "\n"), rl.Vector2{X: t.RealSize.X, Y: t.RealSize.Y}, size, spacing, rl.Black)
}
