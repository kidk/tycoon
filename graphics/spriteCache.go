package graphics

import (
	"fmt"

	"github.com/sirupsen/logrus"
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
	// All textures should follow a certain size pattern, adding some logging to catch mistakes when drawing out the dimensions
	if lx%32 != 0 {
		logrus.Printf("Log %s in file %s lx not devidable by 32", name, filepath)
	}
	if ly%32 != 0 {
		logrus.Printf("Log %s in file %s ly not devidable by 32", name, filepath)
	}
	if sx%32 != 0 {
		logrus.Printf("Log %s in file %s sx not devidable by 32", name, filepath)
	}
	if sy%32 != 0 {
		logrus.Printf("Log %s in file %s sy not devidable by 32", name, filepath)
	}

	image := sc.imageCache.GetImage(filepath)
	sprite := NewSprite(image.texture, lx, ly, sx, sy)
	sc.sprites[name] = &sprite
}

func (sc *SpriteCache) GetSprite(name string) (*Sprite, error) {
	if sprite, ok := sc.sprites[name]; ok {
		return sprite, nil
	}

	return nil, fmt.Errorf("Texture not found, create it first: %v\n", name)
}
