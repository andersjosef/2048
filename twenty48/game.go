package twenty48

import (
	"fmt"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

/* variables and constants */
const (
	SCREENWIDTH  int = 640
	SCREENHEIGHT int = 480
	BOARDSIZE    int = 4
)

type Game struct {
	board             *Board
	state             int //if game is in menu. running, end etc 1: running
	score             int
	shouldClose       bool
	screenControl     *ScreenControl
	scale             float64
	screenSizeChanged bool
	animation         *Animation
}

func NewGame() (*Game, error) {
	// init game struct
	g := &Game{
		state:             2,     // 2: main menu to start
		shouldClose:       false, // if yes will close the game
		scale:             ebiten.DeviceScaleFactor(),
		screenSizeChanged: false,
	}

	var err error

	// initialize new board
	g.animation = InitAnimation(g)
	g.screenControl = InitScreenControl(g)
	g.board, err = NewBoard()
	g.board.game = g

	// initialize text
	initText()

	if err != nil {
		return nil, err
	}
	return g, nil
}

func (g *Game) Update() error {
	m.UpdateInput(g.board)
	switch g.state {
	case 1: //game is running loop
	case 2: //game is in menu
	}

	if g.shouldClose { // quit game check
		return ebiten.Termination
	}
	if g.screenSizeChanged {
		g.ChangeBoardPosition()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(getColor(BEIGE))
	switch g.state {
	case 1: //game is running loop
		if g.animation.isAnimating { // show animation
			g.animation.DrawAnimation(screen)
		} else { // draw normal borad
			g.board.drawBoard(screen)
		}
		DrawScore(screen, g)
	case 2: //game is in menu
		g.DrawMenu(screen)
	}
}

func (game *Game) Layout(_, _ int) (int, int) { panic("use Ebitengine >=v2.5.0") }
func (g *Game) LayoutF(logicWinWidth, logicWinHeight float64) (float64, float64) {
	scale := ebiten.DeviceScaleFactor()
	canvasWidth := math.Ceil(logicWinWidth * scale)
	canvasHeight := math.Ceil(logicWinHeight * scale)
	return canvasWidth, canvasHeight
}

func DrawScore(screen *ebiten.Image, g *Game) {
	myFont := mplusNormalFontSmaller

	//TODO make more dynamig IG
	margin := 10
	var shadowOffsett int = 2
	var score_text string = fmt.Sprintf("%v", g.score)

	text.Draw(screen, score_text, myFont,
		shadowOffsett+margin,
		shadowOffsett+margin+text.BoundString(myFont, score_text).Dy(),
		color.Black)
	text.Draw(screen, score_text, myFont,
		margin,
		margin+text.BoundString(myFont, score_text).Dy(),
		color.White)
}
