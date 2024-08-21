package ui

import (
	"wunkopolis/assets"
)

type Tabs struct {
	UIBase
	Index      int
	Tabs       []UIElem
	TabNames   []string
	TabButtons Flexbox
	FontSize   float32
}

func (t *Tabs) Layout(area Area) {
	t.RealSize = area
	t.TabButtons.Layout(Area{
		Width:  area.Width,
		Height: 43,
		X:      area.X,
		Y:      area.Y,
	})
	for _, tab := range t.Tabs {
		tab.Layout(Area{
			Width:  area.Width,
			Height: area.Height - 50,
			X:      area.X,
			Y:      area.Y + 50,
		})
	}
}

func (t *Tabs) GetSize() Area {
	return t.RealSize
}

func (t *Tabs) Draw(ctx *Context) {
	t.TabButtons.Draw(ctx)
	t.Tabs[t.Index].Draw(ctx)
}

func (t *Tabs) Update() {
	t.TabButtons.Update()
	for _, elem := range t.Tabs {
		elem.Update()
	}
}

func (t *Tabs) Setup() {
	for i := range t.Tabs {
		toggled := false
		if i == t.Index {
			toggled = true
		}
		t.TabButtons.Add(&Button{
			Toggled:  toggled,
			Base:     assets.Manager.GetTexture("LongButton"),
			Hover:    assets.Manager.GetTexture("LongButtonHover"),
			Pressed:  assets.Manager.GetTexture("LongButtonPressed"),
			Toggle:   assets.Manager.GetTexture("LongButtonToggled"),
			Text:     t.TabNames[i],
			FontSize: t.FontSize,
			OnClick: func() {
				t.Index = i
				for bi, elem := range t.TabButtons.Elements {
					if button, ok := elem.(*Button); ok {
						button.Toggled = false
						if bi == i {
							button.Toggled = true
						}
					}
				}
			},
		})
	}
}
