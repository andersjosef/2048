package twenty48

import (
	co "github.com/andersjosef/2048/twenty48/constants"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// // Button Manager ////
type ButtonManager struct {
	game           *Game
	buttonArrayMap map[co.GameState][]*Button
	buttonKeyMap   map[string]*Button
	buttonPressed  bool
}

func InitButtonManager(g *Game) *ButtonManager {
	var bm *ButtonManager = &ButtonManager{
		game:           g,
		buttonArrayMap: make(map[co.GameState][]*Button),
		buttonKeyMap:   make(map[string]*Button),
	}

	// Initialize all buttons
	bm.initButtons()

	return bm
}

// Initiaze buttons here with Addbutton
func (bm *ButtonManager) initButtons() {

	smallOffsett := 1.0
	width, _ := bm.game.GetActualSize()

	// Main Menu
	bm.AddButton(
		"I: Instructions",
		[2]int{0, 0},
		smallOffsett,
		bm.game.fontSet.Mini,
		FontMini,
		toggleInfo,
		co.StateMainMenu,
	)

	// Instructions
	bm.AddButton(
		"Press R to restart",
		[2]int{0, 0},
		smallOffsett,
		bm.game.fontSet.Mini,
		FontMini,
		ResetGame,
		co.StateInstructions,
	)

	bm.AddButton(
		"Press F to toggle Fullscreen",
		[2]int{0, 0},
		smallOffsett,
		bm.game.fontSet.Mini,
		FontMini,
		ToggleFullScreen,
		co.StateInstructions,
	)

	bm.AddButton(
		"Press Q to toggle theme:",
		[2]int{0, 0},
		smallOffsett,
		bm.game.fontSet.Mini,
		FontMini,
		toggleTheme,
		co.StateInstructions,
	)

	bm.AddButton(
		"Press Escape to quit",
		[2]int{0, 0},
		smallOffsett,
		bm.game.fontSet.Mini,
		FontMini,
		CloseGame,
		co.StateInstructions,
	)

	bm.AddButton(
		"Press I to return",
		[2]int{0, 0},
		smallOffsett,
		bm.game.fontSet.Mini,
		FontMini,
		toggleInfo,
		co.StateInstructions,
	)

	// Running loop
	bm.AddButton(
		"II",
		[2]int{width - 20, 20},
		smallOffsett,
		bm.game.fontSet.Mini,
		-1, // Something not in enum for now, needs update to size
		toggleInfo,
		co.StateRunning,
	)

	// Game Over

	bm.AddButton(
		"R: Play again",
		[2]int{0, 0},
		smallOffsett,
		bm.game.fontSet.Mini,
		FontMini,
		ResetGame,
		co.StateGameOver,
	)

}

func (bm *ButtonManager) drawButtons(screen *ebiten.Image) {

	for _, button := range bm.buttonArrayMap[bm.game.state] {
		button.Draw(screen)
	}
}

func (bm *ButtonManager) checkButtons() bool {
	// On mouse click loop over every button in array
	// If cursor is within range of some button do the buttons action
	tapped := bm.game.input.touchInput.checkTapped()

	// Can left, right or wheel click
	var pressed bool = ebiten.IsMouseButtonPressed(ebiten.MouseButton0) ||
		ebiten.IsMouseButtonPressed(ebiten.MouseButton1) ||
		ebiten.IsMouseButtonPressed(ebiten.MouseButton2) || tapped

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
		var tapWithin bool
		for _, tap := range bm.game.input.touchInput.taps {
			if button.CursorWithin(tap.X, tap.Y) {
				tapWithin = true
				bm.game.input.touchInput.taps = bm.game.input.touchInput.taps[:0]
				break
			}
		}
		if button.CursorWithin(curX, curY) || tapWithin {
			button.OnTrigger()
			bm.buttonPressed = true
			return true
		}
	}
	return false
}

func (bm *ButtonManager) AddButton(buttonText string, startPos [2]int, offset float64, font *text.GoTextFace, fontType FontType, actionFunction ActionFunc, state co.GameState) {
	// Create new button obj
	newButton := &Button{
		game:           bm.game,
		identifier:     buttonText,
		text:           buttonText,
		font:           font,
		fontType:       fontType,
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

func (bm *ButtonManager) UpdateFontsForButtons() {
	for _, buttons := range bm.buttonArrayMap {
		for _, button := range buttons {
			// Check fonts and update them
			switch button.fontType {
			case FontMini:
				button.font = bm.game.fontSet.Mini
			case FontSmaller:
				button.font = bm.game.fontSet.Smaller
			case FontNormal:
				button.font = bm.game.fontSet.Normal
			case FontBig:
				button.font = bm.game.fontSet.Big
			}
		}
	}
}
