package engine

type State int

const (
	Idle State = iota
	Move_Down
	Move_Right
	Move_Up
	Move_Left
)

type Player struct {

	// Last known position on grid
	From GridCoord
	// Position in on which the player will be next (on grid)
	// This can be the same as X, Y if player is not moving
	To GridCoord
	// Exact currentPosition
	CurrentPos FloatCoord
	// Future path of the Player, as found by Path(from, to) with "from = nextX,nextY"
	// This path can be adapted at any time
	PlayerPath []Pather

	// State of the player
	// Mainly relevant to know how to draw the character
	State State
}

func NewPlayer() *Player {
	return &Player{
		From:       GridCoord{5, 5},
		To:         GridCoord{5, 5},
		CurrentPos: FloatCoord{5, 5},
		PlayerPath: nil,
		State:      Idle,
	}
}

func (player *Player) UpdateState() {
	delta := player.From.minus(&player.To)

	if delta.Y > 0 {
		player.State = Move_Down
	} else if delta.X > 0 {
		player.State = Move_Right
	} else if delta.Y < 0 {
		player.State = Move_Up
	} else if delta.X < 0 {
		player.State = Move_Left
	} else { // default case
		player.State = Idle
	}

}
