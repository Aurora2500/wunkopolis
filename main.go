package main

import (
	"wunkopolis/ui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(1600, 900, "Wunkopolis")
	defer rl.CloseWindow()

	uiCtx := ui.Context{}
	np1Size := rl.Rectangle{
		X:      225,
		Y:      150,
		Width:  350,
		Height: 600,
	}
	np1 := ui.NPatchBox{Texture: rl.LoadTexture("sprites/Panel.png"), NPatchInfo: rl.NPatchInfo{Left: 8, Right: 8, Top: 8, Bottom: 8, Source: rl.Rectangle{Width: 128, Height: 128}}}

	np1.Layout(np1Size)

	w1 := ui.Window{
		Area: ui.Area{
			X:      600,
			Y:      300,
			Width:  480,
			Height: 360,
		},
		Content: &ui.PieChart{
			Segments: []ui.ChartSegment{
				{Col: rl.DarkGreen, N: 10},
				{Col: rl.DarkPurple, N: 6},
				{Col: rl.DarkBrown, N: 4},
				{Col: rl.DarkGray, N: 1},
			},
		},
	}
	for !rl.WindowShouldClose() {
		w1.Update()
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		np1.Draw(&uiCtx)

		w1.Draw()

		rl.DrawText("Wunkopolis", 285, 200, 40, rl.Black)
		rl.EndDrawing()
	}
}
