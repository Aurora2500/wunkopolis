package main

import (
	"wunkopolis/appstate"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func RunGame() {
	rl.InitWindow(1920, 1080, "Wunkopolis")
	defer rl.CloseWindow()

	appstate := appstate.NewAppState()

	for !rl.WindowShouldClose() && !appstate.IsEmpty() {
		appstate.Update()
		// rl.BeginDrawing()

		// rl.ClearBackground(rl.RayWhite)
		// rl.DrawText("Wunkopolis", 500, 500, 40, rl.Black)

		// rl.EndDrawing()
	}
}
