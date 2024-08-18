package ui

import (
	"wunkopolis/assets"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Button struct {
	UIBase
	base       rl.Texture2D
	hover      rl.Texture2D
	pressed    rl.Texture2D
	Type       string
	Icon       rl.Texture2D
	Text       string
	FontSize   float32
	TextOffset float32
	OnClick    func()
}

func (b *Button) Layout(area Area) {

	if b.FontSize == 0 {
		b.FontSize = 24
	}
	if b.TextOffset == 0 {
		b.TextOffset = 10
	}
	b.base = assets.Manager.GetTexture(b.Type + "Button")
	b.hover = assets.Manager.GetTexture(b.Type + "ButtonHover")
	b.pressed = assets.Manager.GetTexture(b.Type + "ButtonPressed")
	b.RealSize = Area{Width: float32(b.base.Width), Height: float32(b.base.Height), X: area.X, Y: area.Y}
}

func (b *Button) Draw(context *Context) {
	texture := b.base
	textureRect := Area{X: 0, Y: 0, Width: float32(texture.Width), Height: float32(texture.Height)}
	iconRect := Area{X: 0, Y: 0, Width: float32(b.Icon.Width), Height: float32(b.Icon.Height)}
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), b.RealSize) {
		if rl.IsMouseButtonDown(0) {
			texture = b.pressed
		} else {
			texture = b.hover
		}
	}
	rl.DrawTexturePro(texture, textureRect, b.RealSize, rl.Vector2Zero(), 0, rl.White)
	rl.DrawTexturePro(b.Icon, iconRect, b.RealSize, rl.Vector2Zero(), 0, rl.White)
	if b.Text != "" {
		rl.DrawTextEx(assets.Manager.LoadedFont, b.Text, rl.Vector2{X: b.RealSize.X + b.TextOffset, Y: b.RealSize.Y + b.TextOffset}, b.FontSize, 0, rl.Black)
	}
}

func (b *Button) Update() {
	if b.OnClick == nil {
		return
	}
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), b.RealSize) {
		if rl.IsMouseButtonReleased(0) {
			b.OnClick()
		}
	}
}

func (b *Button) GetSize() Area {
	return b.RealSize
}
