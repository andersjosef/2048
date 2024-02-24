package twenty48

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
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
	score int
}

func NewGame() (*Game, error) {
	g := &Game{
		state: 1,
	}

	var err error
	g.board, err = NewBoard()
	g.board.game = g
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
		// g.GetScore()

	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(getColor(BEIGE))
	g.board.drawBoard(screen)
	DrawScore(screen, g)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return SCREENWIDTH, SCREENHEIGHT
}

func DrawScore(screen *ebiten.Image, g *Game) {
	myFont := mplusNormalFontSmaller

	margin := 10
	var shadowOffsett int = 2
	var score_text string = fmt.Sprintf("%v", g.score)

	text.Draw(screen, score_text, myFont,
		shadowOffsett+margin,
		shadowOffsett+margin+text.BoundString(myFont, score_text).Dy(),
		color.Black)
	text.Draw(screen, score_text, myFont,
		10,
		10+text.BoundString(myFont, score_text).Dy(),
		color.White)
}
