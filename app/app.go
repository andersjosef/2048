package app

import (
	"log"

	"github.com/andersjosef/2048/twenty48"
	co "github.com/andersjosef/2048/twenty48/constants"
	"github.com/andersjosef/2048/twenty48/shadertools"
	"github.com/andersjosef/2048/twenty48/state"
	"github.com/hajimehoshi/ebiten/v2"
)

type App struct {
	fsm     *state.FSM[co.GameState]
	sc      twenty48.ScreenControl
	globals []Updater
	overlay Overlay
}

func NewApp() *App {
	f := state.NewFSM[co.GameState]()

	g, err := twenty48.NewGame(twenty48.Deps{
		FSM: f,
		IsGameOver: func() bool {
			return f.Current() == co.StateGameOver
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	f.Register(co.StateMainMenu, &state.MainMenu{
		D: state.DepsMainMenu{
			Menu: g.Menu,
		},
	})

	f.Register(co.StateInstructions, &state.Instructions{
		D: state.DepsInstructions{
			Menu:    g.Menu,
			Buttons: g.Buttons,
		},
	})

	f.Register(co.StateGameOver, &state.GameOver{
		D: state.GameOverDeps{
			Menu:    g.Menu,
			Board:   g.Board,
			Overlay: g.OverlayManager,
		},
	})

	f.Register(co.StateRunning, &state.Running{
		Renderer: g.Renderer,
		ScoreUI:  g.ScoreOverlay,
	})

	f.Register(co.StateQuitGame, &state.QuitGame{})

	f.Start(co.StateMainMenu)

	return &App{
		fsm: f,
		sc:  g.ScreenControl(),
		globals: []Updater{
			updaterFunc(func() error { g.EventBus.Dispatch(); return nil }),
			updaterFunc(func() error { return g.Input.UpdateInput() }),
			updaterFunc(func() error { shadertools.Update(); return nil }),
		},
		overlay: g.OverlayManager,
	}
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
