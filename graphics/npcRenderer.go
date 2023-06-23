package graphics

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kidk/tycoon/engine"
)

type NPCRenderer struct {
	spriteCache *SpriteCache
	sprite      *Sprite

	tx int
	ty int
}

func NewNPCRenderer(spriteCache *SpriteCache, npc *engine.Player, tx int, ty int) NPCRenderer {
	sprite, _ := spriteCache.GetSprite("character_adam_idle_down")
	return NPCRenderer{
		spriteCache: spriteCache,

		sprite: sprite,

		tx: tx,
		ty: ty,
	}
}

func (gr *NPCRenderer) Draw(screen *ebiten.Image) *ebiten.Image {
	image := ebiten.NewImage(32, 64)
	gr.sprite.Draw(image, float64(gr.tx), float64(gr.ty))
	return image
}
