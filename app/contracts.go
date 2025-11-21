package app

import "github.com/hajimehoshi/ebiten/v2"

type Overlay interface {
	BeforeDraw(*ebiten.Image)
	AfterDraw(*ebiten.Image)
	DisableAfter(b bool)
	DisableBefore(b bool)
}
type ScreenControl interface {
	GetActualSize() (x, y int)
	ToggleFullScreen()
	IsFullScreen() bool
	IncrementScale()
	DecrementScale() bool
	GetScale() float64
	LayoutF(float64, float64) (float64, float64)
}
