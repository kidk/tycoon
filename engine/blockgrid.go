package engine

type BlockGrid struct {
	Size   int
	Blocks [][]Block
}

func NewGrid(size int) BlockGrid {
	g := BlockGrid{Size: size}
	g.Blocks = make([][]Block, size)
	for i := 0; i < size; i++ {
		g.Blocks[i] = make([]Block, size)
	}

	return g
}

func (g *BlockGrid) Set(x int, y int, name string) {
	b := Block{x: x, y: y}
	b.Type = Type{
		Name: name,
	}
	g.Blocks[x][y] = b
}

func (g *BlockGrid) FillGrid(nameFunc func(x int, y int) string) {
	for i := 0; i < g.Size; i++ {
		for j := 0; j < g.Size; j++ {
			b := Block{x: i, y: j}
			b.Type = Type{
				Name: nameFunc(i, j),
			}
			g.Blocks[i][j] = b
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
