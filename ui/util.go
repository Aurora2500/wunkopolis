package ui

func AreaCenter(area Area) Vector2 {
	return Vector2{
		X: area.X + area.Width/2,
		Y: area.Y + area.Height/2,
	}
}
