package ui

import (
	"wunkopolis/assets"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Bar struct {
	NPatchBox
	Content Flexbox
}

func (b *Bar) Setup() {
	b.Texture = assets.Manager.GetTexture("Panel")
	b.NPatchInfo = rl.NPatchInfo{Source: rl.Rectangle{Width: 128, Height: 128}, Left: 8, Right: 8, Bottom: 8, Top: 8}
	b.RealSize = rl.Rectangle{Width: 1936, Height: 116, X: -8, Y: 976}
	buttonArea := Area{
		Width:  96,
		Height: 86,
	}
	testButton1 := Button{icon: assets.Manager.GetTexture("Statistics")}
	testButton2 := Button{icon: assets.Manager.GetTexture("Economy")}
	testButton1.Layout(buttonArea)
	testButton2.Layout(buttonArea)
	b.Content = Flexbox{Elements: []UIElem{
		&testButton1,
		&testButton2,
	}, Padding: 16}
	b.Content.Layout(Area{Width: b.RealSize.Width - 32, Height: b.RealSize.Height - 32, X: b.RealSize.X + 12, Y: b.RealSize.Y + 12})
}

func (b *Bar) Draw() {
	ctx := Context{}
	b.NPatchBox.Draw(&ctx)
	b.Content.Draw(&ctx)
}
