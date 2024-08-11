package ui

import rl "github.com/gen2brain/raylib-go/raylib"

type ChartSegment struct {
	Col rl.Color
	N   float32
}

const totalsegments = 256
const pieChartStartAngle = -90.

type PieChart struct {
	UIBase
	Segments []ChartSegment
}

func (pc *PieChart) Layout(area Area) {
	pc.RealSize = area
}

func (pc *PieChart) Draw(ctx *Context) {
	radius := min(pc.RealSize.Height, pc.RealSize.Width) / 2
	center := AreaCenter(pc.RealSize)

	var total float32

	for _, segment := range pc.Segments {
		total = total + segment.N
	}

	var angle float32 = pieChartStartAngle

	for _, segment := range pc.Segments {
		portion := segment.N / total * 360.
		rl.DrawCircleSector(center, radius, angle, angle+portion, int32(totalsegments*portion/360), segment.Col)
		angle = angle + portion
	}
}

type TreemapChart struct {
	UIBase
	Segments []ChartSegment
}

func (tmc *TreemapChart) Layout(area Area) {
	tmc.RealSize = area
}

func (tmc *TreemapChart) Draw(ctx *Context) {
	var total float32

	for _, segment := range tmc.Segments {
		total = total + segment.N
	}

	var horizontal = true

	remainingArea := tmc.RealSize

	for _, segment := range tmc.Segments {
		portion := segment.N / total

		area := remainingArea

		if horizontal {
			area.Width = area.Width * portion
		} else {
			area.Height = area.Height * portion
		}

		rl.DrawRectangleRec(area, segment.Col)

		if horizontal {
			remainingArea.X = remainingArea.X + area.Width
			remainingArea.Width = remainingArea.Width - area.Width
		} else {
			remainingArea.Y = remainingArea.Y + area.Height
			remainingArea.Height = remainingArea.Height - area.Height
		}

		horizontal = !horizontal
		total = total - segment.N
	}
}
