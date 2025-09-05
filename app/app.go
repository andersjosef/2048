package app

import (
	"github.com/andersjosef/2048/app/state"
	"github.com/andersjosef/2048/twenty48"
	co "github.com/andersjosef/2048/twenty48/constants"
	"github.com/andersjosef/2048/twenty48/eventhandler"
	"github.com/andersjosef/2048/twenty48/shadertools"
	"github.com/hajimehoshi/ebiten/v2"
)

type App struct {
	fsm     *state.FSM[co.GameState]
	sc      twenty48.ScreenControl
	globals []Updater
	overlay Overlay
}

func NewApp() (*App, error) {
	f := state.NewFSM[co.GameState]()

	sys, err := twenty48.Build(twenty48.Deps{
		FSM: f,
		IsGameOver: func() bool {
			return f.Current() == co.StateGameOver
		},
	})
	if err != nil {
		return nil, err
	}

	// Events
	sys.EventBus.Register(
		eventhandler.EventResetGame,
		func(eventhandler.Event) {
			sys.Core.Reset()
			sys.SetState(co.StateMainMenu) // Swap to main menu
			shadertools.ResetTimesMapsDissolve()
		},
	)

	// States
	f.Register(co.StateMainMenu, &state.MainMenu{
		D: state.DepsMainMenu{
			Menu: sys.Menu,
		},
	})

	f.Register(co.StateInstructions, &state.Instructions{
		D: state.DepsInstructions{
			Menu:    sys.Menu,
			Buttons: sys.Buttons,
		},
	})

	f.Register(co.StateGameOver, &state.GameOver{
		D: state.GameOverDeps{
			Menu:    sys.Menu,
			Board:   sys.Board,
			Overlay: sys.OverlayManager,
		},
	})

	f.Register(co.StateRunning, &state.Running{
		Renderer: sys.Renderer,
		ScoreUI:  sys.ScoreOverlay,
	})

	f.Register(co.StateQuitGame, &state.QuitGame{})

	f.Start(co.StateMainMenu)

	// Window
	ebiten.SetWindowSize(
		co.LOGICAL_WIDTH*int(sys.ScreenControl().GetScale()),
		co.LOGICAL_HEIGHT*int(sys.ScreenControl().GetScale()),
	)

	return &App{
		fsm: f,
		sc:  sys.ScreenControl(),
		globals: []Updater{
			updaterFunc(func() error { sys.EventBus.Dispatch(); return nil }),
			updaterFunc(func() error { return sys.Input.UpdateInput() }),
			updaterFunc(func() error { shadertools.Update(); return nil }),
		},
		overlay: sys.OverlayManager,
	}, nil
}

func (a *App) Update() error {
	// Global updaters which will run regardless before fsm
	for _, g := range a.globals {
		if err := g.Update(); err != nil {
			return err
		}
	}
	// FSM update
	return a.fsm.Update()
}
func (a *App) Draw(screen *ebiten.Image) {
	// Global overlay before FSM
	a.overlay.BeforeDraw(screen)

	// FSM drawing
	a.fsm.Draw(screen)

	// Global overlay after FSM
	a.overlay.AfterDraw(screen)
}
func (a *App) Layout(_, _ int) (int, int) { panic("use Ebitengine >=v2.5.0") }
func (a *App) LayoutF(w, h float64) (float64, float64) {
	return a.sc.LayoutF(w, h)
}
