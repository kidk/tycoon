package data

import (
	"github.com/kidk/tycoon/graphics"
)

func LoadSprites(spriteCache *graphics.SpriteCache) {
	floorsFile := "resources/Room_Builder_Floors_32x32.png"

	// Floors
	spriteCache.CreateSprite("floor1", floorsFile, 32, 96, 63, 127)
}
