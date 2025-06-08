package menu

import (
	"fmt"

	co "github.com/andersjosef/2048/twenty48/constants"
	"github.com/andersjosef/2048/twenty48/eventhandler"
	"github.com/andersjosef/2048/twenty48/shadertools"
	"github.com/hajimehoshi/ebiten/v2"
)

type Menu struct {
	view        GameView
	dynamicText map[string]string

	titleImage      *ebiten.Image
	TitleInFullView bool
}

// Initialize menu
func NewMenu(v GameView) *Menu {
	var m *Menu = &Menu{
		view: v,
	}
	m.UpdateCenteredTitle() // Inits title image to menu parameter
	m.dynamicText = map[string]string{
		"Press F to toggle Fullscreen": fmt.Sprintf("Press F to toggle Fullscreen: %v", m.view.IsFullScreen()),
		"Press Q to toggle theme:":     fmt.Sprintf("Press Q to toggle theme: %v", m.view.GetCurrentTheme().Name),
	}

	// Register events
	v.GetBusHandler().Register(
		eventhandler.EventScreenChanged,
		func(evt eventhandler.Event) {
			m.UpdateCenteredTitle()
		},
	)

	return m
}

func (m *Menu) Draw(screen *ebiten.Image) {
	switch m.view.GetState() {
	case co.StateMainMenu:
		m.drawMainMenu(screen)
	case co.StateInstructions:
		m.drawInstructions(screen)
	case co.StateGameOver:
		m.drawGameOverScreen(screen)
	}
}

func (m *Menu) drawMainMenu(screen *ebiten.Image) {

	// Title
	m.drawTitle(screen)

	// Instruction key info
	width, height := m.view.GetActualSize()

	insX := width / 2
	insY := (height / 2) + height/10

	m.view.UpdatePosForButton("I: Instructions", insX, insY)

}

func (m *Menu) drawInstructions(screen *ebiten.Image) {

	width, height := m.view.GetActualSize()

	// Title
	m.view.DrawDoubleText(
		screen,
		"Instructions",
		width/2,
		height/10,
		2,
		m.view.GetFontSet().Big,
		true,
	)

	// Instructions messages
	instructions := []string{
		"Use Arrow Keys, WASD or the mouse to move tiles",
		"Combine tiles with the same number",
		"Reach 2048 to win the game!",
		"Press Escape to quit",
		"Press R to restart",
		"Press F to toggle Fullscreen",
		"Press Q to toggle theme:",
	}

	// Render each instruction line
	for i, line := range instructions {
		// Adjust Y-position dynamically based on line index
		rowXPos := width / 2
		lineYPos := (height / 5) + i*(height/18)

		if m.view.ButtonExists(line) { // Buttons
			if newText, doExist := m.dynamicText[line]; doExist {
				m.view.UpdateTextForButton(line, newText)
			}
			m.view.UpdatePosForButton(line, rowXPos, lineYPos)
		} else { // Just text
			m.view.DrawDoubleText(
				screen,
				line,
				rowXPos,
				lineYPos,
				1,
				m.view.GetFontSet().Mini,
				true,
			)
		}

	}

	// Add a back button
	returnButtonText := "Press I to return"
	if m.view.GetPreviousState() == co.StateMainMenu {
		returnButtonText += " to Main Menu"
	} else if m.view.GetPreviousState() == co.StateRunning {
		returnButtonText += " to Game"
	}

	m.view.UpdateTextForButton(
		"Press I to return",
		returnButtonText,
	)
	m.view.UpdatePosForButton(
		"Press I to return",
		width/2,
		height-height/10,
	)
}

func (m *Menu) UpdateDynamicText() {
	m.dynamicText["Press F to toggle Fullscreen"] = fmt.Sprintf("Press F to toggle Fullscreen: %v", m.view.IsFullScreen())
	m.dynamicText["Press Q to toggle theme:"] = fmt.Sprintf("Press Q to toggle theme: %v", m.view.GetCurrentTheme().Name)
}

// Resenter title on change
func (m *Menu) UpdateCenteredTitle() {
	xPos, yPos := m.view.GetActualSize()

	newImage := ebiten.NewImage(xPos, yPos)
	m.view.DrawDoubleText(
		newImage,
		"2048",
		xPos/2,
		yPos/2,
		2,
		m.view.GetFontSet().Big,
		true,
	)
	m.titleImage = newImage

}

// Drawing the title and using shader when animating
func (m *Menu) drawTitle(screen *ebiten.Image) {

	if !m.TitleInFullView {
		shaderImage, isDone := shadertools.GetImageFadeIn(m.titleImage)
		if isDone {
			m.TitleInFullView = true
		}
		screen.DrawImage(shaderImage, &ebiten.DrawImageOptions{})

	} else {
		screen.DrawImage(m.titleImage, &ebiten.DrawImageOptions{})

	}

}

func (m *Menu) drawGameOverScreen(screen *ebiten.Image) {
	scoreString := fmt.Sprintf("Score: %v", m.view.GetScore())

	width, height := m.view.GetActualSize()

	m.view.DrawDoubleText(
		screen,
		"Game Over",
		width/2,
		height/3,
		2,
		m.view.GetFontSet().Big,
		true,
	)
	m.view.DrawDoubleText(
		screen,
		scoreString,
		width/2,
		height-height/2,
		2,
		m.view.GetFontSet().Mini,
		true,
	)
	// Restart Button pos
	m.view.UpdatePosForButton(
		"R: Play again",
		width/2,
		height-height/3,
	)

}
