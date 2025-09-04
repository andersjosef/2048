package state

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type GameOverDeps struct {
	Menu interface {
		DrawGameOver(*ebiten.Image)
	}
	Board   interface{ DrawBoardFadeOut(*ebiten.Image) bool }
	Overlay interface{ DisableAfter(bool) }
}

type GameOver struct {
	D        GameOverDeps
	animDone bool
}

func (s *GameOver) Enter() {
	s.animDone = false
	s.D.Overlay.DisableAfter(true)
}

func (s *GameOver) Exit() {
	s.D.Overlay.DisableAfter(false)
}

func (s *GameOver) Update() error { return nil }

func (s *GameOver) Draw(screen *ebiten.Image) {
	if !s.animDone {
		s.animDone = s.D.Board.DrawBoardFadeOut(screen)
		s.D.Overlay.DisableAfter(!s.animDone)
	} else {
		s.D.Menu.DrawGameOver(screen)
	}
}
