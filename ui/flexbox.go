package ui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

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
	Elements  []UIElem
	Direction FlexDirection
	Anchor    FlexAnchor
	Padding   float32
}

func (fb *Flexbox) Layout(area Area) {
	fb.RealSize = area
	fb.LayoutInside()
}

func (fb *Flexbox) LayoutInside() {
	area := fb.RealSize
	elemoffset := Vector2{X: area.X, Y: area.Y}

	if fb.Anchor == Center {
		size := Vector2{}
		for _, elem := range fb.Elements {
			size.X += elem.GetSize().Width + fb.Padding
			size.Y += elem.GetSize().Height + fb.Padding
		}
		if fb.Direction == Row {
			elemoffset.X = elemoffset.X + area.Width/2 - size.X/2
		} else {
			elemoffset.Y = elemoffset.Y + area.Height/2 - size.Y/2
		}
	}

	for _, elem := range fb.Elements {
		if fb.Direction == Row {
			elem.Layout(Area{
				Width:  elem.GetSize().Width,
				Height: area.Height,
				X:      elemoffset.X,
				Y:      elemoffset.Y,
			})
			elemoffset.X = elemoffset.X + elem.GetSize().Width + fb.Padding
		} else {
			elem.Layout(Area{
				Width:  area.Width,
				Height: elem.GetSize().Height,
				X:      elemoffset.X,
				Y:      elemoffset.Y,
			})
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

func (fb *Flexbox) Add(elem UIElem) {
	fb.Elements = append(fb.Elements, elem)
	fb.LayoutInside()
}

type Scrollbox struct {
	UIBase
	Elements       []UIElem
	Direction      FlexDirection
	Padding        float32
	ScrollPosition float32
	ScrollSize     float32
}

func (sb *Scrollbox) Layout(area Area) {
	sb.RealSize = area
	sb.LayoutInside()
}

func (sb *Scrollbox) LayoutInside() {
	elemoffset := Vector2{
		X: sb.RealSize.X + sb.Padding,
		Y: sb.RealSize.Y + sb.Padding,
	}
	if sb.Direction == Row {
		elemoffset.X += sb.ScrollPosition
	} else {
		elemoffset.Y += sb.ScrollPosition
	}
	for _, elem := range sb.Elements {
		if sb.Direction == Row {
			elem.Layout(Area{
				Width:  elem.GetSize().Width,
				Height: elem.GetSize().Height,
				X:      elemoffset.X,
				Y:      elemoffset.Y,
			})
			elemoffset.X = elemoffset.X + elem.GetSize().Width + sb.Padding
		} else {
			elem.Layout(Area{
				Width:  elem.GetSize().Width,
				Height: elem.GetSize().Height,
				X:      elemoffset.X,
				Y:      elemoffset.Y,
			})
			elemoffset.Y = elemoffset.Y + elem.GetSize().Height + sb.Padding
		}
	}
	if sb.Direction == Row {
		sb.ScrollSize = -(elemoffset.X - sb.RealSize.X)
	} else {
		sb.ScrollSize = -(elemoffset.Y - sb.RealSize.Y)
	}
}

func (sb *Scrollbox) Update() {
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), sb.RealSize) {
		changed := sb.ScrollPosition + rl.GetMouseWheelMove()*24
		if changed > 0 {
			changed = 0
		}
		if changed < sb.ScrollSize {
			changed = sb.ScrollSize
		}
		if changed != sb.ScrollPosition {
			sb.ScrollPosition = changed
			sb.LayoutInside()
		}
	}

	for _, elem := range sb.Elements {
		elem.Update()
	}
}

func (sb *Scrollbox) Draw(ctx *Context) {
	rl.BeginScissorMode(sb.RealSize.ToInt32().X+int32(sb.Padding), sb.RealSize.ToInt32().Y+int32(sb.Padding), sb.RealSize.ToInt32().Width-2*int32(sb.Padding), sb.RealSize.ToInt32().Height-2*int32(sb.Padding))
	for _, elem := range sb.Elements {
		elem.Draw(ctx)
	}
	rl.EndScissorMode()
}

func (sb *Scrollbox) GetSize() Area {
	return sb.RealSize
}

func (sb *Scrollbox) Add(elem UIElem) {
	sb.Elements = append(sb.Elements, elem)
}
