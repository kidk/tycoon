package graphics

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kidk/tycoon/engine"
)

type NPCRenderer struct {
	npc *engine.Player

	spriteIdle *Sprite

	spriteUp    AnimatedSprite
	spriteDown  AnimatedSprite
	spriteLeft  AnimatedSprite
	spriteRight AnimatedSprite

	tx int
	ty int
}

func NewNPCRenderer(spriteCache *SpriteCache, npc *engine.Player, tx int, ty int) NPCRenderer {
	spriteIdle, _ := spriteCache.GetSprite("character_adam_idle_down")
	spriteUp, _ := spriteCache.GetSprite("character_adam_move_up")
	spriteUpAnimation := NewAnimatedSprite(spriteUp, 6, 10)
	spriteDown, _ := spriteCache.GetSprite("character_adam_move_down")
	spriteDownAnimation := NewAnimatedSprite(spriteDown, 6, 10)
	spriteLeft, _ := spriteCache.GetSprite("character_adam_move_left")
	spriteLeftAnimation := NewAnimatedSprite(spriteLeft, 6, 10)
	spriteRight, _ := spriteCache.GetSprite("character_adam_move_right")
	spriteRightAnimation := NewAnimatedSprite(spriteRight, 6, 10)

	return NPCRenderer{
		npc: npc,

		spriteIdle:  spriteIdle,
		spriteUp:    spriteUpAnimation,
		spriteDown:  spriteDownAnimation,
		spriteLeft:  spriteLeftAnimation,
		spriteRight: spriteRightAnimation,

		tx: tx,
		ty: ty,
	}
}

func (gr *NPCRenderer) Update() {
	gr.spriteDown.Update()
	gr.spriteUp.Update()
	gr.spriteLeft.Update()
	gr.spriteRight.Update()
}

func (gr *NPCRenderer) Draw(screen *ebiten.Image) *ebiten.Image {
	// TODO: Hier switch case om verschillende animations op te vangen
	// gr.spriteUp.Draw(image, float64(gr.tx), float64(gr.ty))
	image := ebiten.NewImage(32, 64)
	//gr.spriteIdle.Draw(image, float64(gr.tx), float64(gr.ty))
	gr.spriteDown.Draw(image, float64(gr.tx), float64(gr.ty))
	return image
}
