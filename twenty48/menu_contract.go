package twenty48

import (
	"github.com/andersjosef/2048/twenty48/menu"
	"github.com/hajimehoshi/ebiten/v2"
)

type Menu interface {
	Draw(screen *ebiten.Image)
	UpdateDynamicText()
	UpdateCenteredTitle()
}

func NewMenu(g *Game) Menu {
	d := &menu.Deps{
		Renderer:     g,
		Buttons:      g,
		EventHandler: g,
		GetSnapShot: func() menu.Snapshot {
			w, h := g.GetActualSize()
			return menu.Snapshot{
				State:         g.GetState(),
				PreviousState: g.GetPreviousState(),
				CurrentTheme:  g.GetCurrentTheme(),
				Fonts:         g.GetFontSet(),
				Score:         g.GetScore(),
				Widht:         w,
				Height:        h,
				IsFullScreen:  g.IsFullScreen(),
			}
		},
	}
	return menu.New(d)
}
