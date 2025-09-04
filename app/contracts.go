package app

import "github.com/hajimehoshi/ebiten/v2"

type Overlay interface {
	BeforeDraw(*ebiten.Image)
	AfterDraw(*ebiten.Image)
	DisableAfter(b bool)
	DisableBefore(b bool)
}
