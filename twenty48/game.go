package twenty48

import (
	"image/color"

	"github.com/andersjosef/2048/twenty48/board"
	"github.com/andersjosef/2048/twenty48/buttons"
	"github.com/andersjosef/2048/twenty48/commands"
	co "github.com/andersjosef/2048/twenty48/constants"
	"github.com/andersjosef/2048/twenty48/core"
	"github.com/andersjosef/2048/twenty48/eventhandler"
	"github.com/andersjosef/2048/twenty48/input"
	"github.com/andersjosef/2048/twenty48/menu"
	"github.com/andersjosef/2048/twenty48/shadertools"
	"github.com/andersjosef/2048/twenty48/theme"
	"github.com/andersjosef/2048/twenty48/ui"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	d Deps

	Board          *board.Board
	screenControl  ScreenControl
	animation      Animation
	Menu           *menu.Menu
	Renderer       Renderer
	Input          *input.Input
	Buttons        *buttons.ButtonManager
	utils          Utils
	EventBus       *eventhandler.EventBus
	Cmds           *commands.Commands
	OverlayManager *ui.Manager
	Core           *core.Core
	Theme          *theme.ThemeManager
	ScoreOverlay   *ui.ScoreOverlay
}

func NewGame(d Deps) (*Game, error) {
	// init game struct
	g := &Game{
		d: d,
	}

	g.EventBus = eventhandler.NewEventBus()
	g.screenControl = NewScreenControl(g)
	g.Theme = theme.NewThemeService(theme.ThemeManagerDeps{
		SC: g.screenControl,
	})
	g.Core = core.NewCore()
	g.Board = NewBoard(g)
	g.animation = NewAnimation(g)
	g.Renderer = NewRenderer(g)
	g.utils = NewUtils()
	g.ScoreOverlay = ui.NewScoreOverlay(ui.ScoreOverlayDeps{
		Fonts: g.Theme,
		Score: g.Core,
	})

	g.Cmds = NewCommands(g)
	g.Input = NewInput(g, g.Cmds)
	g.Buttons = NewButtonManager(g, g.Cmds)
	g.Input.GiveButtons(g.Buttons)
	g.Buttons.GiveInput(g.Input)

	g.Menu = NewMenu(g)
	ebiten.SetWindowSize(
		co.LOGICAL_WIDTH*int(g.screenControl.GetScale()),
		co.LOGICAL_HEIGHT*int(g.screenControl.GetScale()),
	)

	g.OverlayManager = ui.NewOverlayManager()
	g.OverlayManager.AddBefore(ui.Background{Color: func() color.RGBA { return g.Theme.Current().ColorScreenBackground }}) // Temporary
	g.OverlayManager.AddAfter(g.Buttons)
	g.OverlayManager.AddAfter(g.Menu)

	g.registerEvents()
	return g, nil
}

func (g *Game) registerEvents() {
	g.EventBus.Register(
		eventhandler.EventResetGame,
		func(eventhandler.Event) {
			g.Core.Reset()
			g.SetState(co.StateMainMenu) // Swap to main menu
			shadertools.ResetTimesMapsDissolve()
		},
	)
}
