package twenty48

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

// Controlling which part of the menu all within the gamestate MainMenu
type MenuState int

const (
	MenuStateMain MenuState = iota + 1
	MenuStateInstructions
)

type Menu struct {
	game  *Game
	state MenuState
}

// Initialize menu
func NewMenu(g *Game) *Menu {
	var m *Menu = &Menu{
		game:  g,
		state: MenuStateMain,
	}

	return m
}

func (m *Menu) DrawMenu(screen *ebiten.Image) {

	switch m.state {
	case MenuStateMain:
		m.DrawMainMenu(screen)
	case MenuStateInstructions:
		ebitenutil.DebugPrint(screen, "Instructions state")
	default:
		ebitenutil.DebugPrint(screen, "Undefined Menu State")
	}
}

func (m *Menu) DrawMainMenu(screen *ebiten.Image) {
	var realWidth, realHeight int = m.game.GetRealWidthHeight()
	m.DrawDoubleText(screen, "2048", realWidth/2, realHeight/2, 2, mplusNormalFontSmaller)

}

func (m *Menu) DrawDoubleText(screen *ebiten.Image, message string, xpos int, ypos int, offset int, fontUsed font.Face) {

	var textPosX int = xpos*int(m.game.scale) - text.BoundString(fontUsed, message).Dx()/2
	var textPosY int = ypos*int(m.game.scale) + text.BoundString(fontUsed, message).Dy()/2

	text.Draw(screen, message, fontUsed,
		textPosX,
		textPosY,
		color.Black)
	text.Draw(screen, message, fontUsed,
		textPosX-(offset)*int(m.game.scale),
		textPosY-(offset)*int(m.game.scale),
		color.White)
}
