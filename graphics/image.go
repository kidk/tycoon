package graphics

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/sirupsen/logrus"
)

type Image struct {
	texture *ebiten.Image
}

func NewImage(filePath string) Image {
	image := Image{}
	texture, _, err := ebitenutil.NewImageFromFile(filePath)
	if err != nil {
		logrus.Fatal(err)
	}
	image.texture = texture

	return image
}
