package twenty48

import (
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Animation struct {
	isAnimating     bool
	arrayOfChange   [BOARDSIZE][BOARDSIZE]int
	game            *Game
	currentDir      string
	animationLength float32           //seconds
	directionMap    map[string][2]int // multiply this to get x y movement of tiles
	startTime       time.Time
}

func InitAnimation(g *Game) *Animation {
	a := &Animation{
		isAnimating:     false,
		game:            g,
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

func (a *Animation) ResetArray() {
	a.arrayOfChange = [BOARDSIZE][BOARDSIZE]int{}
}

func (a *Animation) DrawAnimation(screen *ebiten.Image) {
	// Draw the backgroundimage of the game
	screen.DrawImage(a.game.board.boardImage, a.game.board.boardImageOptions)

	// Calculate animation progress based on time since start
	timeSinceStart := time.Since(a.startTime)
	progress := float32(timeSinceStart.Seconds()) / a.animationLength
	if progress > 1 {
		progress = 1 // Cap at 100%
	}

	// Draw tiles for animation
	for y := 0; y < len(a.game.board.board); y++ {
		for x := 0; x < len(a.game.board.board[0]); x++ {
			var (
				movingDistX float32 = progress * float32(a.directionMap[a.currentDir][0]) * float32(BOARDSIZE-1)
				movingDistY float32 = progress * float32(a.directionMap[a.currentDir][1]) * float32(BOARDSIZE-1)
			)
			if math.Abs(float64(movingDistX)) >= float64(a.arrayOfChange[y][x]) || math.Abs(float64(movingDistY)) >= float64(a.arrayOfChange[y][x]) {
				movingDistX = float32(a.directionMap[a.currentDir][0]) * float32(a.arrayOfChange[y][x])
				movingDistY = float32(a.directionMap[a.currentDir][1]) * float32(a.arrayOfChange[y][x])
			}
			a.game.board.DrawTile(screen, a.game.board.sizes.startPosX, a.game.board.sizes.startPosY, x, y, a.game.board.boardBeforeChange[y][x], movingDistX, movingDistY)
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
