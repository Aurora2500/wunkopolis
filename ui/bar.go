package ui

import (
	"wunkopolis/assets"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Bar struct {
	NPatchBox
	Content Flexbox
}

const barHeight = 116
const sideSize = 8
const padding = 16

func (b *Bar) Setup() {
	b.Texture = assets.Manager.GetTexture("Panel")
	b.NPatchInfo = rl.NPatchInfo{Source: rl.Rectangle{Width: float32(b.Texture.Width) - sideSize*2, Height: float32(b.Texture.Height), X: sideSize}, Left: sideSize, Right: sideSize, Bottom: sideSize, Top: sideSize}
	b.RealSize = rl.Rectangle{Width: float32(rl.GetScreenWidth()), Height: barHeight, X: 0, Y: float32(rl.GetScreenHeight()) - barHeight + sideSize}

	b.Content = Flexbox{Elements: []UIElem{}, Padding: padding}
	b.Content.Layout(Area{Width: b.RealSize.Width, Height: b.RealSize.Height - sideSize*4, X: b.RealSize.X + sideSize, Y: b.RealSize.Y + sideSize*2})
}

func (b *Bar) Draw() {
	ctx := Context{}
	b.NPatchBox.Draw(&ctx)
	b.Content.Draw(&ctx)
}

func (b *Bar) Update() {
	b.Content.Update()
}

func (b *Bar) AddButton(button Button) {
	buttonArea := Area{
		Width:  96,
		Height: 86,
	}
	button.Layout(buttonArea)
	b.Content.Elements = append(b.Content.Elements, &button)
	b.Content.Layout(Area{Width: b.RealSize.Width - 32, Height: b.RealSize.Height - 32, X: b.RealSize.X + 12, Y: b.RealSize.Y + 12})
}
