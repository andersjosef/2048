package twenty48

import (
	"fmt"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Animation struct {
	isAnimating          bool
	arrayOfChange        [BOARDSIZE][BOARDSIZE]int
	game                 *Game
	currentDir           string
	animationCounterOrig int
	animationCounter     int
	animationLength      float32           //seconds
	directionMap         map[string][2]int // multiply this to get x y movement of tiles
}

func InitAnimation(g *Game) *Animation {
	a := &Animation{
		isAnimating: false,
		game:        g,
	}

	a.animationLength = 0.25
	a.directionMap = map[string][2]int{
		"UP":    {0, -1},
		"DOWN":  {0, 1},
		"LEFT":  {-1, 0},
		"RIGHT": {1, 0},
	}

	return a
}

func (a *Animation) ResetArray() {
	fmt.Println("getting reset")
	a.arrayOfChange = [BOARDSIZE][BOARDSIZE]int{}
}

func (a *Animation) DrawAnimation(screen *ebiten.Image) {

	// fmt.Println("animation drawing")
	// fmt.Printf("Current direction %s\n %v\n", a.currentDir, a.directionMap[a.currentDir])

	// draw the backgroundimage of the game
	screen.DrawImage(a.game.board.board_image, a.game.board.board_image_options)

	// draw tiles
	for y := 0; y < len(a.game.board.board); y++ {
		for x := 0; x < len(a.game.board.board[0]); x++ {
			var (
				evenFlow float32 = (float32(a.animationCounterOrig) - float32(a.animationCounter)) / (float32(a.animationCounterOrig)) // to get even increase
				// movingDistX float32 = evenFlow * float32(a.directionMap[a.currentDir][0]) * float32(a.arrayOfChange[y][x])
				// movingDisty float32 = evenFlow * float32(a.directionMap[a.currentDir][1]) * float32(a.arrayOfChange[y][x])
				movingDistX float32 = evenFlow * float32(a.directionMap[a.currentDir][0]) * float32(BOARDSIZE-1)
				movingDistY float32 = evenFlow * float32(a.directionMap[a.currentDir][1]) * float32(BOARDSIZE-1)
			)
			if math.Abs(float64(movingDistX)) >= float64(a.arrayOfChange[y][x]) || math.Abs(float64(movingDistY)) >= float64(a.arrayOfChange[y][x]) {
				movingDistX = float32(a.directionMap[a.currentDir][0]) * float32(a.arrayOfChange[y][x])
				movingDistY = float32(a.directionMap[a.currentDir][1]) * float32(a.arrayOfChange[y][x])
			}
			// fmt.Println(a.animationCounterOrig, a.animationCounter, evenFlow)
			a.game.board.DrawTile(screen, start_pos_x, start_pos_y, x, y, a.game.board.board_before_change[y][x], movingDistX, movingDistY)
		}
	}

	// fmt.Println("animation over")
	a.animationCounter--
	if a.animationCounter < 0 {
		a.isAnimating = false
	}

}

func (a *Animation) ActivateAnimation(direction string) {
	a.currentDir = direction
	a.isAnimating = true
	a.animationCounterOrig = int(float32(60) * a.animationLength)
	a.animationCounter = a.animationCounterOrig

}
