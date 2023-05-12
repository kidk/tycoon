package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type TextureDebugScreen struct {
}

func NewTextureDebugScreen() Screen {
	return &TextureDebugScreen{}
}

func (tds *TextureDebugScreen) Update(g *Game) error {
	return nil
}

func (tds *TextureDebugScreen) Draw(g *Game, screen *ebiten.Image) {
	//ebitenutil.DebugPrint(screen, "Hello, World!")
	g.GridRenderer.Draw(screen)
	sprite := g.SpriteCache.GetSprite("floor1")

	// Available sprites
	sprite.Draw(screen, 32, 32)
}
