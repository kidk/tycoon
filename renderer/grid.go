package renderer

type Block struct {
	x int
	y int

	Visual Visual
}

type Visual struct {
	Name string
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
			name := "ground_grass"
			if i%2 == 0 || j%2 == 0 {
				name = "ground_grass_high"
			}
			b.Visual = Visual{
				Name: name,
			}
			g.blocks[i][j] = b
		}
	}
}
