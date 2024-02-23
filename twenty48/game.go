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
}

func NewGame() (*Game, error) {
	g := &Game{}

	var err error
	g.board, err = NewBoard()
	initText()
	if err != nil {
		return nil, err
	}
	return g, nil
}

func (g *Game) Update() error {
	m.UpdateInput(g.board)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
	screen.Fill(getColor(BEIGE))
	g.board.drawBoard(screen)
	// m.DrawInput(g.board)

}

//	func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
//		return SCREENWIDTH / 2, SCREENHEIGHT / 2
//	}
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return SCREENWIDTH, SCREENHEIGHT
}
