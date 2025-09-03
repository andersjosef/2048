package twenty48

import (
	"github.com/andersjosef/2048/twenty48/menu"
)

// type Menu interface {
// 	Draw(screen *ebiten.Image)
// 	UpdateDynamicText()
// 	UpdateCenteredTitle()
// }

func NewMenu(g *Game) *menu.Menu {
	d := menu.Deps{
		Renderer:     g.utils,
		Buttons:      g.buttonManager,
		EventHandler: g.EventBus,
		GetSnapshot: func() menu.Snapshot {
			w, h := g.screenControl.GetActualSize()
			return menu.Snapshot{
				State:         g.GetState(),
				PreviousState: g.GetPreviousState(),
				CurrentTheme:  g.GetCurrentTheme(),
				Fonts:         g.GetFontSet(),
				Score:         g.GetScore(),
				Width:         w,
				Height:        h,
				IsFullScreen:  g.screenControl.IsFullScreen(),
			}
		},
	}
	return menu.New(d)
}
