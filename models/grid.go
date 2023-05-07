package models

type Block struct {
	x int
	y int
}

type Grid struct {
	size   int
	blocks [][]Block
}

func NewGrid(size int) Grid {
	g := Grid{size: size}
	g.blocks = make([][]Block, size)
	for i := 0; i < size; i++ {
		g.blocks[i] = make([]Block, size)
	}
	return g
}

func (g *Grid) FillGrid() {
	for i := 0; i < g.size; i++ {
		for j := 0; j < g.size; j++ {
			b := Block{x: i, y: j}
			g.blocks[i][j] = b
		}
	}
}
