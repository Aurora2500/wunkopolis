package ui

import rl "github.com/gen2brain/raylib-go/raylib"

type NPatchBox struct {
	UIBase
	Texture    rl.Texture2D
	NPatchInfo rl.NPatchInfo
}

func (b *NPatchBox) Layout(area rl.Rectangle) {
	b.RealSize = area
}

func (b *NPatchBox) Draw(ctx *Context) {
	rl.DrawTextureNPatch(b.Texture, b.NPatchInfo, b.RealSize, rl.Vector2Zero(), 0, rl.White)
}

func (b *NPatchBox) GetSize() Area {
	return b.RealSize
}

func (b *NPatchBox) Update() {

}
