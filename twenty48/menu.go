package twenty48

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

func (g *Game) DrawMenu(screen *ebiten.Image) {

	var realWidth, realHeight int = g.GetRealWidthHeight()

	g.DrawDoubleText(screen, "2048", realWidth/2, realHeight/2, 2, mplusNormalFontSmaller)
}

func (g *Game) DrawDoubleText(screen *ebiten.Image, message string, xpos int, ypos int, offset int, fontUsed font.Face) {

	var textPosX int = xpos*int(g.scale) - text.BoundString(fontUsed, message).Dx()/2
	var textPosY int = ypos*int(g.scale) + text.BoundString(fontUsed, message).Dy()/2

	text.Draw(screen, message, fontUsed,
		textPosX,
		textPosY,
		color.Black)
	text.Draw(screen, message, fontUsed,
		textPosX-(offset)*int(g.scale),
		textPosY-(offset)*int(g.scale),
		color.White)
}
