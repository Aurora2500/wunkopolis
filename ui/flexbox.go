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
	Padding  float32
}

func (fb *Flexbox) Layout(area Area) {
	fb.RealSize = area

	elemoffset := Vector2{X: area.X, Y: area.Y}

	for _, elem := range fb.Elements {
		elem.Layout(Area{
			Width:  elem.GetSize().Width,
			Height: elem.GetSize().Height,
			X:      elemoffset.X,
			Y:      elemoffset.Y,
		})
		if fb.Dir == Row {
			elemoffset.X = elemoffset.X + elem.GetSize().Width + fb.Padding
		} else {
			elemoffset.Y = elemoffset.Y + elem.GetSize().Height + fb.Padding
		}

	}
}

func (fb *Flexbox) Update() {
	for _, elem := range fb.Elements {
		elem.Update()
	}
}

func (fb *Flexbox) Draw(ctx *Context) {
	for _, elem := range fb.Elements {
		elem.Draw(ctx)
	}
}

func (fb *Flexbox) GetSize() Area {
	return fb.RealSize
}
