package graphics

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type AnimatedSprite struct {
	sprite      *Sprite
	Frames      int // number of frames
	count       int // current frame
	speed       int // animation speed
	frameHeight int
	frameWidth  int
}

func NewAnimatedSprite(sprite *Sprite, frames int, speed int) AnimatedSprite {
	return AnimatedSprite{
		sprite:      sprite,
		Frames:      frames,
		count:       0,
		speed:       speed,
		frameHeight: 32,
		frameWidth:  32,
	}
}

func (asprite *AnimatedSprite) Update() {
	asprite.count += 1
}

func (asprite *AnimatedSprite) Draw(screen *ebiten.Image, sx float64, sy float64) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(sx, sy)
	i := (asprite.count / asprite.speed) % asprite.Frames
	ttx, tty := asprite.sprite.tx+i*asprite.frameWidth, asprite.sprite.ty
	tsx, tsy := asprite.sprite.sx+i*asprite.frameWidth, asprite.sprite.sy
	screen.DrawImage(asprite.sprite.image.SubImage(image.Rect(ttx, tty, tsx, tsy)).(*ebiten.Image), op)
}
