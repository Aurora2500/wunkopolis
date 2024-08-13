package assets

import (
	"io/fs"
	"log"
	"os"
)

var Manager AssetManager

type AssetManager struct {
	assetDir fs.FS
}

func init() {
	basePath, err := os.Executable()
	if err != nil {
		panic("Couldn't get base path for the game executable")
	}
	baseFS := os.DirFS(basePath)

	assetDir, err := fs.Sub(baseFS, "assets")
	if err != nil {
		log.Fatal("No assets directory found")
	} else {
		Manager.assetDir = assetDir
	}

}
