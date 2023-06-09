package main

import (
	"errors"
	"fmt"
	"image/color"
	"math"

	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/kidk/tycoon/engine"
	"github.com/kidk/tycoon/graphics"
	"github.com/kidk/tycoon/helpers"
	"github.com/kidk/tycoon/ui"
	camera "github.com/melonfunction/ebiten-camera"
	"github.com/sirupsen/logrus"
)

type GameScreen struct {
	timing helpers.TimingHelper

	engine *engine.Engine

	ui *ui.UITestScreen

	floorGridRenderer    *graphics.GridRenderer
	buildingGridRenderer *graphics.GridRenderer
	itemsGridRenderer    *graphics.GridRenderer
	zonesGridRenderer    *graphics.GridRenderer

	cam      *camera.Camera
	keyboard *helpers.KeyboardHelper
	mouse    *helpers.MouseHelper

	buildMode bool

	playerRenderer graphics.NPCRenderer
}

func NewGameScreen(spriteCache *graphics.SpriteCache) Screen {
	timing := *helpers.NewTimingHelper()
	timing.Disabled = true
	timing.Start("NewGameScreen")
	defer timing.Stop("NewGameScreen")

	engine := engine.NewEngine(ebiten.DefaultTPS / 2)
	mouse := helpers.NewMouseHelper(*spriteCache, engine)

	// Set up UI and mouse listener
	ui := ui.NewUITestScreen(spriteCache, func(event *widget.WidgetMouseButtonPressedEventArgs) {
		if event.Button == ebiten.MouseButton0 {
			logrus.Debugln("Mouse left click")
			logrus.Debugf("Coordinates x: %.3f y: %.3f \n", mouse.X, mouse.Y)
			mouse.LeftClick()
		}
		if event.Button == ebiten.MouseButton1 {
			logrus.Debugln("Mouse middle click")
			logrus.Debugf("Coordinates x: %.3f y: %.3f \n", mouse.X, mouse.Y)
			mouse.MiddleClick()
		}
		if event.Button == ebiten.MouseButton2 {
			logrus.Debugln("Mouse right click")
			logrus.Debugf("Coordinates x: %.3f y: %.3f \n", mouse.X, mouse.Y)
			mouse.RightClick()
		}
	}, mouse)

	return &GameScreen{
		timing: timing,

		engine: engine,

		ui: ui,

		floorGridRenderer:    graphics.NewGridRenderer(spriteCache, engine.FloorGrid, 0, 0),
		buildingGridRenderer: graphics.NewGridRenderer(spriteCache, engine.BuildingGrid, 0, 0),
		itemsGridRenderer:    graphics.NewGridRenderer(spriteCache, engine.ItemGrid, 0, 0),
		zonesGridRenderer:    graphics.NewGridRenderer(spriteCache, engine.ZoneGrid, 0, 0),

		cam:      camera.NewCamera(1920, 1080, 500, 500, 0, 1),
		keyboard: helpers.NewKeyboardHelper(),
		mouse:    mouse,

		buildMode: false,

		playerRenderer: graphics.NewNPCRenderer(spriteCache, engine.Player, 0, 0),
	}
}

func (tds *GameScreen) Update(g *Game) error {
	tds.timing.Start("Update")
	defer tds.timing.Stop("Update")

	// Update engine
	tds.engine.Update()

	// Update UI
	tds.ui.Update()

	// Update player animation
	tds.playerRenderer.Update()

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
	tds.timing.Start("Draw")
	defer tds.timing.Stop("Draw")

	// Clear camera surface
	tds.cam.Surface.Clear()
	tds.cam.Surface.Fill(color.RGBA{255, 128, 128, 255})

	// Shift camera

	// Draw ground
	groundOps := &ebiten.DrawImageOptions{}
	ground := tds.floorGridRenderer.Draw(screen)
	tds.cam.Surface.DrawImage(ground, tds.cam.GetTranslation(groundOps, 0, 0))
	defer ground.Dispose()

	// Draw buildings
	buildingOps := &ebiten.DrawImageOptions{}
	building := tds.buildingGridRenderer.Draw(screen)
	tds.cam.Surface.DrawImage(building, tds.cam.GetTranslation(buildingOps, 0, 0))
	defer building.Dispose()

	// Draw items
	itemsOps := &ebiten.DrawImageOptions{}
	items := tds.itemsGridRenderer.Draw(screen)
	tds.cam.Surface.DrawImage(items, tds.cam.GetTranslation(itemsOps, 0, 0))
	defer items.Dispose()

	// Draw unit
	playerOps := &ebiten.DrawImageOptions{}
	player := tds.playerRenderer.Draw(screen)
	tds.cam.Surface.DrawImage(player, tds.cam.GetTranslation(playerOps, (0+(tds.engine.Player.CurrentPos.X*32)), (-32+(tds.engine.Player.CurrentPos.Y*32)))) // correction because of unit size
	defer player.Dispose()

	// Draw zones
	zoneOps := &ebiten.DrawImageOptions{}
	zones := tds.zonesGridRenderer.Draw(screen)
	tds.cam.Surface.DrawImage(zones, tds.cam.GetTranslation(zoneOps, 0, 0))
	defer items.Dispose()

	// Hightlight tile under mouse
	mouseX, mouseY := tds.cam.GetCursorCoords()
	tds.mouse.X = math.Floor(float64(mouseX) / 32.0)
	tds.mouse.Y = math.Floor(float64(mouseY) / 32.0)

	mouseOps := &ebiten.DrawImageOptions{}
	mouse := tds.mouse.Draw(screen)
	tds.cam.Surface.DrawImage(mouse, tds.cam.GetTranslation(mouseOps, (tds.mouse.X-tds.mouse.OffsetX)*32, (tds.mouse.Y-tds.mouse.OffsetY)*32))
	defer mouse.Dispose()

	// Draw to screen and zoom
	tds.cam.Blit(screen)

	// Draw UI last
	tds.ui.Draw(screen)

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
			tds.mouse.X,
			tds.mouse.Y,
		),
	)
}
