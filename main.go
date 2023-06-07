package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kidk/tycoon/data"
	"github.com/kidk/tycoon/graphics"
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	logger.Print("Starting game")

	logger.Print("Setting window size and title")
	ebiten.SetWindowSize(1680, 980)
	ebiten.SetWindowTitle("Hello, World!")
	g := &Game{}

	logger.Print("Creating image & sprite cache")
	imageCache := graphics.NewImageCache()
	g.SpriteCache = graphics.NewSpriteCache(imageCache)

	logger.Print("Creating sprites")
	data.LoadSprites(&g.SpriteCache)

	logger.Print("Setting initial screen")
	//g.Screen = NewTextureDebugScreen(&g.SpriteCache)
	//g.Screen = NewGameScreen(&g.SpriteCache)
	g.Screen = NewUITestScreen(&g.SpriteCache)

	logger.Print("Everything is ready, starting loops")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
