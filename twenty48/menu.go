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
	game        *Game
	dynamicText map[string]string
}

// Initialize menu
func NewMenu(g *Game) *Menu {
	var m *Menu = &Menu{
		game: g,
	}

	m.dynamicText = map[string]string{
		"Press F to toggle Fullscreen": fmt.Sprintf("Press F to toggle Fullscreen: %v", m.game.screenControl.fullscreen),
		"Press Q to toggle dark mode":  fmt.Sprintf("Press Q to toggle dark mode: %v", m.game.darkMode),
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
	m.DrawDoubleText(screen, "2048", realWidth/2, realHeight/2, 2, m.game.fontSet.Big, true)

	// Instruction key info
	insX := realWidth / 2
	insY := (realHeight / 2) + realHeight/10
	m.game.buttonManager.buttonKeyMap["I: Instructions"].UpdatePos(insX, insY)
	// m.DrawDoubleText(screen, "I: Instructions", insX, insY, 1, mplusNormalFontMini, true)

}

func (m *Menu) DrawInstructions(screen *ebiten.Image) {
	var realWidth, realHeight int = m.game.GetRealWidthHeight()

	// Title
	m.DrawDoubleText(screen, "Instructions", realWidth/2, realHeight/10, 2, m.game.fontSet.Big, true)

	// Instructions messages
	instructions := []string{
		"Use Arrow Keys or WASD to move tiles",
		"Combine tiles with the same number",
		"Reach 2048 to win the game!",
		"Press R to restart",
		// fmt.Sprintf("Press F to toggle Fullscreen: %v", m.game.screenControl.fullscreen),
		// fmt.Sprintf("Press Q to toggle dark mode: %v", m.game.darkMode),
		"Press F to toggle Fullscreen",
		"Press Q to toggle dark mode",
	}

	// Render each instruction line
	for i, line := range instructions {
		// Adjust Y-position dynamically based on line index
		rowXPos := realWidth / 2
		lineYPos := (realHeight / 5) + i*(realHeight/18)

		if button, ok := m.game.buttonManager.buttonKeyMap[line]; ok {
			if newText, ok := m.dynamicText[button.identifier]; ok {
				button.UpdateText(newText)
			}
			button.UpdatePos(rowXPos, lineYPos)
		} else {
			m.DrawDoubleText(screen, line, rowXPos, lineYPos, 1, m.game.fontSet.Mini, true)
		}

	}

	// Add a back button
	returnButtonText := "Press I to return"
	if m.game.previousState == StateMainMenu {
		returnButtonText += " to Main Menu"
	} else if m.game.previousState == StateRunning {
		returnButtonText += " to Game"
	}
	m.game.buttonManager.buttonKeyMap["Press I to return"].UpdateText(returnButtonText)
	m.game.buttonManager.buttonKeyMap["Press I to return"].UpdatePos(realWidth/2, realHeight-realHeight/10)
}

func (m *Menu) DrawDoubleText(screen *ebiten.Image, message string, xpos int, ypos int, offset int, fontUsed font.Face, isCentered bool) {

	scale := m.game.scale

	// Calculate text dimensions
	textWidth := text.BoundString(fontUsed, message).Dx()
	textHeight := text.BoundString(fontUsed, message).Dy()

	// Scale the position
	textPosX := int(scale) * xpos
	textPosY := int(scale) * ypos

	// Adjust for centering
	if isCentered {
		textPosX -= textWidth / 2  // Center horizontally
		textPosY += textHeight / 2 // Center vertically
	} else {
		textPosY += textHeight // This works best :S
	}

	// Draw shadow (black text)
	text.Draw(screen, message, fontUsed,
		textPosX,
		textPosY,
		color.Black)

	// Draw main text (white text) with offset
	text.Draw(screen, message, fontUsed,
		textPosX-int(float64(offset)*scale),
		textPosY-int(float64(offset)*scale),
		color.White)
}

func (m *Menu) UpdateDynamicText() {
	m.dynamicText["Press F to toggle Fullscreen"] = fmt.Sprintf("Press F to toggle Fullscreen: %v", m.game.screenControl.fullscreen)
	m.dynamicText["Press Q to toggle dark mode"] = fmt.Sprintf("Press Q to toggle dark mode: %v", m.game.darkMode)
}
