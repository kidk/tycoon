package engine

import (
	"log"
)

const MAP_SIZE = 30

type Engine struct {
	// We have layers here that define ground, buildings, items, ..
	FloorGrid    BlockGrid
	BuildingGrid BlockGrid

	// Pathfinding
	GridPath *GridPath

	// We have bots: guests, employees, .. with certain priorities and treats
	Player *Player

	// counter, fps, ... for updating the "situation"
	count     int
	updateFPS int
}

func NewEngine(fps int) *Engine {
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

		GridPath:  grid,
		updateFPS: fps,
	}
}

func (engine *Engine) Update() {
	// Need a pointer here in order to pass-by-reference
	// https://stackoverflow.com/questions/73494229/ineffective-assignment-to-field-when-trying-to-update-a-struct-in-go

	// Only move ones every 60 seconds
	engine.count += 1
	if engine.count%engine.updateFPS > 0 {
		return
	}
	engine.count = 0

	if len(engine.Player.PlayerPath) > 0 {
		// take first path and convert to Gridpather to get X,Y variables
		nextPos := engine.Player.PlayerPath[0].(GridPather) // https://yourbasic.org/golang/type-assertion-switch/
		// move player to that location
		engine.Player.X = nextPos.X
		engine.Player.Y = nextPos.Y
		// change array to remove current (new) position
		engine.Player.PlayerPath = engine.Player.PlayerPath[1:]
	}
}

func (engine *Engine) MovePlayer(x int, y int) {
	if !engine.GridPath.IsOpen(x, y) {
		log.Println("Can't go there")
		return
	}

	from := engine.GridPath.Get(engine.Player.X, engine.Player.Y)
	to := engine.GridPath.Get(x, y)
	log.Println(from)
	log.Println(to)

	path, distance, found := Path(from, to)
	log.Println(found)
	log.Println(distance)
	log.Println(path)

	// Assign the found path to the player
	if found {
		// We do not store the first value in the array, as that is the position where we currently are
		// and we don't need this in our (future) path
		engine.Player.PlayerPath = path[1:]
	}
}
