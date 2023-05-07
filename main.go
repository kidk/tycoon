package main

import (
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/kidk/tycoon/data"
	"github.com/kidk/tycoon/graphics"
	"github.com/kidk/tycoon/models"
)

type Game struct {
	grid         models.Grid
	gridRenderer graphics.GridRenderer
	spriteCache  graphics.SpriteCache
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
	g.gridRenderer.Draw(screen)
	sprite := g.spriteCache.GetSprite("floor1")

	// Available sprites
	sprite.Draw(screen, 32, 32)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	logger.Print("Starting game")

	logger.Print("Setting window size and title")
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	g := &Game{}

	logger.Print("Creating grid")
	g.grid = models.NewGrid(10)
	g.grid.FillGrid()

	logger.Print("Creating image & sprite cache")
	imageCache := graphics.NewImageCache()
	g.spriteCache = graphics.NewSpriteCache(imageCache)

	logger.Print("Creating sprites")
	data.LoadSprites(&g.spriteCache)

	logger.Print("Creating grid renderer")
	g.gridRenderer = graphics.NewGridRenderer(g.grid)

	logger.Print("Everything is ready, starting loops")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
