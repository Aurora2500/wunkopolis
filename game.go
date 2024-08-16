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
	font := assets.Manager.GetFont("W95FA")
	bottomBar := ui.Bar{}

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
					Padding:  4,
					Elements: []ui.UIElem{&ui.Button{Type: "Long"}, &ui.Button{Type: "Long"}, &ui.Button{Type: "Long"}},
					Anchor:   ui.Center},
			}},
		Icon:  assets.Manager.GetTexture("Statistics"),
		Title: "Statistics",
	}
	bottomBar.Setup()
	w1.Setup(&bottomBar)
	for !rl.WindowShouldClose() {
		w1.Update()
		rl.BeginDrawing()

		rl.ClearBackground(backgroundColor)
		bottomBar.Draw()
		bottomBar.Update()
		w1.Draw(font)

		rl.EndDrawing()
	}
}
