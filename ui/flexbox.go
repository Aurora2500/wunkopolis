package ui

type FlexDirection int

const (
	Row FlexDirection = iota
	Col
)

type FlexAnchor int

const (
	Start FlexAnchor = iota
	Center
)

type Flexbox struct {
	UIBase
	Elements            []UIElem
	Dir                 FlexDirection
	Anchor              FlexAnchor
	Padding             float32
	VisibilityCondition func() bool
}

func (fb *Flexbox) Layout(area Area) {
	fb.RealSize = area

	elemoffset := Vector2{X: area.X, Y: area.Y}

	if fb.Anchor == Center {
		size := Vector2{}
		for _, elem := range fb.Elements {
			size.X += elem.GetSize().Width + fb.Padding
			size.Y += elem.GetSize().Height + fb.Padding
		}
		if fb.Dir == Row {
			elemoffset.X = elemoffset.X + area.Width/2 - size.X/2
		} else {
			elemoffset.Y = elemoffset.Y + area.Height/2 - size.Y/2
		}
	}

	for _, elem := range fb.Elements {
		if fb.Dir == Row {
			elem.Layout(Area{
				Width:  elem.GetSize().Width,
				Height: area.Height,
				X:      elemoffset.X,
				Y:      elemoffset.Y,
			})
		} else {
			elem.Layout(Area{
				Width:  area.Width,
				Height: elem.GetSize().Height,
				X:      elemoffset.X,
				Y:      elemoffset.Y,
			})
		}

		if fb.Dir == Row {
			elemoffset.X = elemoffset.X + elem.GetSize().Width + fb.Padding
		} else {
			elemoffset.Y = elemoffset.Y + elem.GetSize().Height + fb.Padding
		}

	}
}

func (fb *Flexbox) Update() {
	if fb.VisibilityCondition != nil && !fb.VisibilityCondition() {
		return
	}
	for _, elem := range fb.Elements {
		elem.Update()
	}
}

func (fb *Flexbox) Draw(ctx *Context) {
	if fb.VisibilityCondition != nil && !fb.VisibilityCondition() {
		return
	}
	for _, elem := range fb.Elements {
		elem.Draw(ctx)
	}
}

func (fb *Flexbox) GetSize() Area {
	return fb.RealSize
}
