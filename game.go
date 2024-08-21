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

	windows := []ui.Window{{
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
	},
		{
			Content: &ui.Scrollbox{},
			Title:   "Map",
			Area:    rl.Rectangle{Width: 1000, Height: 800},
			Icon:    assets.Manager.GetTexture("Map"),
		},
	}

	bottomBar.Setup()
	for i := range windows {
		windows[i].Setup(&bottomBar)
	}
	for !rl.WindowShouldClose() {
		bottomBar.Update()
		for i := range windows {
			windows[i].Update()
		}
		rl.BeginDrawing()

		rl.ClearBackground(backgroundColor)
		for _, w := range windows {
			w.Draw()
		}
		bottomBar.Draw()

		rl.EndDrawing()
	}
}
