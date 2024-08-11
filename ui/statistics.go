package ui

import rl "github.com/gen2brain/raylib-go/raylib"

type PieChartSegment struct {
	Col rl.Color
	N   float32
}

const totalsegments = 256

type PieChart struct {
	UIBase
	Segments []PieChartSegment
}

func (pc *PieChart) Layout(area Area) {
	pc.RealSize = area
}

func (pc *PieChart) Draw(ctx *Context) {
	radius := min(pc.RealSize.Height, pc.RealSize.Width)
	center := AreaCenter(pc.RealSize)

	var total float32

	for _, segment := range pc.Segments {
		total = total + segment.N
	}

	var angle float32

	for _, segment := range pc.Segments {
		portion := segment.N / total * 360.
		rl.DrawCircleSector(center, radius, angle, angle+portion, int32(totalsegments*portion/360), segment.Col)
		angle = angle + portion
	}
}
