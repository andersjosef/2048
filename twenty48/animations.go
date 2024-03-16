package twenty48

import "fmt"

type Animation struct {
	isAnimating   bool
	arrayOfChange [BOARDSIZE][BOARDSIZE]int
	game          *Game
}

func InitAnimation(g *Game) *Animation {
	a := &Animation{
		isAnimating: false,
		game:        g,
	}

	return a
}

func (a *Animation) ResetArray() {
	fmt.Println("getting reset")
	a.arrayOfChange = [BOARDSIZE][BOARDSIZE]int{}
}
