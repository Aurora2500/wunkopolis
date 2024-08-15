package ui

import rl "github.com/gen2brain/raylib-go/raylib"

type Button struct {
	UIBase
	base    rl.Texture2D
	hover   rl.Texture2D
	pressed rl.Texture2D
	icon    rl.Texture2D
	onClick func()
}

func (b *Button) Layout(area Area) {
	b.RealSize = area
}

func (b *Button) Draw(context *Context) {
	texture := b.base
	textureRect := rl.Rectangle{X: 0, Y: 0, Width: float32(texture.Width), Height: float32(texture.Height)}
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), b.RealSize) {
		if rl.IsMouseButtonDown(0) {
			texture = b.pressed
		} else {
			texture = b.hover
		}
	}
	rl.DrawTexturePro(texture, textureRect, b.RealSize, rl.Vector2Zero(), 0, rl.White)
	rl.DrawTexturePro(b.icon, textureRect, b.RealSize, rl.Vector2Zero(), 0, rl.White)
}

func (b *Button) Check() {
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), b.RealSize) {
		if rl.IsMouseButtonPressed(0) {
			b.onClick()
		}
	}
}
