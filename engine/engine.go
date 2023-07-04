package engine

import "github.com/sirupsen/logrus"

const MAP_SIZE = 30

type Engine struct {
	// We have layers here that define ground, buildings, items, ..
	FloorGrid    BlockGrid
	BuildingGrid BlockGrid

	// Pathfinding
	GridPath *GridPath

	// We have bots: guests, employees, .. with certain priorities and treats
	Player *Player
}

func NewEngine() *Engine {
	// Initialise the map
	// Currently always same map
	floor, buildings := NewBasicMap(MAP_SIZE)

	// Initialise player and bots
	player := NewPlayer()

	// Init pathfinding
	grid := NewGridPath()
	grid.Process(floor, func(block Block) bool {
		return false
	})
	grid.Process(buildings, func(block Block) bool {
		return block.Type.Name == "wall_brown_up_left"
	})

	return &Engine{
		FloorGrid:    floor,
		BuildingGrid: buildings,
		Player:       player,

		GridPath: grid,
	}
}

func (engine Engine) MovePlayer(x int, y int) {
	if !engine.GridPath.IsOpen(x, y) {
		logrus.Infoln("Can't go there")
		return
	}

	from := engine.GridPath.Get(engine.Player.X, engine.Player.Y)
	to := engine.GridPath.Get(x, y)
	logrus.WithFields(logrus.Fields{
		"from": from,
		"to":   to,
	}).Traceln("Pathfinding start")

	path, distance, found := Path(from, to)
	logrus.WithFields(logrus.Fields{
		"found":    found,
		"distance": distance,
		"path":     path,
	}).Traceln("Pathfinding complete")
}
