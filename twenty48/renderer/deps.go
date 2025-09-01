package renderer

import "github.com/hajimehoshi/ebiten/v2"

type Deps struct {
	Animation
	Base
}

type Animation interface {
	Draw(*ebiten.Image)
	IsAnimating() bool
}

type Base interface {
	Draw(*ebiten.Image)
}
