package engine

func NewBasicMap(size int) (*BlockGrid, *BlockGrid) {
	floor := NewGrid(size)
	floor.FillGrid(func(x int, y int) string {
		name := "ground_grass"
		if x%2 == 0 || y%2 == 0 {
			name = "ground_grass_high"
		}
		return name
	})

	buildings := NewGrid(size)
	buildings.FillGrid(func(x int, y int) string {
		return "empty"
	})

	buildHouse(buildings, 5, 10, 7)
	buildHouse(buildings, 5-1+7, 12, 5)
	buildHouse(buildings, 5-1+7+5, 12, 5)
	buildHouse(buildings, 5-1+7+5+5, 12, 5)

	return floor, buildings
}

func buildHouse(grid *BlockGrid, offsetX int, offsetY int, size int) {
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			if x == 0 || y == 0 || x == size-1 || y == size-1 {
				grid.Set(x+offsetX, y+offsetY, "wall_brown")
			} else {
				grid.Set(x+offsetX, y+offsetY, "floor_wood_light")
			}
		}
	}

	grid.Set(offsetX+2, offsetY+size-1, "floor_wood_light")
}
