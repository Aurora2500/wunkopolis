package statistics

import (
	"math"

	ui "wunkopolis/ui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type ChartSegment struct {
	Col rl.Color
	N   float32
}

const totalsegments = 256
const pieChartStartAngle = -90.

type PieChart struct {
	ui.UIBase
	Segments []ChartSegment
}

func (pc *PieChart) Layout(area ui.Area) {
	pc.RealSize = area
}

func (pc *PieChart) GetSize() ui.Area {
	return pc.RealSize
}

func (pc *PieChart) Update() {

}

func (pc *PieChart) Draw(ctx *ui.Context) {
	radius := min(pc.RealSize.Height, pc.RealSize.Width) / 2
	center := ui.AreaCenter(pc.RealSize)

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
	ui.UIBase
	Segments []ChartSegment
}

func (tmc *TreemapChart) Layout(area ui.Area) {
	tmc.RealSize = area
}

func (tmc *TreemapChart) GetSize() ui.Area {
	return tmc.RealSize
}

func (tmc *TreemapChart) Update() {

}

func (tmc *TreemapChart) Draw(ctx *ui.Context) {
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
	ui.UIBase
	Segments    []ChartSegment
	Height      float32
	Perspective float32
	Tint        float32
}

func (pc *FancyPieChart) Layout(area ui.Area) {
	pc.RealSize = area
}

func (pc *FancyPieChart) GetSize() ui.Area {
	return pc.RealSize
}

func (pc *FancyPieChart) Update() {

}

func (pc *FancyPieChart) Draw(ctx *ui.Context) {
	radius := min(pc.RealSize.Height, pc.RealSize.Width) / 2
	height := pc.Height * float32(math.Sin(float64(pc.Perspective*rl.Deg2rad)))
	center := ui.AreaCenter(pc.RealSize)
	squish := float32(math.Cos(float64(pc.Perspective * rl.Deg2rad)))

	topCenter := center
	topCenter.Y = topCenter.Y - height/2
	bottomCenter := center
	bottomCenter.Y = bottomCenter.Y - height/4
	var total float32
	for _, seg := range pc.Segments {
		total = total + seg.N
	}

	// draw bottom part
	angle := pieChartStartAngle * rl.Deg2rad
	fan := make([]ui.Vector2, 1, totalsegments)
	visible := false
	fan[0] = topCenter
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
			fan = append(fan, ui.Vector2{
				X: topCenter.X - radius,
				Y: topCenter.Y,
			}, ui.Vector2{
				X: bottomCenter.X - radius,
				Y: bottomCenter.Y,
			})
		}
		for i := range numPoints + 1 {
			a := angle + float64(numPoints-i)*deltaAngle
			if a < 0 || a > math.Pi {
				continue
			}
			fan = append(fan, ui.Vector2{
				X: bottomCenter.X + radius*float32(math.Cos(a)),
				Y: bottomCenter.Y + squish*radius*float32(math.Sin(a)),
			})
		}
		// check right wall
		if angle < 0 {
			fan = append(fan, ui.Vector2{
				X: bottomCenter.X + radius,
				Y: bottomCenter.Y,
			}, ui.Vector2{
				X: topCenter.X + radius,
				Y: topCenter.Y,
			})
		}
		fan[0].X = fan[1].X
		fan[len(fan)-1].Y = topCenter.Y
		fan[len(fan)-2].X = fan[len(fan)-1].X
		if fan[0].Y > center.Y && fan[0].X < center.X {
			fan[0].X = fan[len(fan)-1].X
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
			fan = append(fan, ui.Vector2{
				X: topCenter.X + radius*float32(math.Cos(a)),
				Y: topCenter.Y + squish*radius*float32(math.Sin(a)),
			})
		}
		rl.DrawTriangleFan(fan, segment.Col)
		angle = angle + portion
	}
}
