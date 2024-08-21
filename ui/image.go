package ui

import rl "github.com/gen2brain/raylib-go/raylib"

type Image struct {
	UIBase
	Image rl.Texture2D
}

func (i *Image) Layout(area Area) {
	i.RealSize = Area{
		X:      area.X,
		Y:      area.Y,
		Width:  float32(i.Image.Width),
		Height: float32(i.Image.Height),
	}
}

func (i *Image) Update() {

}

func (i *Image) Draw(ctx *Context) {
	rl.DrawTexture(i.Image, i.RealSize.ToInt32().X, i.RealSize.ToInt32().Y, rl.White)
}

func (i *Image) GetSize() Area {
	return i.RealSize
}
