package app

import "github.com/hajimehoshi/ebiten/v2"

type Drawer interface{ Draw(screen *ebiten.Image) }

type Overlay struct {
	before []Drawer // Drawn before FSM
	after  []Drawer // Drawn after FSM

	beforeIsDisabled bool
	afterIsDisabled  bool
}

func (o *Overlay) BeforeDraw(screen *ebiten.Image) {
	if o.beforeIsDisabled {
		return
	}
	for _, g := range o.before {
		g.Draw(screen)
	}
}

func (o *Overlay) AfterDraw(screen *ebiten.Image) {
	if o.afterIsDisabled {
		return
	}
	for _, g := range o.after {
		g.Draw(screen)
	}
}

func (o *Overlay) DisableAfter(b bool) {
	o.afterIsDisabled = b
}

func (o *Overlay) DisableBefore(b bool) {
	o.afterIsDisabled = b
}
