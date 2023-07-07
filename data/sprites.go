package data

import (
	"fmt"

	"github.com/kidk/tycoon/graphics"
)

func LoadSprites(spriteCache *graphics.SpriteCache) {
	// Exceptions
	spriteCache.CreateSprite("error", "resources/error.png", 0, 0, 32, 32)

	// Ground
	groundFile := "resources/1_Terrains_and_Fences_32x32.png"
	spriteCache.CreateSprite("ground_grass", groundFile, 96, 256, 96+32, 256+32)
	spriteCache.CreateSprite("ground_grass_high", groundFile, 32, 224, 32+32, 224+32)

	// Floors
	floorsFile := "resources/Room_Builder_Floors_32x32.png"
	spriteCache.CreateSprite("floor_tiles_light", floorsFile, 32, 96, 64, 128)
	spriteCache.CreateSprite("floor_wood_light", floorsFile, 32, 416, 64, 448)

	// Mouse
	mouseFile := "resources/mouse.png"
	spriteCache.CreateSprite("mouse", mouseFile, 0, 0, 32, 32)

	// Walls
	loadWall(spriteCache, "brown", 0, 0)
	loadWall(spriteCache, "grey", 256, 0)
	loadWall(spriteCache, "silver", 512, 0)

	// Objects
	loadObjects(spriteCache)

	// Zones
	spriteCache.CreateSprite("zone_reception", "resources/zone_reception.png", 0, 0, 32, 32)
	spriteCache.CreateSprite("zone_room", "resources/zone_room.png", 0, 0, 32, 32)

	// Characters
	loadCharacter(spriteCache, "adam", "resources/Adam_32x32.png")
	loadCharacter(spriteCache, "alex", "resources/Alex_32x32.png")
	loadCharacter(spriteCache, "amelia", "resources/Amelia_32x32.png")
	loadCharacter(spriteCache, "ash", "resources/Ash_32x32.png")
	loadCharacter(spriteCache, "bob", "resources/Bob_32x32.png")
	loadCharacter(spriteCache, "bruce", "resources/Bruce_32x32.png")
	loadCharacter(spriteCache, "chef_alex", "resources/Chef_Alex_32x32.png")
	loadCharacter(spriteCache, "chef_lucy", "resources/Chef_Lucy_32x32.png")
	loadCharacter(spriteCache, "chef_molly", "resources/Chef_Molly_32x32.png")
	loadCharacter(spriteCache, "cleaner_boy", "resources/Cleaner_boy_32x32.png")
	loadCharacter(spriteCache, "cleaner_girl", "resources/Cleaner_girl_32x32.png")
	loadCharacter(spriteCache, "conference_man", "resources/Conference_man_32x32.png")
	loadCharacter(spriteCache, "conference_woman", "resources/Conference_woman_32x32.png")
	loadCharacter(spriteCache, "samuel", "resources/Samuel_32x32.png")
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

func loadObjects(spriteCache *graphics.SpriteCache) {
	spriteCache.CreateSprite("bathroom_cabinet_white", "resources/animated_bathroom_cabinet_white_empty_32x32.png", 0, 0, 32, 64)

	genericFile := "resources/1_Generic_32x32.png"
	spriteCache.CreateSprite("reception_desk", genericFile, 0, 512, 64, 32+512)
	spriteCache.CreateSprite("bed", genericFile, 320, 1888, 32+320, 64+1888)
}

func loadCharacter(spriteCache *graphics.SpriteCache, name string, file string) {
	spriteCache.CreateSprite(fmt.Sprintf("character_%s_idle_right", name), file, 0, 0, 32, 64)
	spriteCache.CreateSprite(fmt.Sprintf("character_%s_idle_up", name), file, 0+32, 0, 32+32, 64)
	spriteCache.CreateSprite(fmt.Sprintf("character_%s_idle_left", name), file, 0+64, 0, 32+64, 64)
	spriteCache.CreateSprite(fmt.Sprintf("character_%s_idle_down", name), file, 0+96, 0, 32+96, 64)

	spriteCache.CreateSprite(fmt.Sprintf("character_%s_move_right", name), file, 0, 128, 32, 192)
	spriteCache.CreateSprite(fmt.Sprintf("character_%s_move_up", name), file, 32*6, 128, 32*6+32, 192)
	spriteCache.CreateSprite(fmt.Sprintf("character_%s_move_left", name), file, 32*12, 128, 32*12+32, 192)
	spriteCache.CreateSprite(fmt.Sprintf("character_%s_move_down", name), file, 32*18, 128, 32*18+32, 192)
}
