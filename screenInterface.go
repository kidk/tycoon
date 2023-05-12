package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Screen interface {
	Draw(game *Game, screen *ebiten.Image)
	Update(game *Game) error
}
