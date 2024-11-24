package twenty48

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

type Menu struct {
	game *Game
}

// Initialize menu
func NewMenu(g *Game) *Menu {
	var m *Menu = &Menu{
		game: g,
	}

	return m
}

func (m *Menu) DrawMenu(screen *ebiten.Image) {

	switch m.game.state {
	case StateMainMenu:
		m.DrawMainMenu(screen)
	case StateInstructions:
		m.DrawInstructions(screen)
	default:
		ebitenutil.DebugPrint(screen, "Undefined Menu State")
	}
}

func (m *Menu) DrawMainMenu(screen *ebiten.Image) {
	var realWidth, realHeight int = m.game.GetRealWidthHeight()

	// Title
	m.DrawDoubleText(screen, "2048", realWidth/2, realHeight/2, 2, mplusBigFont, true)

	// Instruction key info
	m.DrawDoubleText(screen, "I: Instructions", realWidth/2, (realHeight/2)+realHeight/10, 1, mplusNormalFontMini, true)

}

func (m *Menu) DrawInstructions(screen *ebiten.Image) {
	var realWidth, realHeight int = m.game.GetRealWidthHeight()

	// Title
	m.DrawDoubleText(screen, "Instructions", realWidth/2, realHeight/10, 2, mplusBigFont, true)

	// Instructions messages
	instructions := []string{
		"Use Arrow Keys or WASD to move tiles",
		"Combine tiles with the same number",
		"Reach 2048 to win the game!",
		"Press R to restart",
		fmt.Sprintf("Press F to toggle Fullscreen: %v", m.game.screenControl.fullscreen),
		fmt.Sprintf("Press Q to toggle dark mode: %v", m.game.darkMode),
	}

	// Render each instruction line
	for i, line := range instructions {
		// Adjust Y-position dynamically based on line index
		lineYPos := (realHeight / 5) + i*(realHeight/18)
		m.DrawDoubleText(screen, line, realWidth/2, lineYPos, 1, mplusNormalFontMini, true)
	}

	// Add a back button
	m.DrawDoubleText(screen, "Press I to return to Main Menu", realWidth/2, realHeight-realHeight/10, 1, mplusNormalFontMini, true)
}

func (m *Menu) DrawDoubleText(screen *ebiten.Image, message string, xpos int, ypos int, offset int, fontUsed font.Face, isCentered bool) {

	var textPosX int = xpos * int(m.game.scale)
	var textPosY int = ypos*int(m.game.scale) + text.BoundString(fontUsed, message).Dy()

	if isCentered {
		textPosX = xpos*int(m.game.scale) - text.BoundString(fontUsed, message).Dx()/2
		textPosY = ypos*int(m.game.scale) + text.BoundString(fontUsed, message).Dy()/2
	}

	text.Draw(screen, message, fontUsed,
		textPosX,
		textPosY,
		color.Black)
	text.Draw(screen, message, fontUsed,
		textPosX-(offset)*int(m.game.scale),
		textPosY-(offset)*int(m.game.scale),
		color.White)
}
