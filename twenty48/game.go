package twenty48

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

/* variables and constants */
const (
	SCREENWIDTH         int = 640
	SCREENHEIGHT        int = 480
	SCREENWIDTH_LAYOUT  int = 640 / 2
	SCREENHEIGHT_LAYOUT int = 480 / 2
	BOARDSIZE           int = 4
)

type Game struct {
	board *Board
	state int //if game is in menu. running, end etc 1: running
}

func NewGame() (*Game, error) {
	g := &Game{
		state: 1,
	}

	var err error
	g.board, err = NewBoard()
	initText()
	if err != nil {
		return nil, err
	}
	return g, nil
}

func (g *Game) Update() error {
	switch g.state {
	case 1: //game is running loop
		m.UpdateInput(g.board)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
	screen.Fill(getColor(BEIGE))
	g.board.drawBoard(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return SCREENWIDTH, SCREENHEIGHT
}
