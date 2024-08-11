package ui

import rl "github.com/gen2brain/raylib-go/raylib"

type Button struct {
	UIBase
	Col        rl.Color
	HoverCol   rl.Color
	PressedCol rl.Color
	OnClick    func()
}

func (b *Button) Layout(area Area) {
	b.RealSize = area
}

func (b *Button) Draw(context *Context) {
	col := b.Col
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), b.RealSize) {
		if rl.IsMouseButtonDown(0) {
			col = b.PressedCol
		} else {
			col = b.HoverCol
		}
	}
	rl.DrawRectangleRec(b.RealSize, col)
}

func (b *Button) Check(area rl.Rectangle) {
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), area) {
		if rl.IsMouseButtonPressed(0) {
			b.OnClick()
		}
	}
}
