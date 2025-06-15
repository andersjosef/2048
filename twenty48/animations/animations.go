package animations

import (
	"math"
	"time"

	co "github.com/andersjosef/2048/twenty48/constants"
	"github.com/hajimehoshi/ebiten/v2"
)

// Describes a single tile moving or merging from one tile to another
type MoveDelta struct {
	FromRow, FromCol int
	ToRow, ToCol     int
	ValueMoved       int
	Merged           bool
}

type Animation struct {
	isAnimating     bool
	ArrayOfChange   [co.BOARDSIZE][co.BOARDSIZE]int // TMP exposure!
	deltas          []MoveDelta
	view            View
	currentDir      string
	animationLength float32           //seconds
	directionMap    map[string][2]int // multiply this to get x y movement of tiles
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

	return a
}

func (a *Animation) Play(deltas []MoveDelta, dir string) {
	// a.deltas = deltas
	// a.currentDir = dir
	// a.isAnimating = true
	// a.startTime = time.Now()
}

func (a *Animation) Draw(screen *ebiten.Image) {
	// Draw the backgroundimage of the game
	a.view.DrawBackgoundBoard(screen)

	// Calculate animation progress based on time since start
	elapsed := float32(time.Since(a.startTime).Seconds())
	progress := float32(min(float64(elapsed/a.animationLength), 1))

	// // Draw tiles for animation
	// mWidth, mHeight := a.view.GetBoardDimentions()
	// for y := range mHeight {
	// 	for x := range mWidth {
	// 		var (
	// 			movingDistX float32 = progress * float32(a.directionMap[a.currentDir][0]) * float32(co.BOARDSIZE-1)
	// 			movingDistY float32 = progress * float32(a.directionMap[a.currentDir][1]) * float32(co.BOARDSIZE-1)
	// 		)
	// 		if math.Abs(float64(movingDistX)) >= float64(a.ArrayOfChange[y][x]) || math.Abs(float64(movingDistY)) >= float64(a.ArrayOfChange[y][x]) {
	// 			movingDistX = float32(a.directionMap[a.currentDir][0]) * float32(a.ArrayOfChange[y][x])
	// 			movingDistY = float32(a.directionMap[a.currentDir][1]) * float32(a.ArrayOfChange[y][x])
	// 		}
	// 		a.view.DrawMovingMatrix(
	// 			screen,
	// 			x,
	// 			y,
	// 			movingDistX,
	// 			movingDistY,
	// 		)
	// 	}
	// }

	for _, d := range a.deltas {
		// How many cells we move in X/Y
		dx := float32(d.ToCol - d.FromCol)
		dy := float32(d.ToRow - d.FromRow)

		// Progress * total cells, clamped so no overshoot
		movingDistX := float32(math.Copysign(math.Min(math.Abs(float64(dx*progress)), math.Abs(float64(dx))), float64(dx)))
		movingDistY := float32(math.Copysign(math.Min(math.Abs(float64(dy*progress)), math.Abs(float64(dy))), float64(dy)))

		// Draw at the original grid cell + offset
		a.view.DrawMovingMatrix(
			screen,
			d.FromCol,
			d.FromRow,
			movingDistX,
			movingDistY,
		)
	}

	if progress >= 1 {
		a.isAnimating = false
	}

}

// Use this function to activate animations
func (a *Animation) ActivateAnimation(direction string) {
	a.currentDir = direction
	a.isAnimating = true
	a.startTime = time.Now()
}

func (a *Animation) IsAnimating() bool {
	return a.isAnimating
}
