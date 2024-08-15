package main

import (
	"wunkopolis/assets"
	"wunkopolis/ui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var backgroundColor = rl.Color{R: 0, G: 130, B: 120, A: 255}

func main() {
	rl.InitWindow(1920, 1080, "Wunkopolis")
	defer rl.CloseWindow()
	defer assets.Manager.Unload()

	rl.SetTargetFPS(60)
	rl.ToggleBorderlessWindowed()

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
	}
	bottomBar.Setup()
	bottomBar.AddButton(ui.Button{Icon: assets.Manager.GetTexture("Statistics"), OnClick: func() { w1.HideShow() }})
	w1.Setup()
	for !rl.WindowShouldClose() {
		w1.Update()
		rl.BeginDrawing()

		rl.ClearBackground(backgroundColor)
		bottomBar.Draw()
		bottomBar.Update()
		w1.Draw()

		rl.EndDrawing()
	}
}
