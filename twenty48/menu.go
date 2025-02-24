package twenty48

import (
	"fmt"

	"github.com/andersjosef/2048/twenty48/shadertools"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Menu struct {
	game        *Game
	dynamicText map[string]string

	titleImage      *ebiten.Image
	titleInFullView bool
}

// Initialize menu
func NewMenu(g *Game) *Menu {
	var m *Menu = &Menu{
		game: g,
	}
	m.initTitle() // Inits title image to menu parameter
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

	// Title
	m.drawTitle(screen)

	// Instruction key info
	insX := m.game.screenControl.actualWidth / 2
	insY := (m.game.screenControl.actualHeight / 2) + m.game.screenControl.actualHeight/10
	m.game.buttonManager.buttonKeyMap["I: Instructions"].UpdatePos(insX, insY)

}

func (m *Menu) DrawInstructions(screen *ebiten.Image) {

	// Title
	m.game.renderer.DrawDoubleText(screen, "Instructions", m.game.screenControl.actualWidth/2, m.game.screenControl.actualHeight/10, 2, m.game.fontSet.Big, true)

	// Instructions messages
	instructions := []string{
		"Use Arrow Keys, WASD or the mouse to move tiles",
		"Combine tiles with the same number",
		"Reach 2048 to win the game!",
		"Press Escape to quit",
		"Press R to restart",
		"Press F to toggle Fullscreen",
		"Press Q to toggle dark mode",
	}

	// Render each instruction line
	for i, line := range instructions {
		// Adjust Y-position dynamically based on line index
		rowXPos := m.game.screenControl.actualWidth / 2
		lineYPos := (m.game.screenControl.actualHeight / 5) + i*(m.game.screenControl.actualHeight/18)

		if button, ok := m.game.buttonManager.buttonKeyMap[line]; ok {
			if newText, ok := m.dynamicText[button.identifier]; ok {
				button.UpdateText(newText)
			}
			button.UpdatePos(rowXPos, lineYPos)
		} else {
			m.game.renderer.DrawDoubleText(screen, line, rowXPos, lineYPos, 1, m.game.fontSet.Mini, true)
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
	m.game.buttonManager.buttonKeyMap["Press I to return"].UpdatePos(m.game.screenControl.actualWidth/2, m.game.screenControl.actualHeight-m.game.screenControl.actualHeight/10)
}

func (m *Menu) UpdateDynamicText() {
	m.dynamicText["Press F to toggle Fullscreen"] = fmt.Sprintf("Press F to toggle Fullscreen: %v", m.game.screenControl.fullscreen)
	m.dynamicText["Press Q to toggle dark mode"] = fmt.Sprintf("Press Q to toggle dark mode: %v", m.game.darkMode)
}

func (m *Menu) initTitle() {
	var xPos, yPos = m.game.screenControl.actualWidth, m.game.screenControl.actualHeight

	newImage := ebiten.NewImage(xPos, yPos)
	m.game.renderer.DrawDoubleText(newImage, "2048", xPos/2, yPos/2, 2, m.game.fontSet.Big, true)
	m.titleImage = newImage

}

// Drawing the title and using shader when animating
func (m *Menu) drawTitle(screen *ebiten.Image) {

	if !m.titleInFullView {
		shaderImage, isDone := shadertools.GetImageFadeIn(m.titleImage)
		if isDone {
			m.titleInFullView = true
		}
		screen.DrawImage(shaderImage, &ebiten.DrawImageOptions{})

	} else {
		screen.DrawImage(m.titleImage, &ebiten.DrawImageOptions{})

	}

}
