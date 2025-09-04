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

	r, err := twenty48.NewRouter(twenty48.Deps{
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
			Menu: r.Menu,
		},
	})

	f.Register(co.StateInstructions, &state.Instructions{
		D: state.DepsInstructions{
			Menu:    r.Menu,
			Buttons: r.Buttons,
		},
	})

	f.Register(co.StateGameOver, &state.GameOver{
		D: state.GameOverDeps{
			Menu:    r.Menu,
			Board:   r.Board,
			Overlay: r.OverlayManager,
		},
	})

	f.Register(co.StateRunning, &state.Running{
		Renderer: r.Renderer,
		ScoreUI:  r.ScoreOverlay,
	})

	f.Register(co.StateQuitGame, &state.QuitGame{})

	f.Start(co.StateMainMenu)

	return &App{
		fsm: f,
		sc:  r.ScreenControl(),
		globals: []Updater{
			updaterFunc(func() error { r.EventBus.Dispatch(); return nil }),
			updaterFunc(func() error { return r.Input.UpdateInput() }),
			updaterFunc(func() error { shadertools.Update(); return nil }),
		},
		overlay: r.OverlayManager,
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
