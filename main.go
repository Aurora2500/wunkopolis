package main

import (
	"wunkopolis/assets"
	"wunkopolis/ui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var backgroundColor = rl.Color{R: 0, G: 130, B: 120, A: 255}
var Font rl.Font

func main() {
	rl.InitWindow(0, 0, "Wunkopolis")
	defer rl.CloseWindow()
	defer assets.Manager.Unload()

	rl.SetTargetFPS(60)
	rl.ToggleFullscreen()
	Font = assets.Manager.GetFont("W95FA")
	bottomBar := ui.Bar{}

	w1 := ui.Window{
		Area: ui.Area{
			X:      0,
			Y:      0,
			Width:  300,
			Height: 350,
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
		Icon:  assets.Manager.GetTexture("Statistics"),
		Title: "Statistics",
	}
	bottomBar.Setup()
	w1.Setup(&bottomBar)
	for !rl.WindowShouldClose() {
		w1.Update()
		rl.BeginDrawing()

		rl.ClearBackground(backgroundColor)
		bottomBar.Draw()
		bottomBar.Update()
		w1.Draw(Font)

		rl.EndDrawing()
	}
}
