package twenty48

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
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

	// Testbutton
	// bm.AddButton(
	// 	"Test Button!",
	// 	[2]int{200, 200},
	// 	mplusNormalFontMini,
	// 	testForButtonAction,
	// 	StateMainMenu,
	// )

	// Main Menu
	bm.AddButton(
		"I: Instructions",
		[2]int{0, 0},
		mplusNormalFontMini,
		toggleInfo,
		StateMainMenu,
	)

	// Instructions
	bm.AddButton(
		"Press R to restart",
		[2]int{0, 0},
		mplusNormalFontMini,
		ResetGame,
		StateInstructions,
	)

	bm.AddButton(
		fmt.Sprintf("Press F to toggle Fullscreen: %v", bm.game.screenControl.fullscreen),
		[2]int{0, 0},
		mplusNormalFontMini,
		ToggleFullScreen,
		StateInstructions,
	)

	bm.AddButton(
		fmt.Sprintf("Press Q to toggle dark mode: %v", bm.game.darkMode),
		[2]int{0, 0},
		mplusNormalFontMini,
		SwitchDefaultDarkMode,
		StateInstructions,
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
			startX, startY, 2, button.font, false)

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
		if button.cursorWithin(curX, curY) {
			button.onTrigger()
			bm.buttonPressed = true
			return true
		}
	}
	return false
}

func (bm *ButtonManager) AddButton(buttonText string, startPos [2]int, font font.Face, actionFunction ActionFunc, state GameState) {
	// Create new button obj
	newButton := &Button{
		game:           bm.game,
		text:           buttonText,
		font:           font,
		actionFunction: actionFunction,
	}

	// Set position of button
	newButton.UpdatePos(startPos[0], startPos[1])

	// Append to list
	bm.buttonArrayMap[state] = append(bm.buttonArrayMap[state], newButton)

	// Store text as key for access other places in the code
	bm.buttonKeyMap[newButton.text] = newButton
}

// // Button ////
type Button struct {
	game           *Game
	startPos       [2]int
	endPos         [2]int
	text           string
	font           font.Face
	actionFunction ActionFunc
}

func (bu *Button) UpdatePos(posX, posY int) {
	dx, dy, err := bu.getDimentions()
	if err != nil {
		log.Fatal(err)
	}

	scale := bu.game.scale

	var textLengt int = (dx / 2) / int(scale)
	var textWidth int = (dy / 2) / int(scale)

	bu.startPos = [2]int{
		posX - textLengt,
		posY - textWidth,
	}
	bu.endPos = [2]int{
		posX + textLengt,
		posY + textWidth,
	}

	fmt.Printf("Button bounds - startPos: %v, endPos: %v\n", bu.startPos, bu.endPos)
}

func (bu *Button) cursorWithin(curX, curY int) bool {
	scale := ebiten.Monitor().DeviceScaleFactor()
	curX = int(float64(curX) / scale)
	curY = int(float64(curY) / scale)

	if curX >= bu.startPos[0] && curX <= bu.endPos[0] {
		if curY >= bu.startPos[1] && curY <= bu.endPos[1] {
			return true
		}
	}
	return false
}

func (b *Button) getDimentions() (int, int, error) {
	if b.font == nil {
		return -1, -1, fmt.Errorf("cant get dimentions, font is not set")
	}
	var x int = text.BoundString(b.font, b.text).Dx()
	var y int = text.BoundString(b.font, b.text).Dy()

	return x, y, nil
}

func (bu *Button) onTrigger() {
	bu.actionFunction(bu.game.input)
}
