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

	t := ui.Text{
		Font:    assets.Manager.LoadedFont,
		Content: lorem_ipsum,
	}

	t.Layout(ui.Area{
		X: 1000, Y: 600,
		Width:  400,
		Height: 800,
	})

	bottomBar.Setup()
	w1.Setup(&bottomBar)
	for !rl.WindowShouldClose() {

		rl.BeginDrawing()

		rl.ClearBackground(backgroundColor)
		w1.Draw()
		w1.Update()
		bottomBar.Draw()
		bottomBar.Update()

		t.Draw(&ui.Context{})

		rl.EndDrawing()
	}
}
