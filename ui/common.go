package ui

import rl "github.com/gen2brain/raylib-go/raylib"

type Area = rl.Rectangle

type UIElem interface {
	Layout(area Area)
	Draw(context *Context)
}

type UIBase struct {
	RealSize      Area
	PrefferedSize Area
}
