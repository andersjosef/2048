package state

import (
	"github.com/andersjosef/2048/twenty48/theme"
	"github.com/hajimehoshi/ebiten/v2"
)

type DepsMainMenu struct {
	G interface {
		DrawMenu(*ebiten.Image)
		DrawUI(*ebiten.Image)
		GetCurrentTheme() theme.Theme
	}
	GoRun func() // Change to running mode
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
	s.D.G.DrawMenu(screen)
}
