package app

import (
	"log"

	"github.com/andersjosef/2048/twenty48"
	co "github.com/andersjosef/2048/twenty48/constants"
	"github.com/andersjosef/2048/twenty48/state"
	"github.com/hajimehoshi/ebiten/v2"
)

type App struct {
	fsm     *state.FSM[co.GameState]
	sc      twenty48.ScreenControl
	globals []Updater
	overlay []Drawer
}

func NewApp() *App {
	f := state.NewFSM[co.GameState]()

	g, err := twenty48.NewGame(twenty48.Deps{
		FMS: f,
		IsGameOver: func() bool {
			return f.Current() == co.StateGameOver
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	menu := &state.MainMenu{
		D: state.DepsMainMenu{
			G: g,
		},
	}

	run := &state.Running{
		D: g,
	}

	// Let keybindings change FSM
	g.Input.SetNavigator(func(gs co.GameState) {
		f.Switch(gs)
	})

	f.Register(co.StateMainMenu, menu)
	f.Register(co.StateRunning, run)
	f.Start(co.StateMainMenu)

	return &App{
		fsm: f,
		sc:  g.ScreenControl(),
		globals: []Updater{
			g,
		},
		// overlay: []Drawer{
		// 	g,
		// },
	}
}

func (a *App) Update() error {
	// Global updaters which will run regardless
	for _, g := range a.globals {
		if err := g.Update(); err != nil {
			return err
		}
	}
	// FSM update
	return a.fsm.Update()
}
func (a *App) Draw(screen *ebiten.Image) {
	// Global overlay drawing
	for _, g := range a.overlay {
		g.Draw(screen)
	}

	// FSM drawing
	a.fsm.Draw(screen)
}
func (a *App) Layout(_, _ int) (int, int) { panic("use Ebitengine >=v2.5.0") }
func (a *App) LayoutF(w, h float64) (float64, float64) {
	return a.sc.LayoutF(w, h)
}
