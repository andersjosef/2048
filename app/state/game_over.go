package state

import (
	"github.com/andersjosef/2048/twenty48/eventhandler"
	"github.com/hajimehoshi/ebiten/v2"
)

type EventHandler interface {
	Register(eventType eventhandler.EventType, handler func(eventhandler.Event))
}

type GameOverDeps struct {
	Menu interface {
		DrawGameOver(*ebiten.Image)
	}
	BoardView interface{ DrawBoardFadeOut(*ebiten.Image) bool }
	Overlay   interface{ DisableAfter(bool) }
	Renderer  interface {
		Draw(*ebiten.Image)
	}
	EventHandler
}

type GameOver struct {
	D          GameOverDeps
	animDone   bool
	shaderDone bool
}

func (s *GameOver) Enter() {
	s.animDone = false
	s.shaderDone = false
	s.D.Overlay.DisableAfter(true)

	s.D.EventHandler.Register(eventhandler.EventAnimationDoneGameOver,
		func(_ eventhandler.Event) {
			s.animDone = true
		})
}

func (s *GameOver) Exit() {
	s.D.Overlay.DisableAfter(false)
}

func (s *GameOver) Update() error { return nil }

func (s *GameOver) Draw(screen *ebiten.Image) {
	if !s.animDone {

		s.D.Renderer.Draw(screen)
		// s.D.Overlay.DisableAfter(!s.animDone)
	} else if !s.shaderDone {
		s.shaderDone = s.D.BoardView.DrawBoardFadeOut(screen)
	} else {
		s.D.Overlay.DisableAfter(false)
		s.D.Menu.DrawGameOver(screen)
	}
}
