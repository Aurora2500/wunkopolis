package ui

import rl "github.com/gen2brain/raylib-go/raylib"

type Context struct {
	// stack of clipped areas of the screen an element can draw onto
	clipping []Area
}

func (c *Context) PushScissor(area Area) {
	if len(c.clipping) > 0 {
		area = rl.GetCollisionRec(area, c.clipping[len(c.clipping)-1])
	}

	c.clipping = append(c.clipping, area)
	areaInt := area.ToInt32()
	rl.BeginScissorMode(areaInt.X, areaInt.Y, areaInt.Width, areaInt.Height)
}

func (c *Context) PopScissor() {
	l := len(c.clipping)
	if l == 0 {
		return
	}

	l--
	c.clipping = c.clipping[:l]

	if l == 0 {
		rl.EndScissorMode()
	} else {
		newArea := c.clipping[l-1].ToInt32()
		rl.BeginScissorMode(newArea.X, newArea.Y, newArea.Width, newArea.Height)
	}
}
