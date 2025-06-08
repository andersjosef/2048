package twenty48

import (
	"math"
	"time"

	co "github.com/andersjosef/2048/twenty48/constants"
	"github.com/hajimehoshi/ebiten/v2"
)

type Animation struct {
	isAnimating     bool
	arrayOfChange   [co.BOARDSIZE][co.BOARDSIZE]int
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

func (a *Animation) Draw(screen *ebiten.Image) {
	// Draw the backgroundimage of the game
	// screen.DrawImage(a.game.board.boardImage, a.game.board.boardImageOptions)
	a.view.DrawBackgoundBoard(screen)

	// Calculate animation progress based on time since start
	timeSinceStart := time.Since(a.startTime)
	progress := min(float32(timeSinceStart.Seconds())/a.animationLength, 1)

	// Draw tiles for animation
	mWidth, mHeight := a.view.GetBoardDimentions()
	for y := range mHeight {
		for x := range mWidth {
			var (
				movingDistX float32 = progress * float32(a.directionMap[a.currentDir][0]) * float32(co.BOARDSIZE-1)
				movingDistY float32 = progress * float32(a.directionMap[a.currentDir][1]) * float32(co.BOARDSIZE-1)
			)
			if math.Abs(float64(movingDistX)) >= float64(a.arrayOfChange[y][x]) || math.Abs(float64(movingDistY)) >= float64(a.arrayOfChange[y][x]) {
				movingDistX = float32(a.directionMap[a.currentDir][0]) * float32(a.arrayOfChange[y][x])
				movingDistY = float32(a.directionMap[a.currentDir][1]) * float32(a.arrayOfChange[y][x])
			}
			a.view.DrawMovingMatrix(
				screen,
				x,
				y,
				movingDistX,
				movingDistY,
			)
			// a.view.DrawTile(
			// 	screen,
			// 	a.view.board.sizes.startPosX,
			// 	a.view.board.sizes.startPosY,
			// 	x,
			// 	y,
			// 	a.view.board.matrixBeforeChange[y][x],
			// 	movingDistX,
			// 	movingDistY,
			// )
		}
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
