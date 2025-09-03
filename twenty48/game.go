package twenty48

import (
	"fmt"
	"image/color"
	"math"

	"github.com/andersjosef/2048/twenty48/buttons"
	co "github.com/andersjosef/2048/twenty48/constants"
	"github.com/andersjosef/2048/twenty48/eventhandler"
	"github.com/andersjosef/2048/twenty48/input"
	"github.com/andersjosef/2048/twenty48/shadertools"
	"github.com/andersjosef/2048/twenty48/shared"
	"github.com/andersjosef/2048/twenty48/theme"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Game struct {
	board         Board
	screenControl ScreenControl
	animation     Animation
	menu          Menu
	renderer      Renderer
	input         *input.Input
	buttonManager *buttons.ButtonManager
	fontSet       *theme.FontSet
	themePicker   *theme.ThemePicker
	utils         Utils
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
	g.fontSet = theme.InitFonts(g.screenControl.GetScale())

	g.board = NewBoard(g)
	g.animation = NewAnimation(g)
	g.renderer = NewRenderer(g)
	g.utils = NewUtils()

	cmds := NewCommands(g)
	g.input = NewInput(g, cmds)
	g.buttonManager = NewButtonManager(g, cmds)
	g.input.GiveButtons(g.buttonManager)
	g.buttonManager.GiveInput(g.input)

	g.menu = NewMenu(g)
	ebiten.SetWindowSize(
		co.LOGICAL_WIDTH*int(g.screenControl.GetScale()),
		co.LOGICAL_HEIGHT*int(g.screenControl.GetScale()),
	)

	g.registerEvents()
	return g, nil
}

func (g *Game) Update() error {
	g.eventBus.Dispatch()
	g.input.UpdateInput()

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
		g.renderer.Draw(screen)
		DrawScore(screen, g)
	}
	g.buttonManager.Draw(screen)
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

	//TODO: make more dynamic
	margin := 10
	shadowOffsett := 2
	score_text := fmt.Sprintf("%v", g.score)

	getOpt := func(x, y float64, col color.Color) *text.DrawOptions {
		opt := &text.DrawOptions{}
		opt.GeoM.Translate(x, y)
		opt.ColorScale.ScaleWithColor(col)
		return opt
	}

	shadowOpt := getOpt(float64((shadowOffsett + margin)), 0, color.Black)
	text.Draw(screen, score_text, myFont, shadowOpt)

	mainOpt := getOpt(float64(margin), 0, color.White)
	text.Draw(screen, score_text, myFont, mainOpt)
}

// For reinitializing a font with a higher dpi
func (g *Game) updateFonts() {
	g.fontSet = theme.InitFonts(g.screenControl.GetScale())
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

// Temporary wrappers
func (g *Game) DrawMenu(screen *ebiten.Image) {
	g.menu.Draw(screen)
}

func (g *Game) DrawUI(screen *ebiten.Image) {
	g.buttonManager.Draw(screen)
}

func (g *Game) DrawRunning(screen *ebiten.Image) {
	g.renderer.Draw(screen)
	DrawScore(screen, g)
}
