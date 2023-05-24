package main

import (
	"errors"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/kidk/tycoon/graphics"
	"github.com/kidk/tycoon/renderer"
	camera "github.com/melonfunction/ebiten-camera"
	"image/color"
)

type GameScreen struct {
	grid         renderer.Grid
	gridRenderer renderer.GridRenderer

	cam *camera.Camera
}

func NewGameScreen(spriteCache *graphics.SpriteCache) Screen {
	grid := renderer.NewGrid(30)
	grid.FillGrid()

	return &GameScreen{
		grid:         grid,
		gridRenderer: renderer.NewGridRenderer(spriteCache, &grid, 32, 32),

		cam: camera.NewCamera(1920, 1080, 500, 500, 0, 1),
	}
}

func (tds *GameScreen) Update(g *Game) error {
	// Keyboard
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		tds.cam.X -= 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		tds.cam.X += 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
		tds.cam.Y -= 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
		tds.cam.Y += 5
	}

	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return errors.New("Normal exit through escape key")
	}

	// Zoom
	_, scrollAmount := ebiten.Wheel()
	if scrollAmount > 0 {
		tds.cam.Zoom(1.1)
		if tds.cam.Scale > 1.5 {
			tds.cam.SetZoom(1.5)
		}
	} else if scrollAmount < 0 {
		tds.cam.Zoom(0.9)
		if tds.cam.Scale < 0.75 {
			tds.cam.SetZoom(0.75)
		}
	}

	return nil
}

func (tds *GameScreen) Draw(g *Game, screen *ebiten.Image) {
	// Clear camera surface
	tds.cam.Surface.Clear()
	tds.cam.Surface.Fill(color.RGBA{255, 128, 128, 255})

	// Draw tiles
	tileOps := &ebiten.DrawImageOptions{}
	tiles := tds.gridRenderer.Draw(screen)
	tds.cam.Surface.DrawImage(tiles, tds.cam.GetTranslation(tileOps, 0, 0))
	tiles.Dispose()

	// Draw to screen and zoom
	tds.cam.Blit(screen)

	ebitenutil.DebugPrint(screen,
		fmt.Sprintf(
			"Camera:\n  X: %3.3f\n  Y: %3.3f\n  W: %d\n  H: %d\n  Rot: %3.3f\n  Zoom: %3.3f\n"+
				"Other:\n  FPS: %3.3f\n  MouseX: %1.0f\n  MouseY: %1.0f",
			tds.cam.X, tds.cam.Y, tds.cam.Surface.Bounds().Size().X, tds.cam.Surface.Bounds().Size().Y, tds.cam.Rot, tds.cam.Scale, ebiten.ActualFPS(), 0, 0,
		))
}
