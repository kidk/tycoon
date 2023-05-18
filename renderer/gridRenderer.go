package renderer

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kidk/tycoon/graphics"
)

type GridRenderer struct {
	spriteCache *graphics.SpriteCache
	grid        *Grid
	tx          int
	ty          int
}

func NewGridRenderer(spriteCache *graphics.SpriteCache, grid *Grid, tx int, ty int) GridRenderer {
	return GridRenderer{
		spriteCache: spriteCache,
		grid:        grid,
		tx:          tx,
		ty:          ty,
	}
}

func (gr *GridRenderer) Draw(screen *ebiten.Image) {
	for x := 0; x < gr.grid.size; x++ {
		for y := 0; y < gr.grid.size; y++ {
			block := gr.grid.blocks[x][y]

			sprite, err := gr.spriteCache.GetSprite(block.Visual.Name)
			if err != nil {
				sprite, _ = gr.spriteCache.GetSprite("error")
			}

			sprite.Draw(screen, float64(gr.tx+(x*32)), float64(gr.ty+(y*32)))
		}
	}
}
