package main

import (
	"wunkopolis/assets"
	"wunkopolis/lua"
	"wunkopolis/ui"
	"wunkopolis/variables"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var backgroundColor = rl.Color{R: 0, G: 130, B: 120, A: 255}

func run_game() {
	rl.InitWindow(int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()), "Wunkopolis")
	defer rl.CloseWindow()
	defer assets.Manager.Unload()
	rl.ToggleFullscreen()

	variables.Variables = make(map[string]int)
	variables.Variables["thriving"] = 12

	rl.SetTargetFPS(60)
	bottomBar := ui.Bar{}
	assets.Manager.LoadFont("W95FA")

	w1 := ui.Window{}
	if window, err := lua.GetWindow("StatisticsWindow"); err == nil {
		w1 = window
	}

	bottomBar.Setup()
	w1.Setup(&bottomBar)
	for !rl.WindowShouldClose() {
		w1.Update()
		bottomBar.Update()

		rl.BeginDrawing()

		rl.ClearBackground(backgroundColor)
		w1.Draw()
		bottomBar.Draw()

		rl.EndDrawing()
	}
}
