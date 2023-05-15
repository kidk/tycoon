package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kidk/tycoon/graphics"
)

type Game struct {
	SpriteCache graphics.SpriteCache
	Screen      Screen
}

func (g *Game) Update() error {
	if err := g.Screen.Update(g); err != nil {
		return err
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Screen.Draw(g, screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1920, 1080
}
