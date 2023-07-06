package engine

import "fmt"

type Block struct {
	x int
	y int

	GetTexture func(*BlockGrid) string

	IsBlocker func() bool
}

type Type struct {
	Name string
}

func BlockFactory(x int, y int, name string) *Block {
	return &Block{
		x: x,
		y: y,

		GetTexture: func(grid *BlockGrid) string {
			return name
		},
		IsBlocker: func() bool {
			return false
		},
	}
}

func WallFactory(x int, y int, name string) *Block {
	return &Block{
		x: x,
		y: y,

		GetTexture: func(grid *BlockGrid) string {
			// up_left, up_middle, up_right, middle_left, middle_right, down

			return fmt.Sprintf("wall_%s_down", name)
		},
		IsBlocker: func() bool {
			return true
		},
	}
}
