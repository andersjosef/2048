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

// Gamestates Enum style
type GameState int

const (
	StateRunning GameState = iota + 1
	StateMainMenu
)

type Game struct {
	board             *Board
	screenControl     *ScreenControl
	animation         *Animation
	menu              *Menu
	input             *Input
	state             GameState //if game is in menu. running, end etc 1: running
	score             int
	shouldClose       bool
	scale             float64
	screenSizeChanged bool
	darkMode          bool
}

func NewGame() (*Game, error) {
	// init game struct
	g := &Game{
		state:             StateMainMenu, // 2: main menu to start
		shouldClose:       false,         // if yes will close the game
		scale:             ebiten.Monitor().DeviceScaleFactor(),
		screenSizeChanged: false,
		darkMode:          false,
	}

	var err error

	// initialize new board
	g.animation = InitAnimation(g)
	g.screenControl = InitScreenControl(g)
	g.board, err = NewBoard(g)
	g.menu = NewMenu(g)
	g.input = InitInput(g)

	// initialize text
	initText(g)

	if err != nil {
		return nil, err
	}
	return g, nil
}

func (g *Game) Update() error {
	g.input.UpdateInput(g.board)
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
	if g.darkMode {
		screen.Fill(getColor(screenBackgroundColorDarkMode))
	} else {
		screen.Fill(getColor(screenBackgroundColorDefault))
	}
	switch g.state {
	case 1: //game is running loop
		if g.animation.isAnimating { // show animation
			g.animation.DrawAnimation(screen)
		} else { // draw normal borad
			g.board.drawBoard(screen)
		}
		DrawScore(screen, g)
	case 2: //game is in menu
		g.menu.DrawMenu(screen)
	}
}

func (game *Game) Layout(_, _ int) (int, int) { panic("use Ebitengine >=v2.5.0") }
func (g *Game) LayoutF(logicWinWidth, logicWinHeight float64) (float64, float64) {
	// scale := ebiten.DeviceScaleFactor()
	canvasWidth := math.Ceil(logicWinWidth * g.scale)
	canvasHeight := math.Ceil(logicWinHeight * g.scale)
	return canvasWidth, canvasHeight
}

func DrawScore(screen *ebiten.Image, g *Game) {
	myFont := mplusNormalFontSmaller

	//TODO make more dynamig IG
	margin := 10
	var shadowOffsett int = 2
	var score_text string = fmt.Sprintf("%v", g.score)

	text.Draw(screen, score_text, myFont,
		(shadowOffsett + margin),
		shadowOffsett+margin+text.BoundString(myFont, score_text).Dy(),
		color.Black)
	text.Draw(screen, score_text, myFont,
		margin,
		margin+text.BoundString(myFont, score_text).Dy(),
		color.White)
}
