package ui

import rl "github.com/gen2brain/raylib-go/raylib"

type Area = rl.Rectangle
type Vector2 = rl.Vector2

type UIElem interface {
	Layout(area Area)
	Draw(context *Context)
	GetSize() Area
}

type UIBase struct {
	RealSize      Area
	PrefferedSize Area
}
