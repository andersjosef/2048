package state

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Running struct {
	Renderer interface {
		Draw(*ebiten.Image)
	}
	ScoreUI interface {
		DrawScore(*ebiten.Image)
	}
}

func (s *Running) Enter() {}

func (s *Running) Exit() {}

func (s *Running) Update() error { return nil }

func (s *Running) Draw(screen *ebiten.Image) {
	s.Renderer.Draw(screen)
	s.ScoreUI.DrawScore(screen)
}
