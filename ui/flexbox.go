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
	Elements   []UIElem
	Direction  FlexDirection
	Anchor     FlexAnchor
	Padding    float32
	Border     float32
	Background NPatchBox
}

func (fb *Flexbox) Layout(area Area) {
	fb.RealSize = area
	fb.Background.Layout(area)
	fb.LayoutInside()
}

func (fb *Flexbox) LayoutInside() {
	area := fb.RealSize
	elemoffset := Vector2{X: area.X + fb.Border, Y: area.Y + fb.Border}

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
	fb.Background.Draw(ctx)
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
	Background     NPatchBox
	ScrollSize     float32
	Border         float32
}

func (sb *Scrollbox) Layout(area Area) {
	sb.RealSize = area
	sb.Background.Layout(area)
	sb.UpdateSize()
	sb.LayoutInside()
}

func (sb *Scrollbox) LayoutInside() {
	elemoffset := Vector2{
		X: sb.RealSize.X + sb.Border,
		Y: sb.RealSize.Y + sb.Border,
	}
	if sb.Direction == Row {
		elemoffset.X -= sb.ScrollPosition
	} else {
		elemoffset.Y -= sb.ScrollPosition
	}
	for _, elem := range sb.Elements {

		if sb.Direction == Row {
			elem.Layout(Area{
				Width:  elem.GetSize().Width,
				Height: sb.RealSize.Height,
				X:      elemoffset.X,
				Y:      elemoffset.Y,
			})
			elemoffset.X = elemoffset.X + elem.GetSize().Width + sb.Padding
		} else {
			elem.Layout(Area{
				Width:  sb.RealSize.Width,
				Height: elem.GetSize().Height,
				X:      elemoffset.X,
				Y:      elemoffset.Y,
			})
			elemoffset.Y = elemoffset.Y + elem.GetSize().Height + sb.Padding
		}
	}
}

func (sb *Scrollbox) Update() {

	dscroll := rl.GetMouseWheelMove()
	if dscroll != 0. {
		sb.ScrollPosition = rl.Clamp(sb.ScrollPosition-dscroll*24, 0, sb.ScrollSize)
		sb.LayoutInside()
	}

	for _, elem := range sb.Elements {
		elem.Update()
	}
}

func (sb *Scrollbox) Draw(ctx *Context) {
	sb.Background.Draw(ctx)
	ctx.PushScissor(InsetArea(sb.RealSize, sb.Border))
	for _, elem := range sb.Elements {
		elem.Draw(ctx)
	}
	ctx.PopScissor()
}

func (sb *Scrollbox) GetSize() Area {
	return sb.RealSize
}

func (sb *Scrollbox) Add(elem UIElem) {
	sb.Elements = append(sb.Elements, elem)
	sb.UpdateSize()
}

func (sb *Scrollbox) UpdateSize() {
	scrollSize := float32(0)

	if sb.Direction == Row {
		scrollSize = scrollSize - InsetArea(sb.RealSize, sb.Border).Width - sb.Padding
	} else {
		scrollSize = scrollSize - InsetArea(sb.RealSize, sb.Border).Height - sb.Padding

	}
	for _, elem := range sb.Elements {
		if sb.Direction == Row {
			scrollSize = scrollSize + elem.GetSize().Width + sb.Padding
		} else {
			scrollSize = scrollSize + elem.GetSize().Height + sb.Padding
		}
	}
	sb.ScrollSize = scrollSize
}

type MapBox struct {
	UIBase
	Map        Maper
	Background NPatchBox
}

func (sb *MapBox) Layout(area Area) {
	sb.RealSize = area
	sb.Background.Layout(area)
}

func (sb *MapBox) Draw(ctx *Context) {
	sb.Background.Draw(ctx)
	sb.Map.DrawMap(sb.RealSize)
}

func (sb *MapBox) GetSize() Area {
	return sb.RealSize
}

func (sb *MapBox) Update() {

}
