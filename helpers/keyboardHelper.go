package helpers

import "github.com/hajimehoshi/ebiten/v2"

type KeyboardHelper struct {
	pressed map[ebiten.Key]bool
}

func NewKeyboardHelper() *KeyboardHelper {
	return &KeyboardHelper{
		pressed: make(map[ebiten.Key]bool),
	}
}

func (kh *KeyboardHelper) IsKeyPressed(key ebiten.Key) bool {
	return ebiten.IsKeyPressed(key)
}

func (kh *KeyboardHelper) IsKeyPressedOnce(key ebiten.Key) bool {
	isPressed := ebiten.IsKeyPressed(key)

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

func (kh *KeyboardHelper) Update() {
	for key, pressed := range kh.pressed {
		if pressed && !ebiten.IsKeyPressed(key) {
			print("setting key to false")
			kh.pressed[key] = false
		}
	}
}
