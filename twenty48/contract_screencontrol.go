package twenty48

import "github.com/andersjosef/2048/twenty48/screencontrol"

type ScreenControl interface {
	GetActualSize() (x, y int)
	ToggleFullScreen()
	IsFullScreen() bool
	IncrementScale()
	DecrementScale() bool
	GetScale() float64
	LayoutF(float64, float64) (float64, float64)
}

func NewScreenControl(g *Game) ScreenControl {
	d := screencontrol.Deps{
		EventHandler: g.EventBus,
	}

	return screencontrol.New(d)
}
