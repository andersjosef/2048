package twenty48

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font"
)

// // Button Manager ////
type ButtonManager struct {
	game           *Game
	buttonArrayMap map[GameState][]*Button
	buttonKeyMap   map[string]*Button
	buttonPressed  bool
}

func InitButtonManager(g *Game) *ButtonManager {
	var bm *ButtonManager = &ButtonManager{
		game:           g,
		buttonArrayMap: make(map[GameState][]*Button),
		buttonKeyMap:   make(map[string]*Button),
	}

	// Initialize all buttons
	bm.initButtons()

	return bm
}

// Initiaze buttons here with Addbutton
func (bm *ButtonManager) initButtons() {

	smallOffsett := 1

	// Main Menu
	bm.AddButton(
		"I: Instructions",
		[2]int{0, 0},
		smallOffsett,
		bm.game.fontSet.Mini,
		toggleInfo,
		StateMainMenu,
	)

	// Instructions
	bm.AddButton(
		"Press R to restart",
		[2]int{0, 0},
		smallOffsett,
		bm.game.fontSet.Mini,
		ResetGame,
		StateInstructions,
	)

	bm.AddButton(
		"Press F to toggle Fullscreen",
		[2]int{0, 0},
		smallOffsett,
		bm.game.fontSet.Mini,
		ToggleFullScreen,
		StateInstructions,
	)

	bm.AddButton(
		"Press Q to toggle dark mode",
		[2]int{0, 0},
		smallOffsett,
		bm.game.fontSet.Mini,
		SwitchDefaultDarkMode,
		StateInstructions,
	)

	bm.AddButton(
		"Press I to return",
		[2]int{0, 0},
		smallOffsett,
		bm.game.fontSet.Mini,
		toggleInfo,
		StateInstructions,
	)

	// Running loop
	bm.AddButton(
		"II",
		[2]int{SCREENWIDTH - 20, 20},
		smallOffsett,
		bm.game.fontSet.Mini,
		toggleInfo,
		StateRunning,
	)

}

func (bm *ButtonManager) drawButtons(screen *ebiten.Image) {

	for _, button := range bm.buttonArrayMap[bm.game.state] {
		startX, startY := button.startPos[0], button.startPos[1]
		width, height := button.endPos[0]-startX, button.endPos[1]-startY

		// Button background
		vector.DrawFilledRect(screen,
			float32(startX), float32(startY),
			float32(width), float32(height),
			color.RGBA{0, 0, 0, 0}, false)

		bm.game.menu.DrawDoubleText(screen,
			button.text,
			startX, startY, button.offset, button.font, false)

	}
}

func (bm *ButtonManager) checkButtons() bool {
	// On mouse click loop over every button in array
	// If cursor is within range of some button do the buttons action

	// Can left, right or wheel click
	var pressed bool = ebiten.IsMouseButtonPressed(ebiten.MouseButton0) ||
		ebiten.IsMouseButtonPressed(ebiten.MouseButton1) ||
		ebiten.IsMouseButtonPressed(ebiten.MouseButton2)

	// Dont check if button isnt pressed
	if !pressed {
		bm.buttonPressed = false
		return false
	}

	// Check to make sure action is only triggered once
	if bm.buttonPressed {
		return true
	}

	curX, curY := ebiten.CursorPosition()

	buttonArray := bm.buttonArrayMap[bm.game.state]
	for _, button := range buttonArray {
		if button.CursorWithin(curX, curY) {
			button.OnTrigger()
			bm.buttonPressed = true
			return true
		}
	}
	return false
}

func (bm *ButtonManager) AddButton(buttonText string, startPos [2]int, offset int, font font.Face, actionFunction ActionFunc, state GameState) {
	// Create new button obj
	newButton := &Button{
		game:           bm.game,
		identifier:     buttonText,
		text:           buttonText,
		font:           font,
		actionFunction: actionFunction,
		offset:         offset,
	}

	// Set position of button
	newButton.UpdatePos(startPos[0], startPos[1])

	// Append to list
	bm.buttonArrayMap[state] = append(bm.buttonArrayMap[state], newButton)

	// Store text as key for access other places in the code
	bm.buttonKeyMap[newButton.identifier] = newButton
}
