package engine

type BlockGrid struct {
	Size   int
	Blocks [][]Block
}

func NewGrid(size int) *BlockGrid {
	g := &BlockGrid{Size: size}
	g.Blocks = make([][]Block, size)
	for i := 0; i < size; i++ {
		g.Blocks[i] = make([]Block, size)
	}

	return g
}

func (g *BlockGrid) Set(x int, y int, name string) {
	if len(name) > 6 && name[0:4] == "wall" {
		g.Blocks[x][y] = *WallFactory(x, y, name[5:])
		return
	}
	g.Blocks[x][y] = *BlockFactory(x, y, name)
}

func (g *BlockGrid) FillGrid(nameFunc func(x int, y int) string) {
	for x := 0; x < g.Size; x++ {
		for y := 0; y < g.Size; y++ {
			g.Blocks[x][y] = *BlockFactory(x, y, nameFunc(x, y))
		}
	}
}

func (g *BlockGrid) ForEach(process func(block Block)) {
	for index := range g.Blocks {
		for _, cel := range g.Blocks[index] {
			process(cel)
		}
	}
}
