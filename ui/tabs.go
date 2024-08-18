package ui

import "wunkopolis/assets"

type Tabs struct {
	UIBase
	Index      int
	Tabs       []UIElem
	TabButtons Flexbox
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
			Height: area.Height - 43,
			X:      area.X,
			Y:      area.Y + 43,
		})
	}
}

func (t *Tabs) GetSize() Area {
	return t.RealSize
}

func (t *Tabs) Draw(ctx *Context) {

	t.Tabs[t.Index].Draw(ctx)
	t.TabButtons.Draw(ctx)
}

func (t *Tabs) Update() {
	t.TabButtons.Update()
}

func (t *Tabs) Setup() {
	for i := range t.Tabs {
		t.TabButtons.Add(&Button{
			Base:    assets.Manager.GetTexture("LongButton"),
			Hover:   assets.Manager.GetTexture("LongButtonHover"),
			Pressed: assets.Manager.GetTexture("LongButtonPressed"),
			OnClick: func() { t.Index = i },
		})
	}
}
