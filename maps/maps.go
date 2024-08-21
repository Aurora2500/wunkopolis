package maps

import rl "github.com/gen2brain/raylib-go/raylib"

type Map struct {
	RoadsTexture rl.RenderTexture2D
}

func (m *Map) DrawMap(area rl.Rectangle) {

	rl.DrawTexture(m.RoadsTexture.Texture, area.ToInt32().X, area.ToInt32().Y, rl.White)

}

func (m *Map) Create(size rl.Rectangle) {
	m.RoadsTexture = rl.LoadRenderTexture(size.ToInt32().Width, size.ToInt32().Height)
	m.DrawRoads()
}

func (m *Map) DrawRoads() {

	rl.BeginTextureMode(m.RoadsTexture)
	rl.EndTextureMode()

}
