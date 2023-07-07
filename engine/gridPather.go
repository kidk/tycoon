package engine

type GridPath struct {
	field [][]GridPather
}

func NewGridPath() *GridPath {
	// Create path data
	field := make([][]GridPather, MAP_SIZE)
	for i := 0; i < MAP_SIZE; i++ {
		field[i] = make([]GridPather, MAP_SIZE)
	}

	gridPath := &GridPath{
		field: field,
	}

	for x, row := range field {
		for y := range row {
			field[x][y].world = gridPath
			field[x][y].X = x
			field[x][y].Y = y
		}
	}

	return gridPath
}

func (gp GridPath) Process(grid *BlockGrid, isBlocked func(block Block) bool) {
	grid.ForEach(func(block Block) {
		if isBlocked(block) {
			gp.field[block.x][block.y].Blocked = true
		}
	})
}

func (gp GridPath) Get(x int, y int) Pather {
	if x >= 0 && x < MAP_SIZE && y >= 0 && y < MAP_SIZE {
		return Pather(gp.field[x][y])
	}

	return nil
}

func (gp GridPath) IsOpen(x int, y int) bool {
	if x >= 0 && x < MAP_SIZE && y >= 0 && y < MAP_SIZE {
		return !gp.field[x][y].Blocked
	}

	return false
}

type GridPather struct {
	world   *GridPath
	X, Y    int
	Blocked bool
}

func NewGridPather(world *GridPath, x int, y int) Pather {
	return &GridPather{
		world: world,

		X: x,
		Y: y,

		Blocked: false,
	}
}

func (gp GridPather) PathNeighbors() []Pather {
	neighbors := []Pather{}

	// Left
	if gp.world.IsOpen(gp.X-1, gp.Y-0) {
		neighbors = append(neighbors, Pather(gp.world.field[gp.X-1][gp.Y-0]))
	}
	// Right
	if gp.world.IsOpen(gp.X+1, gp.Y-0) {
		neighbors = append(neighbors, Pather(gp.world.field[gp.X+1][gp.Y-0]))
	}
	// Up
	if gp.world.IsOpen(gp.X+0, gp.Y-1) {
		neighbors = append(neighbors, Pather(gp.world.field[gp.X+0][gp.Y-1]))
	}
	// Down
	if gp.world.IsOpen(gp.X+0, gp.Y+1) {
		neighbors = append(neighbors, Pather(gp.world.field[gp.X+0][gp.Y+1]))
	}

	return neighbors
}

func (gp GridPather) PathNeighborCost(to Pather) float64 {
	return 1
}

func (gp GridPather) PathEstimatedCost(to Pather) float64 {
	toT := to.(GridPather)
	absX := toT.X - gp.X
	if absX < 0 {
		absX = -absX
	}
	absY := toT.Y - gp.Y
	if absY < 0 {
		absY = -absY
	}
	return float64(absX + absY)
}
