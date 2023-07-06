package graphics

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kidk/tycoon/engine"
)

type GridRenderer struct {
	spriteCache *SpriteCache
	grid        *engine.BlockGrid
	tx          int
	ty          int
}

func NewGridRenderer(spriteCache *SpriteCache, grid *engine.BlockGrid, tx int, ty int) GridRenderer {
	return GridRenderer{
		spriteCache: spriteCache,
		grid:        grid,
		tx:          tx,
		ty:          ty,
	}
}

func (gr *GridRenderer) Draw(screen *ebiten.Image) *ebiten.Image {
	image := ebiten.NewImage(gr.grid.Size*32, gr.grid.Size*32)
	for x := 0; x < gr.grid.Size; x++ {
		for y := 0; y < gr.grid.Size; y++ {
			block := gr.grid.Blocks[x][y]
			texture := block.GetTexture(gr.grid)

			if texture == "empty" {
				continue
			}

			sprite, err := gr.spriteCache.GetSprite(texture)
			if err != nil {
				sprite, _ = gr.spriteCache.GetSprite("error")
			}

			sprite.Draw(image, float64(gr.tx+(x*32)), float64(gr.ty+(y*32)))
		}
	}

	return image
}
