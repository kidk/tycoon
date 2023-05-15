package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kidk/tycoon/graphics"
)

type TextureDebugScreen struct {
	cabinet graphics.AnimatedSprite
}

func NewTextureDebugScreen(spriteCache *graphics.SpriteCache) Screen {
	sprite := spriteCache.GetSprite("bathroom_cabinet_white")
	cabinet := graphics.NewAnimatedSprite(sprite, 10, 10)
	return &TextureDebugScreen{
		cabinet: cabinet,
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

	tds.cabinet.Draw(screen, 32, 512)
}

func (tds *TextureDebugScreen) DrawRoomTiles(g *Game, screen *ebiten.Image, name string, ox float64, oy float64) {
	wallUpLeft := g.SpriteCache.GetSprite(fmt.Sprintf("wall_%s_up_left", name))
	wallUpMiddle := g.SpriteCache.GetSprite(fmt.Sprintf("wall_%s_up_middle", name))
	wallUpRight := g.SpriteCache.GetSprite(fmt.Sprintf("wall_%s_up_right", name))
	wallMiddleLeft := g.SpriteCache.GetSprite(fmt.Sprintf("wall_%s_middle_left", name))
	wallMiddleRight := g.SpriteCache.GetSprite(fmt.Sprintf("wall_%s_middle_right", name))
	wallDown := g.SpriteCache.GetSprite(fmt.Sprintf("wall_%s_down", name))

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
	wallUpLeft := g.SpriteCache.GetSprite(fmt.Sprintf("wall_%s_up_left", name))
	wallUpMiddle := g.SpriteCache.GetSprite(fmt.Sprintf("wall_%s_up_middle", name))
	wallUpRight := g.SpriteCache.GetSprite(fmt.Sprintf("wall_%s_up_right", name))
	wallMiddleLeft := g.SpriteCache.GetSprite(fmt.Sprintf("wall_%s_middle_left", name))
	wallMiddleRight := g.SpriteCache.GetSprite(fmt.Sprintf("wall_%s_middle_right", name))
	wallDown := g.SpriteCache.GetSprite(fmt.Sprintf("wall_%s_down", name))

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
	floor := g.SpriteCache.GetSprite(fmt.Sprintf("floor_%s", name))

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
