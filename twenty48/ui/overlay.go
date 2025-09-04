package ui

import "github.com/hajimehoshi/ebiten/v2"

type Overlay interface{ Draw(screen *ebiten.Image) }

type Manager struct {
	before []Overlay // Drawn before FSM
	after  []Overlay // Drawn after FSM

	beforeIsDisabled bool
	afterIsDisabled  bool
}

func NewOverlayManager() *Manager { return &Manager{} }

func (o *Manager) AddBefore(ov Overlay) {
	o.before = append(o.before, ov)
}

func (o *Manager) AddAfter(ov Overlay) {
	o.after = append(o.after, ov)
}

func (o *Manager) BeforeDraw(screen *ebiten.Image) {
	if o.beforeIsDisabled {
		return
	}
	for _, g := range o.before {
		g.Draw(screen)
	}
}

func (o *Manager) AfterDraw(screen *ebiten.Image) {
	if o.afterIsDisabled {
		return
	}
	for _, g := range o.after {
		g.Draw(screen)
	}
}

func (o *Manager) DisableAfter(b bool) {
	o.afterIsDisabled = b
}

func (o *Manager) DisableBefore(b bool) {
	o.afterIsDisabled = b
}
