package animations

import (
	"math"
	"time"

	co "github.com/andersjosef/2048/twenty48/constants"
	"github.com/andersjosef/2048/twenty48/eventhandler"
	"github.com/andersjosef/2048/twenty48/shared"
	"github.com/hajimehoshi/ebiten/v2"
)

type Animation struct {
	view            View
	isAnimating     bool
	deltas          []shared.MoveDelta
	currentDir      string
	animationLength float32           // Seconds
	directionMap    map[string][2]int // Multiply this to get x y movement of tiles
	startTime       time.Time
}

func InitAnimation(g View) *Animation {
	a := &Animation{
		isAnimating:     false,
		view:            g,
		animationLength: 0.20, // Animation duration in seconds
		directionMap: map[string][2]int{
			"UP":    {0, -1},
			"DOWN":  {0, 1},
			"LEFT":  {-1, 0},
			"RIGHT": {1, 0},
		},
	}

	a.view.Register(
		eventhandler.EventMoveMade,
		func(e eventhandler.Event) {
			moveData, ok := e.Data.(shared.MoveData)
			if !ok {
				return
			}

			a.Play(moveData.MoveDeltas, moveData.Dir)

		},
	)

	return a
}

func (a *Animation) Play(deltas []shared.MoveDelta, dir string) {
	for i, d := range deltas {
		nd := d
		switch dir {
		case "RIGHT":
			nd.FromCol = co.BOARDSIZE - 1 - d.FromCol
			nd.ToCol = co.BOARDSIZE - 1 - d.ToCol

		case "UP":
			nd.FromRow, nd.FromCol = d.FromCol, d.FromRow
			nd.ToRow, nd.ToCol = d.ToCol, d.ToRow

		case "DOWN":
			nd.FromRow, nd.FromCol = d.FromCol, d.FromRow
			nd.ToRow, nd.ToCol = d.ToCol, d.ToRow
			nd.FromRow = co.BOARDSIZE - 1 - nd.FromRow
			nd.ToRow = co.BOARDSIZE - 1 - nd.ToRow
		}
		deltas[i] = nd
	}

	a.deltas = deltas
	a.currentDir = dir
	a.isAnimating = true
	a.startTime = time.Now()
}

func (a *Animation) Draw(screen *ebiten.Image) {
	a.view.DrawBackgoundBoard(screen)

	elapsed := float32(time.Since(a.startTime).Seconds())
	progress := float32(min(float64(elapsed/a.animationLength), 1))

	// How far any tile would go if it had to move the full width/height
	fullDist := float32(co.BOARDSIZE-1) * progress

	for _, d := range a.deltas {

		// Signed cell‐deltas
		dxCells := d.ToCol - d.FromCol
		dyCells := d.ToRow - d.FromRow

		dirX := float32(sign(dxCells))
		dirY := float32(sign(dyCells))

		// How many cells this tile needs in each axis
		needX := float32(math.Abs(float64(dxCells)))
		needY := float32(math.Abs(float64(dyCells)))

		// Cap fullDist at each axis need
		moveX := dirX * min(fullDist, needX)
		moveY := dirY * min(fullDist, needY)

		a.view.DrawMovingMatrix(
			screen,
			d.FromCol,
			d.FromRow,
			moveX,
			moveY,
		)
	}

	if progress >= 1 {
		a.isAnimating = false
	}
}

// Helper to get sign of an int
func sign(n int) int {
	switch {
	case n < 0:
		return -1
	case n > 0:
		return +1
	default:
		return 0
	}
}

func (a *Animation) IsAnimating() bool {
	return a.isAnimating
}
