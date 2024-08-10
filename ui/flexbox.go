package ui

import rl "github.com/gen2brain/raylib-go/raylib"

type FlexDirection int

const (
	Row FlexDirection = iota
	Col
)

type Flexbox struct {
	Elements []UIElem
	Dir      FlexDirection
}

func (fb *Flexbox) Draw(area rl.Rectangle, ctx *Context) {
	numElements := float32(len(fb.Elements))
	elemArea := area
	if fb.Dir == Row {
		elemArea.Width = area.Width / numElements
	} else {
		elemArea.Height = area.Height / float32(numElements)
	}

	for _, elem := range fb.Elements {
		elem.Draw(elemArea, ctx)
		if fb.Dir == Row {
			elemArea.X = elemArea.X + elemArea.Width
		} else {
			elemArea.Y = elemArea.Y + elemArea.Height
		}
	}
}
