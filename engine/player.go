package engine

type PlayerState int

const (
	Player_Idle PlayerState = iota
	Player_MoveDown
	Player_MoveRight
	Player_MoveUp
	Player_MoveLeft
)

type Player struct {

	// Last known position on grid
	From *GridCoord
	// Position in on which the player will be next (on grid)
	// This can be the same as X, Y if player is not moving
	To *GridCoord
	// Exact currentPosition
	CurrentPos *FloatCoord
	// Future path of the Player, as found by Path(from, to) with "from = nextX,nextY"
	// This path can be adapted at any time
	PlayerPath []Pather

	// State of the player
	// Mainly relevant to know how to draw the character
	State PlayerState
}

func NewPlayer() *Player {
	return &Player{
		From:       &GridCoord{5, 5},
		To:         &GridCoord{5, 5},
		CurrentPos: &FloatCoord{5, 5},
		PlayerPath: nil,
		State:      Player_Idle,
	}
}

func (player *Player) UpdateState() {
	delta := player.To.minus(player.From)

	if delta.Y > 0 {
		player.State = Player_MoveDown
	} else if delta.X > 0 {
		player.State = Player_MoveRight
	} else if delta.Y < 0 {
		player.State = Player_MoveUp
	} else if delta.X < 0 {
		player.State = Player_MoveLeft
	} else { // default case
		player.State = Player_Idle
	}

}
