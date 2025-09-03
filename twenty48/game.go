package twenty48

import (
	"fmt"
	"image/color"
	"math"

	"github.com/andersjosef/2048/twenty48/buttons"
	"github.com/andersjosef/2048/twenty48/commands"
	co "github.com/andersjosef/2048/twenty48/constants"
	"github.com/andersjosef/2048/twenty48/eventhandler"
	"github.com/andersjosef/2048/twenty48/input"
	"github.com/andersjosef/2048/twenty48/menu"
	"github.com/andersjosef/2048/twenty48/shadertools"
	"github.com/andersjosef/2048/twenty48/shared"
	"github.com/andersjosef/2048/twenty48/theme"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Game struct {
	d Deps

	board         Board
	screenControl ScreenControl
	animation     Animation
	Menu          *menu.Menu
	renderer      Renderer
	Input         *input.Input
	buttonManager *buttons.ButtonManager
	fontSet       *theme.FontSet
	themePicker   *theme.ThemePicker
	utils         Utils
	EventBus      *eventhandler.EventBus
	Cmds          *commands.Commands

	score        int
	shouldClose  bool // If yes will close the game
	currentTheme theme.Theme
	gameOver     bool
}

func NewGame(d Deps) (*Game, error) {
	// init game struct
	g := &Game{
		d:           d,
		shouldClose: false,
	}

	g.EventBus = eventhandler.NewEventBus()
	g.themePicker = theme.NewThemePicker()
	g.currentTheme = g.themePicker.GetCurrentTheme()
	g.screenControl = NewScreenControl(g)

	// initialize text
	g.fontSet = theme.InitFonts(g.screenControl.GetScale())

	g.board = NewBoard(g)
	g.animation = NewAnimation(g)
	g.renderer = NewRenderer(g)
	g.utils = NewUtils()

	g.Cmds = NewCommands(g)
	g.Input = NewInput(g, g.Cmds)
	g.buttonManager = NewButtonManager(g, g.Cmds)
	g.Input.GiveButtons(g.buttonManager)
	g.buttonManager.GiveInput(g.Input)

	g.Menu = NewMenu(g)
	ebiten.SetWindowSize(
		co.LOGICAL_WIDTH*int(g.screenControl.GetScale()),
		co.LOGICAL_HEIGHT*int(g.screenControl.GetScale()),
	)

	g.registerEvents()
	return g, nil
}

// Global update which is run regardless of state
func (g *Game) Update() error {
	g.EventBus.Dispatch()
	g.Input.UpdateInput()

	if g.shouldClose { // quit game check
		return ebiten.Termination
	}

	shadertools.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(g.currentTheme.ColorScreenBackground)
	g.buttonManager.Draw(screen)
	g.Menu.Draw(screen)
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
	g.EventBus.Register(
		eventhandler.EventResetGame,
		func(_ eventhandler.Event) {
			g.score = 0
			g.SetState(co.StateMainMenu) // Swap to main menu
			g.gameOver = false
			shadertools.ResetTimesMapsDissolve()

		},
	)
	g.EventBus.Register(
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
	g.Menu.Draw(screen)
}

func (g *Game) DrawUI(screen *ebiten.Image) {
	g.buttonManager.Draw(screen)
}

func (g *Game) DrawRunning(screen *ebiten.Image) {
	g.renderer.Draw(screen)
	DrawScore(screen, g)
}
