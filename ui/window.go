package ui

import (
	"wunkopolis/assets"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const topBarSize = 56
const topBarOffset = 16
const borderSize = 16

var topBarColor = rl.Color{
	R: 17,
	G: 24,
	B: 136,
	A: 255,
}

type Window struct {
	Content    UIElem
	Title      string
	Icon       rl.Texture2D
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

func (w *Window) Setup(bottomBar *Bar) {
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
		Icon:    assets.Manager.GetTexture("x"),
		OnClick: func() { w.HideShow() },
	}
	bottomBar.AddButton(Button{Icon: w.Icon, OnClick: w.HideShow, Base: assets.Manager.GetTexture("BigButton"), Hover: assets.Manager.GetTexture("BigButtonHover"), Pressed: assets.Manager.GetTexture("BigButtonPressed")})
}

func (w *Window) Update() {
	if w.hidden {
		return
	}
	w.Content.Update()
	w.button.Update()

	w.Dragging()
}

func (w *Window) Draw() {
	if w.hidden {
		return
	}
	contentArea := Area{
		X:      w.Area.X + borderSize,
		Y:      w.Area.Y + topBarSize + topBarOffset + borderSize,
		Width:  w.Area.Width - 2*borderSize,
		Height: w.Area.Height - 2*borderSize - topBarSize - topBarOffset,
	}
	w.background.Layout(w.Area)
	w.Content.Layout(contentArea)
	w.button.Layout(rl.Rectangle{Y: w.Area.Y + 23, X: w.Area.X + w.Area.Width - 72})

	ctx := Context{}
	w.background.Draw(&ctx)
	w.Content.Draw(&ctx)
	rl.DrawRectangleRec(w.barArea(), topBarColor)
	rl.DrawTextEx(assets.Manager.LoadedFont, w.Title, rl.Vector2{X: w.barArea().X + 8, Y: w.barArea().Y + 8}, 40, 0, rl.White)
	w.button.Draw(&ctx)
}

func (w *Window) HideShow() {
	w.hidden = !w.hidden
}

func (w *Window) Dragging() {
	if !rl.CheckCollisionPointRec(rl.GetMousePosition(), w.button.GetSize()) {

		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) && rl.CheckCollisionPointRec(rl.GetMousePosition(), w.barArea()) {
			w.dragging = true
			rl.SetMouseCursor(rl.MouseCursorResizeAll)
			return
		}

		if w.dragging && rl.IsMouseButtonReleased(rl.MouseButtonLeft) {
			w.dragging = false
			rl.SetMouseCursor(rl.MouseCursorArrow)
			return
		}
	}

	if w.dragging {
		delta := rl.GetMouseDelta()

		w.Area.X = w.Area.X + delta.X
		w.Area.Y = w.Area.Y + delta.Y
	}
}
