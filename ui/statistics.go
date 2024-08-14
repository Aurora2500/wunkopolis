package ui

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

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

type FancyPieChart struct {
	UIBase
	Segments    []ChartSegment
	Height      float32
	Perspective float32
	Tint        float32
}

func (pc *FancyPieChart) Layout(area Area) {
	pc.RealSize = area
}

func (pc *FancyPieChart) Draw(ctx *Context) {
	radius := min(pc.RealSize.Height, pc.RealSize.Width) / 2
	height := pc.Height * float32(math.Sin(float64(pc.Perspective*rl.Deg2rad)))
	center := AreaCenter(pc.RealSize)
	squish := float32(math.Cos(float64(pc.Perspective * rl.Deg2rad)))
	println(height)

	topCenter := center
	topCenter.Y = topCenter.Y - height/2
	bottomCenter := center
	bottomCenter.Y = bottomCenter.Y + height/2

	var total float32
	for _, seg := range pc.Segments {
		total = total + seg.N
	}

	// draw bottom part
	angle := pieChartStartAngle * rl.Deg2rad
	fan := make([]Vector2, 1, totalsegments)
	fan[0] = bottomCenter
	visible := false
	for _, segment := range pc.Segments {
		portion := float64(segment.N / total * 2 * math.Pi)
		if !visible {
			if angle+portion > 0 {
				visible = true
			} else {
				angle = angle + portion
				continue
			}
		}
		col := rl.ColorTint(segment.Col, rl.Color{R: uint8(pc.Tint * 255), G: uint8(pc.Tint * 255), B: uint8(pc.Tint * 255), A: 255})
		numPoints := int32(totalsegments * portion / (2 * math.Pi))
		deltaAngle := portion / float64(numPoints)
		// check left wall
		if angle+portion > math.Pi {
			fan = append(fan, Vector2{
				X: topCenter.X - radius,
				Y: topCenter.Y,
			}, Vector2{
				X: bottomCenter.X - radius,
				Y: bottomCenter.Y,
			})
		}
		for i := range numPoints + 1 {
			a := angle + float64(numPoints-i)*deltaAngle
			if a < 0 || a > math.Pi {
				continue
			}
			println(a)
			fan = append(fan, Vector2{
				X: bottomCenter.X + radius*float32(math.Cos(a)),
				Y: bottomCenter.Y + squish*radius*float32(math.Sin(a)),
			})
		}
		// check right wall
		if angle < 0 {
			fan = append(fan, Vector2{
				X: bottomCenter.X + radius,
				Y: bottomCenter.Y,
			}, Vector2{
				X: topCenter.X + radius,
				Y: topCenter.Y,
			})
		}
		rl.DrawTriangleFan(fan, col)
		angle = angle + portion
		fan = fan[:1]
		if angle > math.Pi {
			break
		}
	}

	// draw top part
	angle = pieChartStartAngle * rl.Deg2rad
	fan = fan[:1]
	fan[0] = topCenter
	for _, segment := range pc.Segments {
		fan = fan[:1]
		portion := float64(segment.N / total * 2 * math.Pi)
		numPoints := int32(totalsegments * portion / (2 * math.Pi))
		deltaAngle := portion / float64(numPoints)
		for i := range numPoints + 1 {
			a := angle + float64(numPoints-i)*deltaAngle
			fan = append(fan, Vector2{
				X: topCenter.X + radius*float32(math.Cos(a)),
				Y: topCenter.Y + squish*radius*float32(math.Sin(a)),
			})
		}
		rl.DrawTriangleFan(fan, segment.Col)
		angle = angle + portion
	}
}
