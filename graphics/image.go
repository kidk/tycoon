package graphics

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Image struct {
	texture *ebiten.Image
}

func NewImage(filePath string) Image {
	image := Image{}
	texture, _, err := ebitenutil.NewImageFromFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	image.texture = texture

	return image
}
