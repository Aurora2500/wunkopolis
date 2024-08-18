package assets

import (
	"os"
	"path"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var Manager AssetManager

type AssetManager struct {
	assetDir       string
	loadedTextures map[string]rl.Texture2D
	LoadedFont     rl.Font
}

func init() {
	basePath, err := os.Getwd()
	if err != nil {
		panic("Couldn't get base path for the game executable")
	}

	assetDir := path.Join(basePath, "data")

	Manager.assetDir = assetDir
	Manager.loadedTextures = make(map[string]rl.Texture2D)
}

func (am *AssetManager) LoadFont(name string) {
	Manager.LoadedFont = rl.LoadFont(path.Join(am.assetDir, name+".otf"))
}

func (am *AssetManager) GetTexture(name string) rl.Texture2D {
	tex, ok := am.loadedTextures[name]
	if ok {
		return tex
	}

	imagePath := path.Join(am.assetDir, name+".png")
	tex = rl.LoadTexture(imagePath)
	am.loadedTextures[name] = tex
	return tex
}

func (am *AssetManager) Unload() {
	for _, tex := range am.loadedTextures {
		rl.UnloadTexture(tex)
	}
	rl.UnloadFont(am.LoadedFont)

}
