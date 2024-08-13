package ui

import rl "github.com/gen2brain/raylib-go/raylib"

const topBarSize = 32
const borderSize = 2

type Window struct {
	Content  UIElem
	Title    string
	Area     Area
	dragging bool
}

func (w *Window) barArea() Area {
	return Area{
		X:      w.Area.X,
		Y:      w.Area.Y,
		Width:  w.Area.Width,
		Height: topBarSize,
	}
}

func (w *Window) Update() {
	barArea := w.barArea()
	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) && rl.CheckCollisionPointRec(rl.GetMousePosition(), barArea) {
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
	rl.DrawRectangleRec(w.Area, rl.LightGray)
	rl.DrawRectangleRec(w.barArea(), rl.DarkBlue)
	contentArea := Area{
		X:      w.Area.X + borderSize,
		Y:      w.Area.Y + topBarSize,
		Width:  w.Area.Width - 2*borderSize,
		Height: w.Area.Height - borderSize - topBarSize,
	}
	w.Content.Layout(contentArea)
	ctx := Context{}
	w.Content.Draw(&ctx)
}
