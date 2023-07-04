package main

import (
	"github.com/sirupsen/logrus"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kidk/tycoon/data"
	"github.com/kidk/tycoon/graphics"
)

func main() {
	logrus.Print("Starting game")
	logrus.SetLevel(logrus.TraceLevel)

	logrus.Print("Setting window size and title")
	ebiten.SetWindowSize(1680, 980)
	ebiten.SetWindowTitle("Hello, World!")
	g := &Game{}

	logrus.Print("Creating image & sprite cache")
	imageCache := graphics.NewImageCache()
	g.SpriteCache = graphics.NewSpriteCache(imageCache)

	logrus.Print("Creating sprites")
	data.LoadSprites(&g.SpriteCache)

	logrus.Print("Setting initial screen")
	//g.Screen = NewTextureDebugScreen(&g.SpriteCache)
	g.Screen = NewGameScreen(&g.SpriteCache)
	//g.Screen = NewUITestScreen(&g.SpriteCache)

	logrus.Print("Everything is ready, starting loops")
	if err := ebiten.RunGame(g); err != nil {
		logrus.Fatal(err)
	}
}
