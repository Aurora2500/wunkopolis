package main

import (
	"wunkopolis/assets"
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

	w1 := ui.Window{
		Area: ui.Area{
			X:      0,
			Y:      0,
			Width:  632,
			Height: 350,
		},
		Content: &ui.Flexbox{
			Dir: ui.Col,
			Elements: []ui.UIElem{
				&ui.Flexbox{
					Dir: ui.Col,
					Elements: []ui.UIElem{
						&ui.Flexbox{
							Padding: 4,
							Elements: []ui.UIElem{
								&ui.Button{Type: "Long", Text: "Test one"},
								&ui.Button{Type: "Long", Text: "Test two"},
								&ui.Button{Type: "Long", Text: "Test three"},
							},
							Anchor: ui.Center},
					},
				},
			},
		},
		Icon:  assets.Manager.GetTexture("Statistics"),
		Title: "Statistics",
	}
	bottomBar.Setup()
	w1.Setup(&bottomBar)
	for !rl.WindowShouldClose() {

		rl.BeginDrawing()

		rl.ClearBackground(backgroundColor)
		bottomBar.Draw()
		bottomBar.Update()
		w1.Draw()
		w1.Update()

		rl.EndDrawing()
	}
}
