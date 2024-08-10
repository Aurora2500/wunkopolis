package ui

import rl "github.com/gen2brain/raylib-go/raylib"

type UIElem interface {
	Draw(area rl.Rectangle, context *Context)
}
