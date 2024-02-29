package twenty48

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

func (g *Game) DrawMenu(screen *ebiten.Image) {
	DrawDoubleText(screen, "2048", SCREENWIDTH/2, SCREENHEIGHT/3, 2)
	DrawDoubleText(screen, "Press any button", SCREENWIDTH/2, int(float32(SCREENHEIGHT)/1.5), 2)
}

func DrawDoubleText(screen *ebiten.Image, message string, xpos int, ypos int, offsett int) {
	myFont := mplusNormalFontSmaller
	text.Draw(screen, message, myFont,
		xpos-text.BoundString(myFont, message).Dx()/2,
		ypos+text.BoundString(myFont, message).Dy()/2,
		color.Black)
	text.Draw(screen, message, myFont,
		xpos-offsett-text.BoundString(myFont, message).Dx()/2,
		ypos-offsett+text.BoundString(myFont, message).Dy()/2,
		color.White)
}
