package ui

import rl "github.com/gen2brain/raylib-go/raylib"

type Box struct {
	UIBase
	Col rl.Color
}

func (b *Box) Layout(area rl.Rectangle) {
	b.RealSize = area
}

func (b *Box) Draw(ctx *Context) {
	rl.DrawRectangleRec(b.RealSize, b.Col)
}

func (b *Box) GetSize() Area {
	return b.RealSize
}

func (b *Box) Update() {

}
