package data

import (
	"fmt"
	"github.com/kidk/tycoon/graphics"
)

func LoadSprites(spriteCache *graphics.SpriteCache) {
	floorsFile := "resources/Room_Builder_Floors_32x32.png"

	// Floors
	spriteCache.CreateSprite("floor_tiles_light", floorsFile, 32, 96, 64, 128)
	spriteCache.CreateSprite("floor_wood_light", floorsFile, 32, 416, 64, 448)

	// Walls
	loadWall(spriteCache, "brown", 0, 0)
	loadWall(spriteCache, "grey", 256, 0)
	loadWall(spriteCache, "silver", 512, 0)
}

func loadWall(spriteCache *graphics.SpriteCache, name string, ox int, oy int) {
	wallsFile := "resources/Room_Builder_3d_walls_32x32.png"

	spriteCache.CreateSprite(fmt.Sprintf("wall_%s_up_left", name), wallsFile, 96+ox, 0+oy, 128+ox, 32+oy)
	spriteCache.CreateSprite(fmt.Sprintf("wall_%s_up_middle", name), wallsFile, 96+ox, 64+oy, 128+ox, 96+oy)
	spriteCache.CreateSprite(fmt.Sprintf("wall_%s_up_right", name), wallsFile, 128+ox, 0+oy, 160+ox, 32+oy)
	spriteCache.CreateSprite(fmt.Sprintf("wall_%s_middle_left", name), wallsFile, 64+ox, 64+oy, 64+32+ox, 64+32+oy)
	spriteCache.CreateSprite(fmt.Sprintf("wall_%s_middle_right", name), wallsFile, 160+ox, 64+oy, 160+32+ox, 64+32+oy)
	spriteCache.CreateSprite(fmt.Sprintf("wall_%s_down", name), wallsFile, 96+ox, 64+oy, 128+ox, 96+oy)
}
