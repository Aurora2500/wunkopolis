package main

import (
	"wunkopolis/assets"
	"wunkopolis/lua"
	"wunkopolis/ui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var backgroundColor = rl.Color{R: 0, G: 130, B: 120, A: 255}

const lorem_ipsum = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."

func run_game() {
	rl.InitWindow(int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()), "Wunkopolis")
	defer rl.CloseWindow()
	defer assets.Manager.Unload()
	rl.ToggleFullscreen()

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
