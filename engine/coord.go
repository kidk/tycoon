package engine

type GridCoord struct {
	X, Y int
}

type FloatCoord struct {
	X, Y float64
}

func (gridCoord1 *GridCoord) minus(gridCoord2 *GridCoord) GridCoord {
	gridCoord := GridCoord{gridCoord1.X - gridCoord2.X, gridCoord1.Y - gridCoord2.Y}
	return gridCoord
}
