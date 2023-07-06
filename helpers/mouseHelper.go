package helpers

import (
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
	mouseImage := ebiten.NewImage(64, 64)
	if kh.activeTexture != nil {
		kh.activeTexture.Draw(mouseImage, 0, 0)
	} else {
		kh.mouseTexture.Draw(mouseImage, 0, 0)
	}

	return mouseImage
}

func (kh *MouseHelper) LeftClick() {
	if kh.activeType != "" {
		switch kh.activeType {
		case "building":
			kh.engine.BuildingGrid.Set(int(kh.X), int(kh.Y), kh.activeName)
		case "floor":
			kh.engine.FloorGrid.Set(int(kh.X), int(kh.Y), kh.activeName)
		}
	} else {
		kh.engine.MovePlayer(int(kh.X), int(kh.Y))
	}
}

func (kh *MouseHelper) MiddleClick() {

}

func (kh *MouseHelper) RightClick() {
	kh.activeTexture = nil
	kh.activeType = ""
}

func (kh *MouseHelper) SetCursor(typ string, block string) {
	kh.activeTexture, _ = kh.spriteCache.GetSprite(block)
	kh.activeName = block
	kh.activeType = typ
}
