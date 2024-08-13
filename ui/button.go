package ui

import rl "github.com/gen2brain/raylib-go/raylib"

type Button struct {
	UIBase
	Base    rl.Texture2D
	Hover   rl.Texture2D
	Pressed rl.Texture2D
	OnClick func()
}

func (b *Button) Layout(area Area) {
	b.RealSize = area
}

func (b *Button) Draw(context *Context) {
	texture := b.Base
	textureRect := rl.Rectangle{X: 0, Y: 0, Width: float32(texture.Width), Height: float32(texture.Height)}
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), b.RealSize) {
		if rl.IsMouseButtonDown(0) {
			texture = b.Pressed
		} else {
			texture = b.Hover
		}
	}
	rl.DrawTexturePro(texture, textureRect, b.RealSize, rl.Vector2Zero(), 0, rl.White)
}

func (b *Button) Check() {
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), b.RealSize) {
		if rl.IsMouseButtonPressed(0) {
			b.OnClick()
		}
	}
}
