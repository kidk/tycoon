package engine

type Player struct {
	X, Y       int
	PlayerPath []Pather
}

func NewPlayer() *Player {
	return &Player{
		X:          5,
		Y:          5,
		PlayerPath: nil,
	}
}
