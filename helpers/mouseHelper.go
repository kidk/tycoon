package helpers

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kidk/tycoon/engine"
	"github.com/kidk/tycoon/graphics"
)

type MouseHelper struct {
	spriteCache graphics.SpriteCache

	engine *engine.Engine

	pressed      map[ebiten.MouseButton]bool
	X, Y         float64
	mouseTexture *graphics.Sprite

	draggable              bool
	dragging               bool
	dragStartX, dragStartY float64
	OffsetX, OffsetY       float64

	activeTexture *graphics.Sprite
	activeName    string
	activeType    string
}

func NewMouseHelper(spriteCache graphics.SpriteCache, engine *engine.Engine) *MouseHelper {
	mouseTexture, _ := spriteCache.GetSprite("mouse")

	return &MouseHelper{
		spriteCache: spriteCache,

		engine: engine,

		pressed:      make(map[ebiten.MouseButton]bool),
		mouseTexture: mouseTexture,
	}
}

func (kh *MouseHelper) Draw(screen *ebiten.Image) *ebiten.Image {
	mouseImage := ebiten.NewImage(512, 512)
	if kh.activeTexture != nil {
		kh.activeTexture.Draw(mouseImage, 0, 0)
	} else {
		kh.mouseTexture.Draw(mouseImage, 0, 0)
	}

	kh.OffsetX = 0.0
	kh.OffsetY = 0.0

	if kh.activeType != "" && kh.dragging {
		sizeX := math.Abs(kh.X-kh.dragStartX) + 1
		sizeY := math.Abs(kh.Y-kh.dragStartY) + 1

		if kh.X > kh.dragStartX {
			kh.OffsetX = kh.X - kh.dragStartX
		}
		if kh.Y > kh.dragStartY {
			kh.OffsetY = kh.Y - kh.dragStartY
		}

		for x := 0.0; x < sizeX; x++ {
			for y := 0.0; y < sizeY; y++ {
				kh.activeTexture.Draw(mouseImage, x*32, y*32)
			}
		}
	}

	return mouseImage
}

func (kh *MouseHelper) LeftClick() {
	if kh.activeType != "" {
		if kh.dragging {
			kh.dragging = false
			sizeX := math.Abs(kh.X-kh.dragStartX) + 1
			sizeY := math.Abs(kh.Y-kh.dragStartY) + 1
			for x := 0; x < int(sizeX); x++ {
				for y := 0; y < int(sizeY); y++ {
					xCorr := int(math.Min(kh.X, kh.dragStartX)) + x
					yCorr := int(math.Min(kh.Y, kh.dragStartY)) + y
					kh.setBlock(xCorr, yCorr)
				}
			}
		} else {
			if kh.draggable {
				kh.dragging = true
				kh.dragStartX = kh.X
				kh.dragStartY = kh.Y
			} else {
				kh.setBlock(int(kh.X), int(kh.Y))
			}
		}
	} else {
		kh.engine.MovePlayer(int(kh.X), int(kh.Y))
	}
}

func (kh *MouseHelper) setBlock(x int, y int) {
	switch kh.activeType {
	case "building":
		kh.engine.BuildingGrid.Set(x, y, kh.activeName)
	case "floor":
		kh.engine.FloorGrid.Set(x, y, kh.activeName)
	case "item":
		kh.engine.ItemGrid.Set(x, y, kh.activeName)
	case "zone":
		kh.engine.ZoneGrid.Set(x, y, kh.activeName)
	}
}

func (kh *MouseHelper) MiddleClick() {

}

func (kh *MouseHelper) RightClick() {
	kh.activeTexture = nil
	kh.activeType = ""
}

func (kh *MouseHelper) SetCursor(typ string, block string, draggable bool) {
	kh.activeTexture, _ = kh.spriteCache.GetSprite(block)
	kh.activeName = block
	kh.activeType = typ
	kh.draggable = draggable
}
