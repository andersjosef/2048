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

func NewMenu(v menu.View) Menu {
	return menu.New(v)
}
