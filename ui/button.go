package ui

import (
	"wunkopolis/assets"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Button struct {
	UIBase
	Base     rl.Texture2D
	Hover    rl.Texture2D
	Pressed  rl.Texture2D
	Toggle   rl.Texture2D
	Icon     rl.Texture2D
	Text     string
	FontSize float32
	OnClick  func()
	Toggled  bool
}

func (b *Button) Layout(area Area) {
	b.RealSize = Area{Width: float32(b.Base.Width), Height: float32(b.Base.Height), X: area.X, Y: area.Y}
}

func (b *Button) Draw(context *Context) {
	texture := b.Base
	textureRect := Area{X: 0, Y: 0, Width: float32(texture.Width), Height: float32(texture.Height)}
	iconRect := Area{X: 0, Y: 0, Width: float32(b.Icon.Width), Height: float32(b.Icon.Height)}
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), b.RealSize) {
		if rl.IsMouseButtonDown(0) {
			texture = b.Pressed
		} else {
			texture = b.Hover
		}
	}
	if b.Toggled {
		texture = b.Toggle
	}
	rl.DrawTexturePro(texture, textureRect, b.RealSize, rl.Vector2Zero(), 0, rl.White)
	rl.DrawTexturePro(b.Icon, iconRect, b.RealSize, rl.Vector2Zero(), 0, rl.White)
	if b.Text != "" {
		textSize := rl.MeasureTextEx(assets.Manager.LoadedFont, b.Text, b.FontSize, 0)
		rl.DrawTextEx(assets.Manager.LoadedFont, b.Text, rl.Vector2{X: b.RealSize.X + (b.RealSize.Width-textSize.X)/2, Y: b.RealSize.Y + (b.RealSize.Height-textSize.Y)/2}, b.FontSize, 0, rl.Black)
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
