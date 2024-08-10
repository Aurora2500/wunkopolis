package main

import (
	"wunkopolis/ui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(1920, 1080, "Wunkopolis")
	defer rl.CloseWindow()

	uiCtx := ui.Context{}
	fb1 := ui.Flexbox{
		Dir: ui.Row,
		Elements: []ui.UIElem{
			&ui.Box{rl.Red},
			&ui.Box{rl.Green},
			&ui.Box{rl.Blue},
			&ui.Box{rl.Pink},
			&ui.Box{rl.Brown},
			&ui.Box{rl.Yellow},
		},
	}
	fb1Area := rl.Rectangle{
		X:      500,
		Y:      600,
		Width:  1000,
		Height: 300,
	}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("Wunkopolis", 500, 500, 40, rl.Black)

		fb1.Draw(fb1Area, &uiCtx)

		rl.EndDrawing()
	}
}
