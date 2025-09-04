package state

import (
	"github.com/andersjosef/2048/twenty48/theme"
	"github.com/hajimehoshi/ebiten/v2"
)

type DepsMainMenu struct {
	Menu interface {
		Draw(*ebiten.Image)
	}
	G interface {
		DrawUI(*ebiten.Image)
		GetCurrentTheme() theme.Theme
	}
}

type MainMenu struct {
	D DepsMainMenu
}

func (s *MainMenu) Enter() {}

func (s *MainMenu) Exit() {}

func (s *MainMenu) Update() error {
	return nil
}

func (s *MainMenu) Draw(screen *ebiten.Image) {
	s.D.Menu.Draw(screen)
}
