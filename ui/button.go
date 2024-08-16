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
	Scale   float32
	Type    string
	Icon    rl.Texture2D
	OnClick func()
}

func (b *Button) Layout(area Area) {
	if b.Scale == 0 {
		b.Scale = 1
	}
	b.base = assets.Manager.GetTexture(b.Type + "Button")
	b.hover = assets.Manager.GetTexture(b.Type + "ButtonHover")
	b.pressed = assets.Manager.GetTexture(b.Type + "ButtonPressed")
	b.RealSize = Area{Width: float32(b.base.Width) * b.Scale, Height: float32(b.base.Height) * b.Scale, X: area.X, Y: area.Y}
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
