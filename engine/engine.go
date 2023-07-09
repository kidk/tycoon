package engine

import "github.com/sirupsen/logrus"

const MAP_SIZE = 30

type Engine struct {
	// We have layers here that define ground, buildings, items, ..
	FloorGrid    *BlockGrid
	BuildingGrid *BlockGrid
	ItemGrid     *BlockGrid
	ZoneGrid     *BlockGrid

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
	floor, buildings, items, zones := NewBasicMap(MAP_SIZE)

	// Initialise player and bots
	player := NewPlayer()

	// Init pathfinding
	grid := NewGridPath()
	grid.Process(floor, func(block Block) bool {
		return false
	})
	grid.Process(buildings, func(block Block) bool {
		return block.IsBlocker()
	})

	return &Engine{
		FloorGrid:    floor,
		BuildingGrid: buildings,
		ItemGrid:     items,
		ZoneGrid:     zones,

		Player: player,

		GridPath:  grid,
		updateFPS: fps,
	}
}

func (engine *Engine) Update() {
	// Need a pointer here in order to pass-by-reference
	// https://stackoverflow.com/questions/73494229/ineffective-assignment-to-field-when-trying-to-update-a-struct-in-go

	if engine.count%engine.updateFPS == 0 {
		// Reset count
		engine.count = 0
		// We are currently located on the "new" tile
		engine.Player.From = engine.Player.To
		// Continue to move if there is a next position
		if len(engine.Player.PlayerPath) > 0 { // There is a future path
			// Take first path and convert to Gridpather to get next X,Y variables
			nextPos := engine.Player.PlayerPath[0].(GridPather) // https://yourbasic.org/golang/type-assertion-switch/
			// Save this as the nextPosition
			engine.Player.To = &GridCoord{nextPos.X, nextPos.Y}
			// Change array to remove nextPost from the PlayerPath
			engine.Player.PlayerPath = engine.Player.PlayerPath[1:]
		}
		// Update Player state
		engine.Player.UpdateState()
	}

	// Update count
	engine.count += 1

	// Update real time position
	engine.Player.CurrentPos = &FloatCoord{
		float64(engine.Player.From.X) + float64(engine.Player.To.X-engine.Player.From.X)*float64(engine.count)/float64(engine.updateFPS), // X-coord
		float64(engine.Player.From.Y) + float64(engine.Player.To.Y-engine.Player.From.Y)*float64(engine.count)/float64(engine.updateFPS), // Y-coord
	}

	// Only update the position of the Player ones every second (= FPS frames)
	if engine.count%engine.updateFPS > 0 {
		return
	}

	/*
		// Need a pointer here in order to pass-by-reference
		// https://stackoverflow.com/questions/73494229/ineffective-assignment-to-field-when-trying-to-update-a-struct-in-go

		if engine.count%engine.updateFPS == 0 {
			// Reset count
			engine.count = 0
			// We are currently located on the "new" tile
			engine.Player.X = engine.Player.NextX
			engine.Player.Y = engine.Player.NextY
			// Continue to move if there is a next position
			if len(engine.Player.PlayerPath) > 0 { // There is a future path
				// Take first path and convert to Gridpather to get next X,Y variables
				nextPos := engine.Player.PlayerPath[0].(GridPather) // https://yourbasic.org/golang/type-assertion-switch/
				// Save this as the nextPosition
				engine.Player.NextX = nextPos.X
				engine.Player.NextY = nextPos.Y
				// Change array to remove nextPost from the PlayerPath
				engine.Player.PlayerPath = engine.Player.PlayerPath[1:]
			}
		}

		// Update count
		engine.count += 1

		// Only update the position of the Player ones every second (= FPS frames)
		if engine.count%engine.updateFPS > 0 {
			return
		}
	*/
}

func (engine *Engine) MovePlayer(x int, y int) {
	if !engine.GridPath.IsOpen(x, y) {
		logrus.Infoln("Can't go there")
		return
	}

	from := engine.GridPath.Get(engine.Player.To.X, engine.Player.To.Y) // We use the future position for new path, if Player is currently not moving the future position will be the same as current position
	//from := engine.GridPath.Get(engine.Player.NextX, engine.Player.NextY)
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

	// Assign the found path to the player
	if found {
		// We do not store the first value in the array, as that is the position where we currently are
		// and we don't need this in our (future) path
		engine.Player.PlayerPath = path[1:]
	}
}
