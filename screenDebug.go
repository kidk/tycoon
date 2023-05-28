package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kidk/tycoon/graphics"
)

type TextureDebugScreen struct {
	cabinet    graphics.AnimatedSprite
	animations []graphics.AnimatedSprite
}

func NewTextureDebugScreen(spriteCache *graphics.SpriteCache) Screen {
	sprite, _ := spriteCache.GetSprite("bathroom_cabinet_white")
	cabinet := graphics.NewAnimatedSprite(sprite, 10, 10)
	return &TextureDebugScreen{
		cabinet:    cabinet,
		animations: make([]graphics.AnimatedSprite, 100),
	}
}

func (tds *TextureDebugScreen) Update(g *Game) error {
	tds.cabinet.Update()

	return nil
}

func (tds *TextureDebugScreen) Draw(g *Game, screen *ebiten.Image) {
	tds.DrawExampleFloor(g, screen, "wood_light", 32, 32)
	tds.DrawExampleFloor(g, screen, "tiles_light", 32+(32*4), 32)

	tds.DrawRoomTiles(g, screen, "brown", 32, 160)
	tds.DrawRoomTiles(g, screen, "grey", 32+(4*32), 160)
	tds.DrawRoomTiles(g, screen, "silver", 32+(8*32), 160)

	tds.DrawExampleRoom(g, screen, "brown", 32, 256+32)
	tds.DrawExampleRoom(g, screen, "grey", 32+(6*32), 256+32)
	tds.DrawExampleRoom(g, screen, "silver", 32+(12*32), 256+32)

	tds.DrawCharacter(g, screen, "adam", 32, 480+(64*1))
	tds.DrawCharacter(g, screen, "alex", 32, 480+(64*2))
	tds.DrawCharacter(g, screen, "amelia", 32, 480+(64*3))
	tds.DrawCharacter(g, screen, "ash", 32, 480+(64*4))
	tds.DrawCharacter(g, screen, "bob", 32, 480+(64*5))
	tds.DrawCharacter(g, screen, "bruce", 32, 480+(64*6))
	tds.DrawCharacter(g, screen, "chef_alex", 32, 480+(64*7))
	tds.DrawCharacter(g, screen, "chef_lucy", 256, 480+(64*1))
	tds.DrawCharacter(g, screen, "chef_molly", 256, 480+(64*2))
	tds.DrawCharacter(g, screen, "cleaner_boy", 256, 480+(64*3))
	tds.DrawCharacter(g, screen, "cleaner_girl", 256, 480+(64*4))
	tds.DrawCharacter(g, screen, "conference_man", 256, 480+(64*5))
	tds.DrawCharacter(g, screen, "conference_woman", 256, 480+(64*6))
	tds.DrawCharacter(g, screen, "samuel", 256, 480+(64*7))

	tds.cabinet.Draw(screen, 32, 512-32)
}

func (tds *TextureDebugScreen) DrawRoomTiles(g *Game, screen *ebiten.Image, name string, ox float64, oy float64) {
	wallUpLeft, _ := g.SpriteCache.GetSprite(fmt.Sprintf("wall_%s_up_left", name))
	wallUpMiddle, _ := g.SpriteCache.GetSprite(fmt.Sprintf("wall_%s_up_middle", name))
	wallUpRight, _ := g.SpriteCache.GetSprite(fmt.Sprintf("wall_%s_up_right", name))
	wallMiddleLeft, _ := g.SpriteCache.GetSprite(fmt.Sprintf("wall_%s_middle_left", name))
	wallMiddleRight, _ := g.SpriteCache.GetSprite(fmt.Sprintf("wall_%s_middle_right", name))
	wallDown, _ := g.SpriteCache.GetSprite(fmt.Sprintf("wall_%s_down", name))

	wallUpLeft.Draw(screen, ox+(32*0), oy+(32*0))
	wallUpMiddle.Draw(screen, ox+(32*1), oy+(32*0))
	wallUpRight.Draw(screen, ox+(32*2), oy+(32*0))
	wallMiddleLeft.Draw(screen, ox+(32*0), oy+(32*1))
	wallMiddleRight.Draw(screen, ox+(32*2), oy+(32*1))
	wallDown.Draw(screen, ox+(32*0), oy+(32*2))
	wallDown.Draw(screen, ox+(32*1), oy+(32*2))
	wallDown.Draw(screen, ox+(32*2), oy+(32*2))
}

func (tds *TextureDebugScreen) DrawExampleRoom(g *Game, screen *ebiten.Image, name string, ox float64, oy float64) {
	wallUpLeft, _ := g.SpriteCache.GetSprite(fmt.Sprintf("wall_%s_up_left", name))
	wallUpMiddle, _ := g.SpriteCache.GetSprite(fmt.Sprintf("wall_%s_up_middle", name))
	wallUpRight, _ := g.SpriteCache.GetSprite(fmt.Sprintf("wall_%s_up_right", name))
	wallMiddleLeft, _ := g.SpriteCache.GetSprite(fmt.Sprintf("wall_%s_middle_left", name))
	wallMiddleRight, _ := g.SpriteCache.GetSprite(fmt.Sprintf("wall_%s_middle_right", name))
	wallDown, _ := g.SpriteCache.GetSprite(fmt.Sprintf("wall_%s_down", name))

	wallUpLeft.Draw(screen, ox+(32*0), oy+(32*0))
	wallUpMiddle.Draw(screen, ox+(32*1), oy+(32*0))
	wallUpMiddle.Draw(screen, ox+(32*2), oy+(32*0))
	wallUpMiddle.Draw(screen, ox+(32*3), oy+(32*0))
	wallUpRight.Draw(screen, ox+(32*4), oy+(32*0))
	wallMiddleLeft.Draw(screen, ox+(32*0), oy+(32*1))
	wallMiddleRight.Draw(screen, ox+(32*4), oy+(32*1))
	wallMiddleLeft.Draw(screen, ox+(32*0), oy+(32*2))
	wallMiddleRight.Draw(screen, ox+(32*4), oy+(32*2))
	wallMiddleLeft.Draw(screen, ox+(32*0), oy+(32*3))
	wallMiddleRight.Draw(screen, ox+(32*4), oy+(32*3))
	wallDown.Draw(screen, ox+(32*0), oy+(32*4))
	wallDown.Draw(screen, ox+(32*1), oy+(32*4))
	wallDown.Draw(screen, ox+(32*2), oy+(32*4))
	wallDown.Draw(screen, ox+(32*3), oy+(32*4))
	wallDown.Draw(screen, ox+(32*4), oy+(32*4))

	// Floor
	tds.DrawExampleFloor(g, screen, "wood_light", ox+(32*1), oy+(32*1))
}

func (tds *TextureDebugScreen) DrawExampleFloor(g *Game, screen *ebiten.Image, name string, ox float64, oy float64) {
	floor, _ := g.SpriteCache.GetSprite(fmt.Sprintf("floor_%s", name))

	floor.Draw(screen, ox+(32*0), oy+(32*0))
	floor.Draw(screen, ox+(32*1), oy+(32*0))
	floor.Draw(screen, ox+(32*2), oy+(32*0))
	floor.Draw(screen, ox+(32*0), oy+(32*1))
	floor.Draw(screen, ox+(32*1), oy+(32*1))
	floor.Draw(screen, ox+(32*2), oy+(32*1))
	floor.Draw(screen, ox+(32*0), oy+(32*2))
	floor.Draw(screen, ox+(32*1), oy+(32*2))
	floor.Draw(screen, ox+(32*2), oy+(32*2))
}

func (tds *TextureDebugScreen) DrawCharacter(g *Game, screen *ebiten.Image, name string, ox float64, oy float64) {
	characterRight, _ := g.SpriteCache.GetSprite(fmt.Sprintf("character_%s_idle_right", name))
	characterUp, _ := g.SpriteCache.GetSprite(fmt.Sprintf("character_%s_idle_up", name))
	characterLeft, _ := g.SpriteCache.GetSprite(fmt.Sprintf("character_%s_idle_left", name))
	characterDown, _ := g.SpriteCache.GetSprite(fmt.Sprintf("character_%s_idle_down", name))

	characterRight.Draw(screen, ox, oy)
	characterUp.Draw(screen, ox+32, oy)
	characterLeft.Draw(screen, ox+64, oy)
	characterDown.Draw(screen, ox+96, oy)
}
