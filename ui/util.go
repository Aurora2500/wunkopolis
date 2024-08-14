package ui

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func AreaCenter(area Area) Vector2 {
	return Vector2{
		X: area.X + area.Width/2,
		Y: area.Y + area.Height/2,
	}
}

func ScaleAngle3D(angle, perspective float32) float32 {
	x, y := math.Sin(float64(angle)*rl.Deg2rad), math.Cos(float64(angle)*rl.Deg2rad)

	y = y * math.Cos(float64(perspective))

	return float32(math.Atan2(y, x))
}
