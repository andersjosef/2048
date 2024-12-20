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
	m.titleImage = m.initTitle()
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
	var realWidth, realHeight int = m.game.screenControl.GetRealWidthHeight()

	// Title
	m.drawTitle(screen)

	// Instruction key info
	insX := realWidth / 2
	insY := (realHeight / 2) + realHeight/10
	m.game.buttonManager.buttonKeyMap["I: Instructions"].UpdatePos(insX, insY)

}

func (m *Menu) DrawInstructions(screen *ebiten.Image) {
	var realWidth, realHeight int = m.game.screenControl.GetRealWidthHeight()

	// Title
	m.game.renderer.DrawDoubleText(screen, "Instructions", realWidth/2, realHeight/10, 2, m.game.fontSet.Big, true)

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
		rowXPos := realWidth / 2
		lineYPos := (realHeight / 5) + i*(realHeight/18)

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
	m.game.buttonManager.buttonKeyMap["Press I to return"].UpdatePos(realWidth/2, realHeight-realHeight/10)
}

func (m *Menu) UpdateDynamicText() {
	m.dynamicText["Press F to toggle Fullscreen"] = fmt.Sprintf("Press F to toggle Fullscreen: %v", m.game.screenControl.fullscreen)
	m.dynamicText["Press Q to toggle dark mode"] = fmt.Sprintf("Press Q to toggle dark mode: %v", m.game.darkMode)
}

func (m *Menu) initTitle() *ebiten.Image {
	var realWidth, realHeight int = m.game.screenControl.GetRealWidthHeight()
	newImage := ebiten.NewImage(realWidth*int(m.game.scale), realHeight*int(m.game.scale))
	m.game.renderer.DrawDoubleText(newImage, "2048", realWidth/2, realHeight/2, 2, m.game.fontSet.Big, true)

	return newImage

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
