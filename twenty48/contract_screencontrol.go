package twenty48

import "github.com/andersjosef/2048/twenty48/screencontrol"

type ScreenControl interface {
	GetActualSize() (x, y int)
	ToggleFullScreen()
	IsFullScreen() bool
	IncrementScale()
	DecrementScale() bool
	GetScale() float64
}

func NewScreenControl(g *Game) ScreenControl {
	d := screencontrol.Deps{
		EventHandler: g.eventBus,
	}

	return screencontrol.New(d)
}
