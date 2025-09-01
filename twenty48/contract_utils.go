package twenty48

import (
	"github.com/andersjosef/2048/twenty48/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Utils interface {
	DrawDoubleText(screen *ebiten.Image, message string, xpos int, ypos int, offset float64, fontUsed *text.GoTextFace, isCentered bool)
}

func NewUtils() Utils {
	return utils.Utils{}
}
