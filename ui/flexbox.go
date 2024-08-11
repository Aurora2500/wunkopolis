package ui

type FlexDirection int

const (
	Row FlexDirection = iota
	Col
)

type Flexbox struct {
	UIBase
	Elements []UIElem
	Dir      FlexDirection
}

func (fb *Flexbox) Layout(area Area) {
	fb.RealSize = area
	numElements := float32(len(fb.Elements))
	elemArea := area
	if fb.Dir == Row {
		elemArea.Width = area.Width / numElements
	} else {
		elemArea.Height = area.Height / float32(numElements)
	}

	for _, elem := range fb.Elements {
		elem.Layout(elemArea)
		if fb.Dir == Row {
			elemArea.X = elemArea.X + elemArea.Width
		} else {
			elemArea.Y = elemArea.Y + elemArea.Height
		}
	}
}

func (fb *Flexbox) Draw(ctx *Context) {
	for _, elem := range fb.Elements {
		elem.Draw(ctx)
	}
}
