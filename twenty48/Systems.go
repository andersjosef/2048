package twenty48

import (
	"image/color"

	"github.com/andersjosef/2048/twenty48/animations"
	"github.com/andersjosef/2048/twenty48/board"
	"github.com/andersjosef/2048/twenty48/board_view.go"
	"github.com/andersjosef/2048/twenty48/buttons"
	"github.com/andersjosef/2048/twenty48/commands"
	"github.com/andersjosef/2048/twenty48/core"
	"github.com/andersjosef/2048/twenty48/eventhandler"
	"github.com/andersjosef/2048/twenty48/input"
	"github.com/andersjosef/2048/twenty48/menu"
	"github.com/andersjosef/2048/twenty48/theme"
	"github.com/andersjosef/2048/twenty48/ui"
	"github.com/andersjosef/2048/twenty48/ui/layout"
)

type Systems struct {
	d Deps

	Board          *board.Board
	BoardView      *board_view.BoardView
	screenControl  ScreenControl
	animation      *animations.Animation
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
	Layout         *layout.Layout
}

func Build(d Deps) (*Systems, error) {
	g := &Systems{
		d: d,
	}

	g.EventBus = eventhandler.NewEventBus()
	g.screenControl = NewScreenControl(g)
	g.Theme = theme.NewThemeService(theme.ThemeManagerDeps{
		SC: g.screenControl,
	})
	g.Core = core.NewCore()
	g.Board = NewBoard(g)
	g.Layout = layout.New(layout.SizesDeps{
		EventHandler:  g.EventBus,
		ScreenControl: g.screenControl,
	})
	g.BoardView = board_view.NewBoardView(board_view.BoardViewDeps{
		EventHandler:  g.EventBus,
		ScreenControl: g.screenControl,
		Board:         g.Board,
		Theme:         g.Theme,
		Layout:        g.Layout,
		IsGameOver:    d.IsGameOver,
	})
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
	g.Input.GiveButtons(g.Buttons) // TODO: fix this
	g.Buttons.GiveInput(g.Input)

	g.Menu = NewMenu(g)

	g.OverlayManager = ui.NewOverlayManager()
	g.OverlayManager.AddBefore(ui.Background{Color: func() color.RGBA { return g.Theme.Current().ColorScreenBackground }})
	g.OverlayManager.AddAfter(g.Buttons)
	g.OverlayManager.AddAfter(g.Menu)

	return g, nil
}
