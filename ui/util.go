package ui

func AreaCenter(area Area) Vector2 {
	return Vector2{
		X: area.X + area.Width/2,
		Y: area.Y + area.Height/2,
	}
}

func InsetArea(area Area, inset float32) Area {
	return Area{
		X:      area.X + inset,
		Y:      area.Y + inset,
		Width:  area.Width - 2*inset,
		Height: area.Height - 2*inset,
	}
}
