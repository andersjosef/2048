package app

import (
	"github.com/andersjosef/2048/twenty48"
	co "github.com/andersjosef/2048/twenty48/constants"
	"github.com/andersjosef/2048/twenty48/state"
	"github.com/hajimehoshi/ebiten/v2"
)

type App struct {
	fsm *state.FSM[co.GameState]
	sc  twenty48.ScreenControl
}

func NewApp(g *twenty48.Game) *App {
	f := state.NewFSM[co.GameState]()

	menu := &state.MainMenu{
		D: g,
	}

	run := &state.Running{
		D: g,
	}

	f.Register(co.StateMainMenu, menu)
	f.Register(co.StateRunning, run)
	f.Start(co.StateMainMenu)

	return &App{fsm: f, sc: g.ScreenControl()}
}

func (a *App) Update() error              { return a.fsm.Update() }
func (a *App) Draw(screen *ebiten.Image)  { a.fsm.Draw(screen) }
func (a *App) Layout(_, _ int) (int, int) { panic("use Ebitengine >=v2.5.0") }
func (a *App) LayoutF(w, h float64) (float64, float64) {
	return a.sc.LayoutF(w, h)
}
