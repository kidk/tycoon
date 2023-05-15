package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kidk/tycoon/graphics"
	"github.com/kidk/tycoon/renderer"
)

type GameScreen struct {
	grid         renderer.Grid
	gridRenderer renderer.GridRenderer
}

func NewGameScreen(spriteCache *graphics.SpriteCache) Screen {
	grid := renderer.NewGrid(30)
	grid.FillGrid()

	return &GameScreen{
		grid:         grid,
		gridRenderer: renderer.NewGridRenderer(spriteCache, &grid, 32, 32),
	}
}

func (tds *GameScreen) Update(g *Game) error {
	return nil
}

func (tds *GameScreen) Draw(g *Game, screen *ebiten.Image) {
	tds.gridRenderer.Draw(screen)
}
