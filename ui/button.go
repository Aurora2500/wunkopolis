package ui

import rl "github.com/gen2brain/raylib-go/raylib"

type Button struct {
	Col        rl.Color
	HoverCol   rl.Color
	PressedCol rl.Color
	OnClick    func()
}

func (b *Button) Draw(area rl.Rectangle, context *Context) {
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), area) {
		if rl.IsMouseButtonDown(0) {
			rl.DrawRectangleRec(area, b.PressedCol)
		} else {
			rl.DrawRectangleRec(area, b.HoverCol)
		}
	} else {
		rl.DrawRectangleRec(area, b.Col)
	}
}

func (b *Button) Check(area rl.Rectangle) {
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), area) {
		if rl.IsMouseButtonPressed(0) {
			b.OnClick()
		}

	}
}
