package engine

func NewBasicMap(size int) (BlockGrid, BlockGrid) {
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
	buildings.Set(10, 10, "wall_brown_up_left")

	return floor, buildings
}
