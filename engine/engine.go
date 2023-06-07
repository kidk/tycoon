package engine

type Engine struct {
	// We have layers here that define ground, buildings, items, ..
	FloorGrid    BlockGrid
	BuildingGrid BlockGrid

	// We have bots: guests, employees, .. with certain priorities and treats

}

func NewEngine() *Engine {
	floor := NewGrid(30)
	floor.FillGrid(func(x int, y int) string {
		name := "ground_grass"
		if x%2 == 0 || y%2 == 0 {
			name = "ground_grass_high"
		}
		return name
	})

	buildings := NewGrid(30)
	buildings.FillGrid(func(x int, y int) string {
		return "empty"
	})
	buildings.Set(10, 10, "wall_brown_up_left")

	return &Engine{
		FloorGrid:    floor,
		BuildingGrid: buildings,
	}
}
