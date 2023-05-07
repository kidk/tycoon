package graphics

type ImageCache struct {
	images map[string]*Image
}

func NewImageCache() ImageCache {
	return ImageCache{
		images: make(map[string]*Image),
	}
}

func (ic *ImageCache) GetImage(filePath string) *Image {
	if image, ok := ic.images[filePath]; ok {
		return image
	}

	image := NewImage(filePath)
	ic.images[filePath] = &image
	return &image
}
