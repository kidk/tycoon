package graphics

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kidk/tycoon/models"
)

type GridRenderer struct {
	grid models.Grid
}

func NewGridRenderer(grid models.Grid) GridRenderer {
	return GridRenderer{
		grid: grid,
	}
}

func (gr *GridRenderer) Draw(screen *ebiten.Image) {
	// image := ebiten.NewImage(30, 30)
	// image.Fill(color.RGBA{0xff, 0xff, 0xff, 0xff})
	// location := ebiten.GeoM{}
	// location.Translate(10, 10)
	// screen.DrawImage(image, &ebiten.DrawImageOptions{
	// 	GeoM: location,
	// })

}
