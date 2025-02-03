package twenty48

import (
	"fmt"
	"image/color"

	"github.com/andersjosef/2048/twenty48/renderer"
	"github.com/andersjosef/2048/twenty48/shadertools"
	"github.com/andersjosef/2048/twenty48/theme"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

/* variables and constants */
const (
	logicalWidth  int = 640
	logicalHeight int = 480
	BOARDSIZE     int = 4
)

// Gamestates Enum style
type GameState int

const (
	StateRunning GameState = iota + 1
	StateMainMenu
	StateInstructions
	StateGameOver
)

type Game struct {
	board             *Board
	screenControl     *ScreenControl
	animation         *Animation
	menu              *Menu
	input             *Input
	buttonManager     *ButtonManager
	fontSet           *theme.FontSet
	renderer          *renderer.Renderer
	state             GameState // Game is in menu, running, etc
	previousState     GameState
	score             int
	shouldClose       bool // If yes will close the game
	scale             float64
	screenSizeChanged bool
	darkMode          bool
	currentTheme      theme.Theme
	gameOver          bool
}

func NewGame() (*Game, error) {
	// init game struct
	g := &Game{
		state:             StateMainMenu,
		previousState:     StateMainMenu,
		shouldClose:       false,
		scale:             ebiten.Monitor().DeviceScaleFactor(),
		screenSizeChanged: false,
		darkMode:          true,
	}

	if g.darkMode {
		g.currentTheme = theme.DarkTheme
	} else {
		g.currentTheme = theme.DefaultTheme

	}

	var err error

	// initialize text
	g.fontSet, err = theme.InitFonts(g.scale)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize fonts: %v", err)
	}

	// initialize new board
	g.animation = InitAnimation(g)
	g.screenControl = InitScreenControl(g)
	g.board, err = NewBoard(g)
	g.renderer = renderer.InitRenderer(g.fontSet)
	g.menu = NewMenu(g)
	g.input = InitInput(g)
	g.buttonManager = InitButtonManager(g)

	if err != nil {
		return nil, err
	}

	ebiten.SetWindowSize(logicalWidth*int(g.scale), logicalHeight*int(g.scale))
	return g, nil
}

func (g *Game) Update() error {
	g.input.UpdateInput(g.board)

	if g.shouldClose { // quit game check
		return ebiten.Termination
	}
	// if g.screenSizeChanged {
	// 	g.screenControl.ChangeBoardPosition()
	// }
	shadertools.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(g.currentTheme.ColorScreenBackground)
	switch g.state {
	case StateRunning: //game is running loop
		if g.animation.isAnimating { // show animation
			g.animation.DrawAnimation(screen)
		} else { // draw normal borad
			g.board.drawBoard(screen)
		}
		DrawScore(screen, g)
	case StateMainMenu, StateInstructions: //game is in menu
		g.menu.DrawMenu(screen)

	case StateGameOver:
		g.DrawGameOverScreen(screen)

	}
	g.buttonManager.drawButtons(screen)
}

// Layout defines the logical screen size and how it scales to the actual screen.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return logicalWidth, logicalHeight
}

func DrawScore(screen *ebiten.Image, g *Game) {
	myFont := g.fontSet.Smaller

	//TODO make more dynamig IG
	margin := 10
	var shadowOffsett int = 2
	var score_text string = fmt.Sprintf("%v", g.score)

	shadowOpt := &text.DrawOptions{}

	shadowOpt.GeoM.Translate(
		float64((shadowOffsett + margin)),
		0)
	shadowOpt.ColorScale.ScaleWithColor(color.Black)

	text.Draw(screen, score_text, myFont, shadowOpt)

	mainOpt := &text.DrawOptions{}
	mainOpt.GeoM.Translate(
		float64(margin),
		0)
	mainOpt.ColorScale.ScaleWithColor(color.White)

	text.Draw(screen, score_text, myFont,
		mainOpt)
}

// For reinitializing a font with a higher dpi
func (g *Game) updateFonts() {
	var err error
	g.fontSet, err = theme.InitFonts(g.scale)
	if err != nil {
		fmt.Println("Error changing fontsiz")
	}
}
