package ui

import (
	"wunkopolis/assets"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Button struct {
	UIBase
	base    rl.Texture2D
	hover   rl.Texture2D
	pressed rl.Texture2D
	Icon    rl.Texture2D
	OnClick func()
}

func (b *Button) Layout(area Area) {
	b.RealSize = area
	b.base = assets.Manager.GetTexture("Button")
	b.hover = assets.Manager.GetTexture("ButtonHover")
	b.pressed = assets.Manager.GetTexture("ButtonPressed")
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
	rl.DrawTexturePro(b.Icon, textureRect, b.RealSize, rl.Vector2Zero(), 0, rl.White)
}

func (b *Button) Update() {
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), b.RealSize) {
		if rl.IsMouseButtonReleased(0) {
			b.OnClick()
		}
	}
}

func (b *Button) GetSize() Area {
	return b.RealSize
}
