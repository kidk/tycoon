package engine

type Player struct {
	X, Y int
}

func NewPlayer() *Player {
	return &Player{
		X: 5,
		Y: 5,
	}
}
