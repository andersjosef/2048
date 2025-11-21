package twenty48

import "github.com/andersjosef/2048/twenty48/screencontrol"

func NewScreenControl(g *Systems) *screencontrol.ScreenControl {
	d := screencontrol.Deps{
		EventHandler: g.EventBus,
	}

	return screencontrol.New(d)
}
