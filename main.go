package main

import (
	rand "math/rand"
	"wunkopolis/ui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(1600, 900, "Wunkopolis")
	defer rl.CloseWindow()

	uiCtx := ui.Context{}
	fb1 := ui.Flexbox{
		Elements: []ui.UIElem{},
	}

	fb1Area := rl.Rectangle{
		X:      500,
		Y:      600,
		Width:  1000,
		Height: 300,
	}
	b1 := ui.Button{Col: rl.Gray, HoverCol: rl.DarkGray, PressedCol: rl.Black, OnClick: func() {
		fb1.Elements = append(fb1.Elements, &ui.Box{Col: rl.ColorFromHSV(rand.Float32()*360, 1., 1.)})
		fb1.Layout(fb1.RealSize)
	}}
	b1Area := rl.Rectangle{X: 10, Y: 10, Width: 100, Height: 50}

	fb1.Layout(fb1Area)
	b1.Layout(b1Area)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("Wunkopolis", 500, 500, 40, rl.Black)

		fb1.Draw(&uiCtx)

		b1.Draw(&uiCtx)
		b1.Check()

		rl.EndDrawing()
	}
}
