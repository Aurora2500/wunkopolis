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

	uiCtx := ui.Context{}
	np1Size := rl.Rectangle{
		X:      225,
		Y:      150,
		Width:  350,
		Height: 600,
	}
	np1 := ui.NPatchBox{Texture: rl.LoadTexture("sprites/Panel.png"), NPatchInfo: rl.NPatchInfo{Left: 8, Right: 8, Top: 8, Bottom: 8, Source: rl.Rectangle{Width: 128, Height: 128}}}

	np1.Layout(np1Size)
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		np1.Draw(&uiCtx)

		rl.DrawText("Wunkopolis", 285, 200, 40, rl.Black)
		rl.EndDrawing()
	}
}
