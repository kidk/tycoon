package helpers

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type MouseHelper struct {
	pressed map[ebiten.MouseButton]bool
	X, Y    float64
}

func NewMouseHelper() *MouseHelper {
	return &MouseHelper{
		pressed: make(map[ebiten.MouseButton]bool),
	}
}

func (kh *MouseHelper) IsKeyPressed(key ebiten.MouseButton) bool {
	return ebiten.IsMouseButtonPressed(key)
}

func (kh *MouseHelper) IsKeyPressedOnce(key ebiten.MouseButton) bool {
	isPressed := ebiten.IsMouseButtonPressed(key)

	// If key does not exist in cache, create it we current value
	if _, ok := kh.pressed[key]; !ok {
		kh.pressed[key] = isPressed

		// This is the first time we see this key, so let's return the current value
		return isPressed
	}

	// Save the previous value
	wasPressed := kh.pressed[key]
	result := false

	// If key was not previously set, but is pressed now, return true for one time
	if !wasPressed && isPressed {
		result = true

		// Update the cache
		kh.pressed[key] = isPressed
	}

	return result
}

func (kh *MouseHelper) Update() {
	for key, pressed := range kh.pressed {
		if pressed && !ebiten.IsMouseButtonPressed(key) {
			kh.pressed[key] = false
		}
	}
}
