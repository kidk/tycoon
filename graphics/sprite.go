package graphics

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	image *ebiten.Image
	tx    int // translate x, so the location in the image
	ty    int // translate y, so the location in the image
	sx    int // size x
	sy    int // size y
}

func NewSprite(image *ebiten.Image, tx int, ty int, sx int, sy int) Sprite {
	return Sprite{
		image: image,
		tx:    tx,
		ty:    ty,
		sx:    sx,
		sy:    sy,
	}
}

func (sprite *Sprite) Draw(screen *ebiten.Image, sx float64, sy float64) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(sx, sy)

	screen.DrawImage(sprite.image.SubImage(image.Rect(sprite.tx, sprite.ty, sprite.sx, sprite.sy)).(*ebiten.Image), op)
}
