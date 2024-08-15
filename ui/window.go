package ui

import (
	"wunkopolis/assets"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const topBarSize = 56
const topBarOffset = 16
const borderSize = 10

var topBarColor = rl.Color{
	R: 17,
	G: 24,
	B: 136,
	A: 255,
}

type Window struct {
	Content    UIElem
	Title      string
	Area       Area
	dragging   bool
	hidden     bool
	background NPatchBox
	button     Button
}

func (w *Window) barArea() Area {
	return Area{
		X:      w.Area.X + topBarOffset,
		Y:      w.Area.Y + topBarOffset,
		Width:  w.Area.Width - 32,
		Height: topBarSize,
	}
}

func (w *Window) Setup() {
	bgtex := assets.Manager.GetTexture("Panel")

	w.background = NPatchBox{
		Texture: bgtex,
		NPatchInfo: rl.NPatchInfo{
			Top:    8,
			Right:  8,
			Bottom: 8,
			Left:   8,
			Source: Area{
				Width:  float32(bgtex.Width),
				Height: float32(bgtex.Height),
			},
		},
	}
	w.button = Button{
		base:    assets.Manager.GetTexture("Button"),
		hover:   assets.Manager.GetTexture("ButtonHover"),
		pressed: assets.Manager.GetTexture("ButtonPressed"),
		icon:    assets.Manager.GetTexture("x"), onClick: func() { w.Hide() }}
}

func (w *Window) Update() {
	if w.hidden {
		return
	}
	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) && rl.CheckCollisionPointRec(rl.GetMousePosition(), w.Area) {
		w.dragging = true
		rl.SetMouseCursor(rl.MouseCursorResizeAll)
		return
	}

	if w.dragging && rl.IsMouseButtonReleased(rl.MouseButtonLeft) {
		w.dragging = false
		rl.SetMouseCursor(rl.MouseCursorArrow)
		return
	}

	if w.dragging {
		delta := rl.GetMouseDelta()

		w.Area.X = w.Area.X + delta.X
		w.Area.Y = w.Area.Y + delta.Y
	}
}

func (w *Window) Draw() {
	if w.hidden {
		return
	}
	contentArea := Area{
		X:      w.Area.X + borderSize,
		Y:      w.Area.Y + topBarSize + topBarOffset,
		Width:  w.Area.Width - 2*borderSize,
		Height: w.Area.Height - borderSize - topBarSize - topBarOffset,
	}
	w.background.Layout(w.Area)
	w.Content.Layout(contentArea)
	w.button.Layout(rl.Rectangle{Width: 48, Height: 43, Y: w.Area.Y + 23, X: w.Area.X + w.Area.Width - 72})

	ctx := Context{}
	w.background.Draw(&ctx)
	w.Content.Draw(&ctx)
	rl.DrawRectangleRec(w.barArea(), topBarColor)
	w.button.Draw(&ctx)
	w.button.Check()
}

func (w *Window) Hide() {
	w.hidden = true
}
