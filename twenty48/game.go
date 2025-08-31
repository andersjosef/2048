package twenty48

import (
	"fmt"
	"image/color"
	"math"

	"github.com/andersjosef/2048/twenty48/animations"
	co "github.com/andersjosef/2048/twenty48/constants"
	"github.com/andersjosef/2048/twenty48/eventhandler"
	"github.com/andersjosef/2048/twenty48/renderer"
	"github.com/andersjosef/2048/twenty48/shadertools"
	"github.com/andersjosef/2048/twenty48/shared"
	"github.com/andersjosef/2048/twenty48/theme"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Game struct {
	board         *Board
	screenControl ScreenControl
	animation     *animations.Animation
	menu          Menu
	input         *Input
	buttonManager *ButtonManager
	fontSet       *theme.FontSet
	themePicker   *theme.ThemePicker
	renderer      *renderer.Renderer
	eventBus      *eventhandler.EventBus
	state         co.GameState // Game is in menu, running, etc
	previousState co.GameState
	score         int
	shouldClose   bool // If yes will close the game
	currentTheme  theme.Theme
	gameOver      bool
}

func NewGame() (*Game, error) {
	// init game struct
	g := &Game{
		state:         co.StateMainMenu,
		previousState: co.StateMainMenu,
		shouldClose:   false,
	}

	g.eventBus = eventhandler.NewEventBus()
	g.themePicker = theme.NewThemePicker()
	g.currentTheme = g.themePicker.GetCurrentTheme()
	g.screenControl = NewScreenControl(g)

	// initialize text
	var err error
	g.fontSet, err = theme.InitFonts(g.screenControl.GetScale())
	if err != nil {
		return nil, fmt.Errorf("failed to initialize fonts: %v", err)
	}

	// initialize new board
	g.board, err = NewBoard(g)
	g.animation = animations.InitAnimation(g.board)
	g.renderer = renderer.InitRenderer(g.fontSet)
	g.input = InitInput(g)
	g.buttonManager = InitButtonManager(g)
	g.menu = NewMenu(g)

	if err != nil {
		return nil, err
	}

	ebiten.SetWindowSize(
		co.LOGICAL_WIDTH*int(g.screenControl.GetScale()),
		co.LOGICAL_HEIGHT*int(g.screenControl.GetScale()),
	)

	g.registerEvents()
	return g, nil
}

func (g *Game) Update() error {
	g.eventBus.Dispatch()
	g.input.UpdateInput(g.board)

	if g.shouldClose { // quit game check
		return ebiten.Termination
	}

	shadertools.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(g.currentTheme.ColorScreenBackground)
	switch g.state {
	case co.StateRunning: //game is running loop
		if g.animation.IsAnimating() { // show animation
			g.animation.Draw(screen)
		} else { // draw normal borad
			g.board.Draw(screen)
		}
		DrawScore(screen, g)
	}
	g.buttonManager.drawButtons(screen)
	g.menu.Draw(screen)
}
func (game *Game) Layout(_, _ int) (int, int) { panic("use Ebitengine >=v2.5.0") }
func (g *Game) LayoutF(logicWinWidth, logicWinHeight float64) (float64, float64) {
	scale := ebiten.Monitor().DeviceScaleFactor()
	canvasWidth := math.Ceil(logicWinWidth * scale)
	canvasHeight := math.Ceil(logicWinHeight * scale)
	return canvasWidth, canvasHeight
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
	g.fontSet, err = theme.InitFonts(g.screenControl.GetScale())
	if err != nil {
		fmt.Println("Error changing fontsiz")
	}
}

func (g *Game) registerEvents() {
	g.eventBus.Register(
		eventhandler.EventResetGame,
		func(_ eventhandler.Event) {
			g.score = 0
			g.state = co.StateMainMenu // Swap to main menu
			g.gameOver = false
			shadertools.ResetTimesMapsDissolve()

		},
	)
	g.eventBus.Register(
		eventhandler.EventMoveMade,
		func(e eventhandler.Event) {
			data, ok := e.Data.(shared.MoveData)
			if !ok {
				return
			}

			g.score += data.ScoreGain
			g.gameOver = data.IsGameOver
		},
	)
}
