package menu

import (
	"fmt"

	co "github.com/andersjosef/2048/twenty48/constants"
	"github.com/andersjosef/2048/twenty48/eventhandler"
	"github.com/andersjosef/2048/twenty48/shadertools"
	"github.com/hajimehoshi/ebiten/v2"
)

type Menu struct {
	d           Deps
	dynamicText map[string]string

	titleImage      *ebiten.Image
	titleInFullView bool

	lastSnap Snapshot
}

// Initialize menu
func New(d Deps) *Menu {
	var m *Menu = &Menu{
		d: d,
	}
	snap := m.d.GetSnapshot()
	m.lastSnap = snap

	m.UpdateCenteredTitle() // Inits title image to menu parameter
	m.dynamicText = map[string]string{
		"Press F to toggle Fullscreen": fmt.Sprintf("Press F to toggle Fullscreen: %v", snap.IsFullScreen),
		"Press Q to toggle theme:":     fmt.Sprintf("Press Q to toggle theme: %v", snap.CurrentTheme.Name),
	}

	m.registerEvents()

	return m
}

func (m *Menu) registerEvents() {
	// Register events
	m.d.Register(
		eventhandler.EventScreenChanged,
		func(evt eventhandler.Event) {
			m.UpdateCenteredTitle()
			m.UpdateDynamicText()
		},
	)
	m.d.Register(
		eventhandler.EventResetGame,
		func(_ eventhandler.Event) {
			m.titleInFullView = false
		},
	)

	m.d.Register(
		eventhandler.EventThemeChanged,
		func(_ eventhandler.Event) {
			m.UpdateDynamicText()
		},
	)
}

func (m *Menu) Draw(screen *ebiten.Image) {
	m.lastSnap = m.d.GetSnapshot()
	switch m.lastSnap.State {
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
	width, height := m.lastSnap.Width, m.lastSnap.Height

	insX := width / 2
	insY := (height / 2) + height/10

	m.d.UpdatePosForButton("I: Instructions", insX, insY)

}

func (m *Menu) drawInstructions(screen *ebiten.Image) {

	width, height := m.lastSnap.Width, m.lastSnap.Height

	// Title
	m.d.DrawDoubleText(
		screen,
		"Instructions",
		width/2,
		height/10,
		2,
		m.lastSnap.Fonts.Big,
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

		if m.d.ButtonExists(line) { // Buttons
			if newText, doExist := m.dynamicText[line]; doExist {
				m.d.UpdateTextForButton(line, newText)
			}
			m.d.UpdatePosForButton(line, rowXPos, lineYPos)
		} else { // Just text
			m.d.DrawDoubleText(
				screen,
				line,
				rowXPos,
				lineYPos,
				1,
				m.lastSnap.Fonts.Mini,
				true,
			)
		}

	}

	// Add a back button
	returnButtonText := "Press I to return"
	switch m.lastSnap.PreviousState {
	case co.StateMainMenu:
		returnButtonText += " to Main Menu"
	case co.StateRunning:
		returnButtonText += " to Game"
	}

	m.d.UpdateTextForButton(
		"Press I to return",
		returnButtonText,
	)
	m.d.UpdatePosForButton(
		"Press I to return",
		width/2,
		height-height/10,
	)
}

func (m *Menu) UpdateDynamicText() {
	snap := m.d.GetSnapshot()
	m.dynamicText["Press F to toggle Fullscreen"] = fmt.Sprintf("Press F to toggle Fullscreen: %v", snap.IsFullScreen)
	m.dynamicText["Press Q to toggle theme:"] = fmt.Sprintf("Press Q to toggle theme: %v", snap.CurrentTheme.Name)
}

// Resenter title on change
func (m *Menu) UpdateCenteredTitle() {
	m.lastSnap = m.d.GetSnapshot()
	xPos, yPos := m.lastSnap.Width, m.lastSnap.Height

	newImage := ebiten.NewImage(xPos, yPos)
	m.d.DrawDoubleText(
		newImage,
		"2048",
		xPos/2,
		yPos/2,
		2,
		m.lastSnap.Fonts.Big,
		true,
	)
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

func (m *Menu) drawGameOverScreen(screen *ebiten.Image) {
	scoreString := fmt.Sprintf("Score: %v", m.lastSnap.Score)

	width, height := m.lastSnap.Width, m.lastSnap.Height

	m.d.DrawDoubleText(
		screen,
		"Game Over",
		width/2,
		height/3,
		2,
		m.lastSnap.Fonts.Big,
		true,
	)
	m.d.DrawDoubleText(
		screen,
		scoreString,
		width/2,
		height-height/2,
		2,
		m.lastSnap.Fonts.Mini,
		true,
	)
	// Restart Button pos
	m.d.UpdatePosForButton(
		"R: Play again",
		width/2,
		height-height/3,
	)

}
