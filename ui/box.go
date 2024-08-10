package ui

import rl "github.com/gen2brain/raylib-go/raylib"

type Box struct {
	Col rl.Color
}

func (b *Box) Draw(area rl.Rectangle, ctx *Context) {
	rl.DrawRectangleRec(area, b.Col)
}
