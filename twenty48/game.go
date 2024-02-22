package twenty48

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

/* variables */

type Game struct {
	board *Board
}

func NewGame() (*Game, error) {
	g := &Game{}

	var err error
	g.board, err = NewBoard()

	if err != nil {
		return nil, err
	}
	return g, nil
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
	screen.Fill(getColor(BEIGE))
	g.board.drawBoard(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
