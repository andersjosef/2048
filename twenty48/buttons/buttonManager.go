package buttons

import (
	"github.com/andersjosef/2048/twenty48/commands"
	co "github.com/andersjosef/2048/twenty48/constants"
	"github.com/andersjosef/2048/twenty48/eventhandler"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// // Button Manager ////
type ButtonManager struct {
	d              Deps
	Cmds           commands.Commands
	buttonArrayMap map[co.GameState][]*Button
	buttonKeyMap   map[string]*Button
	buttonPressed  bool
}

func NewButtonManager(d Deps, cmds commands.Commands) *ButtonManager {
	bm := &ButtonManager{
		d:              d,
		Cmds:           cmds,
		buttonArrayMap: make(map[co.GameState][]*Button),
		buttonKeyMap:   make(map[string]*Button),
	}

	// Initialize all buttons
	bm.initButtons()

	bm.d.Register(
		eventhandler.EventScreenChanged,
		func(_ eventhandler.Event) {
			bm.updatePauseButtonLocation()
		},
	)

	return bm
}

func (bm *ButtonManager) GiveInput(i Input) {
	bm.d.Input = i
}

// Initiaze buttons here with Addbutton
func (bm *ButtonManager) initButtons() {

	smallOffsett := 1.0
	width, _ := bm.d.GetActualSize()

	// Main Menu
	bm.AddButton(
		"I: Instructions",
		[2]int{0, 0},
		smallOffsett,
		bm.d.GetFontSet().Mini,
		FontMini,
		bm.Cmds.ToggleInfo,
		co.StateMainMenu,
	)

	// Instructions
	bm.AddButton(
		"Press R to restart",
		[2]int{0, 0},
		smallOffsett,
		bm.d.GetFontSet().Mini,
		FontMini,
		bm.Cmds.ResetGame,
		co.StateInstructions,
	)

	bm.AddButton(
		"Press F to toggle Fullscreen",
		[2]int{0, 0},
		smallOffsett,
		bm.d.GetFontSet().Mini,
		FontMini,
		bm.Cmds.ToggleFullscreen,
		co.StateInstructions,
	)

	bm.AddButton(
		"Press Q to toggle theme:",
		[2]int{0, 0},
		smallOffsett,
		bm.d.GetFontSet().Mini,
		FontMini,
		bm.Cmds.ToggleTheme,
		co.StateInstructions,
	)

	bm.AddButton(
		"Press Escape to quit",
		[2]int{0, 0},
		smallOffsett,
		bm.d.GetFontSet().Mini,
		FontMini,
		bm.Cmds.CloseGame,
		co.StateInstructions,
	)

	bm.AddButton(
		"Press I to return",
		[2]int{0, 0},
		smallOffsett,
		bm.d.GetFontSet().Mini,
		FontMini,
		bm.Cmds.ToggleInfo,
		co.StateInstructions,
	)

	// Running loop
	bm.AddButton(
		"II",
		[2]int{width - 20, 20},
		smallOffsett,
		bm.d.GetFontSet().Mini,
		-1, // Something not in enum for now, needs update to size
		bm.Cmds.ToggleInfo,
		co.StateRunning,
	)

	// Game Over

	bm.AddButton(
		"R: Play again",
		[2]int{0, 0},
		smallOffsett,
		bm.d.GetFontSet().Mini,
		FontMini,
		bm.Cmds.ResetGame,
		co.StateGameOver,
	)

}

func (bm *ButtonManager) Draw(screen *ebiten.Image) {

	for _, button := range bm.buttonArrayMap[bm.d.GetState()] {
		button.Draw(screen)
	}
}

func (bm *ButtonManager) CheckButtons() bool {
	// On mouse click loop over every button in array
	// If cursor is within range of some button do the buttons action
	tapped := bm.d.Input.CheckTapped()

	// Can left, right or wheel click
	pressed := ebiten.IsMouseButtonPressed(ebiten.MouseButton0) ||
		ebiten.IsMouseButtonPressed(ebiten.MouseButton1) ||
		ebiten.IsMouseButtonPressed(ebiten.MouseButton2) ||
		tapped

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

	buttonArray := bm.buttonArrayMap[bm.d.GetState()]
	for _, button := range buttonArray {
		var tapWithin bool
		for _, tap := range bm.d.Input.GetTaps() {
			if button.CursorWithin(tap.X, tap.Y) {
				tapWithin = true
				bm.d.Input.ClearTaps()
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

func (bm *ButtonManager) AddButton(buttonText string, startPos [2]int, offset float64, font *text.GoTextFace, fontType FontType, actionFunction func(), state co.GameState) {
	// Create new button obj
	newButton := &Button{
		d:              Deps{Utils: bm.d.Utils},
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
				button.font = bm.d.GetFontSet().Mini
			case FontSmaller:
				button.font = bm.d.GetFontSet().Smaller
			case FontNormal:
				button.font = bm.d.GetFontSet().Normal
			case FontBig:
				button.font = bm.d.GetFontSet().Big
			}
		}
	}
}

// Providing
func (bm *ButtonManager) ButtonExists(keyName string) (exists bool) {
	_, exists = bm.buttonKeyMap[keyName]
	return
}

func (bm *ButtonManager) UpdatePosForButton(keyName string, posX, posY int) (exists bool) {
	if button, doExist := bm.buttonKeyMap[keyName]; doExist {
		button.UpdatePos(posX, posY)
		return true
	}
	return false
}

func (bm *ButtonManager) UpdateTextForButton(keyName, newText string) (exists bool) {
	if button, doExist := bm.buttonKeyMap[keyName]; doExist {
		button.UpdateText(newText)
		return true
	}
	return false
}

func (bm *ButtonManager) GetButton(identifier string) (button *Button, exists bool) {
	if button, doExist := bm.buttonKeyMap[identifier]; doExist {
		return button, true
	}
	return nil, false
}

// Helper function for updating the pause button location
// When changing screen size
func (i *ButtonManager) updatePauseButtonLocation() {
	width, _ := i.d.ScreenControl.GetActualSize()
	i.UpdatePosForButton("II", width-20, 20)
}
