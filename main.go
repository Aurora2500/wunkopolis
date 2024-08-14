package main

import (
	"wunkopolis/assets"
	"wunkopolis/ui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(1600, 900, "Wunkopolis")
	defer rl.CloseWindow()
	defer assets.Manager.Unload()

	w1 := ui.Window{
		Area: ui.Area{
			X:      600,
			Y:      300,
			Width:  480,
			Height: 360,
		},
		Content: &ui.FancyPieChart{
			Segments: []ui.ChartSegment{
				{Col: rl.DarkGreen, N: 10},
				{Col: rl.DarkPurple, N: 6},
				{Col: rl.DarkBrown, N: 4},
				{Col: rl.DarkGray, N: 1},
			},
			Height:      40,
			Perspective: 76,
			Tint:        0.7,
		},
	}
	w1.Setup()
	for !rl.WindowShouldClose() {
		w1.Update()
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		w1.Draw()

		rl.DrawText("Wunkopolis", 285, 200, 40, rl.Black)
		rl.EndDrawing()
	}
}
