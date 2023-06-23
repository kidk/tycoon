package main

import (
	"errors"
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/kidk/tycoon/engine"
	"github.com/kidk/tycoon/graphics"
	"github.com/kidk/tycoon/helpers"
	camera "github.com/melonfunction/ebiten-camera"
)

type GameScreen struct {
	state *engine.Engine

	floorGridRenderer    graphics.GridRenderer
	buildingGridRenderer graphics.GridRenderer

	cam      *camera.Camera
	keyboard *helpers.KeyboardHelper

	buildMode bool

	playerRenderer graphics.NPCRenderer
}

func NewGameScreen(spriteCache *graphics.SpriteCache) Screen {
	state := engine.NewEngine()

	return &GameScreen{
		state:                state,
		floorGridRenderer:    graphics.NewGridRenderer(spriteCache, &state.FloorGrid, 0, 0),
		buildingGridRenderer: graphics.NewGridRenderer(spriteCache, &state.BuildingGrid, 0, 0),

		cam:      camera.NewCamera(1920, 1080, 500, 500, 0, 1),
		keyboard: helpers.NewKeyboardHelper(),

		buildMode: false,

		playerRenderer: graphics.NewNPCRenderer(spriteCache, state.Player, 0, 0),
	}
}

func (tds *GameScreen) Update(g *Game) error {
	// Keyboard
	tds.keyboard.Update()

	if tds.keyboard.IsKeyPressed(ebiten.KeyArrowLeft) || tds.keyboard.IsKeyPressed(ebiten.KeyA) {
		tds.cam.X -= 5
	}
	if tds.keyboard.IsKeyPressed(ebiten.KeyArrowRight) || tds.keyboard.IsKeyPressed(ebiten.KeyD) {
		tds.cam.X += 5
	}
	if tds.keyboard.IsKeyPressed(ebiten.KeyArrowUp) || tds.keyboard.IsKeyPressed(ebiten.KeyW) {
		tds.cam.Y -= 5
	}
	if tds.keyboard.IsKeyPressed(ebiten.KeyArrowDown) || tds.keyboard.IsKeyPressed(ebiten.KeyS) {
		tds.cam.Y += 5
	}

	if tds.keyboard.IsKeyPressedOnce(ebiten.KeyQ) {
		tds.buildMode = !tds.buildMode
	}

	if tds.keyboard.IsKeyPressedOnce(ebiten.KeyEscape) {
		return errors.New("normal exit through escape key")
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

	// Draw ground
	groundOps := &ebiten.DrawImageOptions{}
	ground := tds.floorGridRenderer.Draw(screen)
	tds.cam.Surface.DrawImage(ground, tds.cam.GetTranslation(groundOps, 0, 0))
	ground.Dispose()

	// Draw buildings
	buildingOps := &ebiten.DrawImageOptions{}
	building := tds.buildingGridRenderer.Draw(screen)
	tds.cam.Surface.DrawImage(building, tds.cam.GetTranslation(buildingOps, 0, 0))
	building.Dispose()

	// Draw unit
	playerOps := &ebiten.DrawImageOptions{}
	player := tds.playerRenderer.Draw(screen)
	tds.cam.Surface.DrawImage(player, tds.cam.GetTranslation(playerOps, float64(0+(tds.state.Player.X*32)), float64(-32+(tds.state.Player.Y*32))))
	player.Dispose()

	// Draw to screen and zoom
	tds.cam.Blit(screen)

	ebitenutil.DebugPrint(screen,
		fmt.Sprintf(`
		State:
			Buildmode: %t
		Camera:
			X: %3.3f
			Y: %3.3f
			W: %d
			H: %d
			Rot: %3.3f
			Zoom: %3.3f
		Other:
			FPS: %3.3f
			MouseX: %1.0f
			MouseY: %1.0f
		`,
			tds.buildMode,
			tds.cam.X,
			tds.cam.Y,
			tds.cam.Surface.Bounds().Size().X,
			tds.cam.Surface.Bounds().Size().Y,
			tds.cam.Rot,
			tds.cam.Scale,
			ebiten.ActualFPS(),
			0.0,
			0.0,
		),
	)
}
