package ui

import rl "github.com/gen2brain/raylib-go/raylib"

type Context struct {
	// stack of clipped areas of the screen an element can draw onto
	clipping []rl.Rectangle
}
