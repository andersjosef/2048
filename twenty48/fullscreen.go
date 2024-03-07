package twenty48

import "github.com/hajimehoshi/ebiten/v2"

type ScreenControl struct {
	fullscreen bool
	game       *Game
}

func InitScreenControl(g *Game) *ScreenControl {
	sc := &ScreenControl{
		fullscreen: false,
		game:       g,
	}

	return sc
}

func (b *Board) ToggleFullScreen() {
	if b.game.screenControl.fullscreen {
		ebiten.SetFullscreen(false)
		b.game.screenControl.fullscreen = false
	} else {
		ebiten.SetFullscreen(true)
		b.game.screenControl.fullscreen = true
	}
}
