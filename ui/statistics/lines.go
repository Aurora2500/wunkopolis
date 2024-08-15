package statistics

import (
	"wunkopolis/ui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type LineType int

const (
	SolidLine LineType = iota
	DashedLine
	DottedLine
	DashDottedLine
)

type Line struct {
	Type      LineType
	Thickness float32
	Color     rl.Color
	DashRatio float32
}

func (l *Line) Draw(from, to ui.Vector2) {
	switch l.Type {
	case SolidLine:
		rl.DrawLineEx(from, to, l.Thickness, l.Color)
	case DashedLine:
		lineLength := rl.Vector2Distance(to, from)
		numDashes := int32(lineLength / l.Thickness / (l.DashRatio + 1))
		for i := range numDashes {
			t1 := float32(i) / float32(numDashes)
			t2 := t1 + (l.DashRatio/(l.DashRatio+1))/float32(numDashes)
			a := rl.GetSplinePointLinear(from, to, t1)
			b := rl.GetSplinePointLinear(from, to, t2)
			rl.DrawLineEx(a, b, l.Thickness, l.Color)
		}
	case DottedLine:
		lineLength := rl.Vector2Distance(to, from)
		radius := l.Thickness / 2
		numDots := int32(lineLength / radius)
		for i := range numDots {
			center := rl.GetSplinePointLinear(to, from, float32(i)/float32(numDots))
			rl.DrawCircleV(center, radius, l.Color)
		}
	case DashDottedLine:

	}
}
