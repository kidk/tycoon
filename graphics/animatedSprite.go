package graphics

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

type AnimatedSprite struct {
	Sprite      *Sprite
	Frames      int // number of frames
	count       int // current frame
	Speed       int // animation speed
	FrameHeight int
	FrameWidth  int
}

func NewAnimatedSprite(sprite *Sprite, frames int, speed int) AnimatedSprite {
	return AnimatedSprite{
		Sprite:      sprite,
		Frames:      frames,
		count:       0,
		Speed:       speed,
		FrameHeight: 32,
		FrameWidth:  32,
	}
}

func (asprite *AnimatedSprite) Update() {
	asprite.count += 1
}

func (asprite *AnimatedSprite) Draw(screen *ebiten.Image, sx float64, sy float64) {
	op := &ebiten.DrawImageOptions{}
	//op.GeoM.Translate(-float64(asprite.FrameWidth)/2, -float64(asprite.FrameHeight)/2)
	op.GeoM.Translate(sx, sy)
	i := (asprite.count / asprite.Speed) % asprite.Frames
	ttx, tty := asprite.Sprite.tx+i*asprite.FrameWidth, asprite.Sprite.ty
	tsx, tsy := asprite.Sprite.sx+i*asprite.FrameWidth, asprite.Sprite.sy
	screen.DrawImage(asprite.Sprite.image.SubImage(image.Rect(ttx, tty, tsx, tsy)).(*ebiten.Image), op)
}
