package graphics

import (
	"fmt"
	"os"
)

type SpriteCache struct {
	sprites    map[string]*Sprite
	imageCache ImageCache
}

func NewSpriteCache(imageCache ImageCache) SpriteCache {
	return SpriteCache{
		sprites:    make(map[string]*Sprite),
		imageCache: imageCache,
	}
}

func (sc *SpriteCache) CreateSprite(name string, filepath string, lx int, ly int, sx int, sy int) {
	image := sc.imageCache.GetImage(filepath)
	sprite := NewSprite(image.texture, lx, ly, sx, sy)
	sc.sprites[name] = &sprite
}

func (sc *SpriteCache) GetSprite(name string) *Sprite {
	if sprite, ok := sc.sprites[name]; ok {
		return sprite
	}

	fmt.Fprintf(os.Stderr, "Texture not found, create it first: %v\n", name)
	os.Exit(1)

	return nil
}
