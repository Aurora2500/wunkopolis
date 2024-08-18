package main

import (
	"wunkopolis/assets"
	"wunkopolis/lua"
	"wunkopolis/ui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var backgroundColor = rl.Color{R: 0, G: 130, B: 120, A: 255}

func run_game() {
	rl.InitWindow(0, 0, "Wunkopolis")
	defer rl.CloseWindow()
	defer assets.Manager.Unload()

	rl.SetTargetFPS(60)
	rl.ToggleFullscreen()
	bottomBar := ui.Bar{}
	assets.Manager.LoadFont("W95FA")

	w1 := ui.Window{}
	if window, err := lua.GetWindow("StatisticsWindow"); err == nil {
		w1 = window
	}

	bottomBar.Setup()
	w1.Setup(&bottomBar)
	for !rl.WindowShouldClose() {

		rl.BeginDrawing()

		rl.ClearBackground(backgroundColor)
		w1.Draw()
		w1.Update()
		bottomBar.Draw()
		bottomBar.Update()

		rl.EndDrawing()
	}
}
