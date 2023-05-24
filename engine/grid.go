package engine

type Block struct {
	x int
	y int

	Visual Visual
}

type Visual struct {
	Name string
}

type Grid struct {
	Size   int
	Blocks [][]Block
}

func NewGrid(size int) Grid {
	g := Grid{Size: size}
	g.Blocks = make([][]Block, size)
	for i := 0; i < size; i++ {
		g.Blocks[i] = make([]Block, size)
	}

	return g
}

func (g *Grid) FillGrid() {
	for i := 0; i < g.Size; i++ {
		for j := 0; j < g.Size; j++ {
			b := Block{x: i, y: j}
			name := "ground_grass"
			if i%2 == 0 || j%2 == 0 {
				name = "ground_grass_high"
			}
			b.Visual = Visual{
				Name: name,
			}
			g.Blocks[i][j] = b
		}
	}
}
