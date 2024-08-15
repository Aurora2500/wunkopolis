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
			X:      0,
			Y:      0,
			Width:  900,
			Height: 700,
		},
		Content: &ui.FancyPieChart{
			Segments: []ui.ChartSegment{
				{Col: rl.DarkGreen, N: 3},
				{Col: rl.DarkPurple, N: 2},
				{Col: rl.DarkBrown, N: 1},
				{Col: rl.DarkGray, N: 1},
				{Col: rl.Blue, N: 2},
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

		rl.EndDrawing()
	}
}
