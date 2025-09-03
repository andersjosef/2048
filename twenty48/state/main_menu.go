package state

import (
	"github.com/andersjosef/2048/twenty48/theme"
	"github.com/hajimehoshi/ebiten/v2"
)

type MainMenu struct {
	D interface {
		DrawMenu(*ebiten.Image)
		DrawUI(*ebiten.Image)
		GetCurrentTheme() theme.Theme
	}
}

func (s *MainMenu) Enter() {}

func (s *MainMenu) Exit() {}

func (s *MainMenu) Update() error {
	return nil
}

func (s *MainMenu) Draw(screen *ebiten.Image) {
	screen.Fill(s.D.GetCurrentTheme().ColorScreenBackground)
	s.D.DrawMenu(screen)
	s.D.DrawUI(screen)
}
