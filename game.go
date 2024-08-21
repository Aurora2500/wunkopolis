package main

import (
	"wunkopolis/assets"
	"wunkopolis/ui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var backgroundColor = rl.Color{R: 0, G: 130, B: 120, A: 255}
var Variables map[string]int

func run_game() {
	rl.InitWindow(int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()), "Wunkopolis")
	defer rl.CloseWindow()
	defer assets.Manager.Unload()
	rl.ToggleFullscreen()

	Variables = make(map[string]int)
	rl.SetTargetFPS(60)
	bottomBar := ui.Bar{}
	assets.Manager.LoadFont("W95FA")

	w1 := ui.Window{
		Content: &ui.Flexbox{
			Direction: ui.Col,
			Border:    10,
			Background: ui.NPatchBox{
				Texture: assets.Manager.GetTexture("InteriorPanel"),
				NPatchInfo: rl.NPatchInfo{
					Source: rl.Rectangle{Width: float32(assets.Manager.GetTexture("InteriorPanel").Width), Height: float32(assets.Manager.GetTexture("InteriorPanel").Height)},
					Left:   10,
					Right:  10,
					Top:    10,
					Bottom: 10,
				},
			},
		},
		Title: "Statistics",
		Area:  rl.Rectangle{Width: 620, Height: 900},
		Icon:  assets.Manager.GetTexture("Statistics"),
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
