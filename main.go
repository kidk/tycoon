package main

import (
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kidk/tycoon/data"
	"github.com/kidk/tycoon/graphics"
	"github.com/kidk/tycoon/renderer"
)

func main() {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	logger.Print("Starting game")

	logger.Print("Setting window size and title")
	ebiten.SetWindowSize(1920, 1080)
	ebiten.SetWindowTitle("Hello, World!")
	g := &Game{}

	logger.Print("Creating grid")
	g.Grid = renderer.NewGrid(10)
	g.Grid.FillGrid()

	logger.Print("Creating image & sprite cache")
	imageCache := graphics.NewImageCache()
	g.SpriteCache = graphics.NewSpriteCache(imageCache)

	logger.Print("Creating sprites")
	data.LoadSprites(&g.SpriteCache)

	logger.Print("Setting initial screen")
	g.Screen = NewTextureDebugScreen(g.SpriteCache)

	logger.Print("Creating grid renderer")
	g.GridRenderer = renderer.NewGridRenderer(g.Grid)

	logger.Print("Everything is ready, starting loops")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
