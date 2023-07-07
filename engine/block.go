package engine

import "fmt"

type Block struct {
	name string
	x    int
	y    int

	GetTexture func(*BlockGrid) string

	IsBlocker func() bool
}

type Type struct {
	Name string
}

func BlockFactory(x int, y int, name string) *Block {
	return &Block{
		name: name,
		x:    x,
		y:    y,

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
		name: name,
		x:    x,
		y:    y,

		GetTexture: func(grid *BlockGrid) string {
			// up_left, up_middle, up_right, middle_left, middle_right, down
			top := y-1 > 0 && len(grid.Blocks[x][y-1].name) > 4 && grid.Blocks[x][y-1].name[0:4] == "wall"
			down := y+1 < len(grid.Blocks) && len(grid.Blocks[x][y+1].name) > 4 && grid.Blocks[x][y+1].name[0:4] == "wall"
			left := x-1 > 0 && len(grid.Blocks[x-1][y].name) > 4 && grid.Blocks[x-1][y].name[0:4] == "wall"
			right := x+1 < len(grid.Blocks) && len(grid.Blocks[x+1][y].name) > 4 && grid.Blocks[x+1][y].name[0:4] == "wall"

			direction := "down"
			if !top && down && !left && right {
				direction = "up_left"
			}
			if left && right {
				direction = "up_middle"
			}
			if !top && down && left && !right {
				direction = "up_right"
			}
			if !left && !right {
				direction = "middle_right"
			}

			return fmt.Sprintf("wall_%s_%s", name[5:], direction)
		},
		IsBlocker: func() bool {
			return true
		},
	}
}
