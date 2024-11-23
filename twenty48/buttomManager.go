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
}

func InitButtonManager(g *Game) *ButtonManager {
	var bm *ButtonManager = &ButtonManager{
		game:           g,
		buttonArrayMap: make(map[GameState][]*Button),
	}

	return bm
}

func (bm *ButtonManager) drawButtons(screen *ebiten.Image) {

	for _, button := range bm.buttonArrayMap[bm.game.state] {
		startX, startY := button.startPos[0], button.startPos[1]
		height, width := startX-button.endPos[0], startY-button.endPos[1]

		// Button background
		vector.DrawFilledRect(screen,
			float32(startX), float32(startY),
			float32(width), float32(height),
			color.Black, false)

		bm.game.menu.DrawDoubleText(screen,
			button.text,
			startX, startY, 2, button.font)

	}
}

func (bm *ButtonManager) checkButtons() {
	// On mouse click loop over every button in array
	// If cursor is within range of some button do the buttons action

	// Can left, right or wheel click
	var pressed bool = ebiten.IsMouseButtonPressed(ebiten.MouseButton0) ||
		ebiten.IsMouseButtonPressed(ebiten.MouseButton1) ||
		ebiten.IsMouseButtonPressed(ebiten.MouseButton2)

	if !pressed {
		return
	}

	curX, curY := ebiten.CursorPosition()

	buttonArray := bm.buttonArrayMap[bm.game.state]
	for _, button := range buttonArray {
		if button.cursorWithin(curX, curY) {
			button.onTrigger()
		}
	}
}

func (bm *ButtonManager) AddButton(text string, startPos [2]int, font font.Face, actionFunction ActionFunc, state GameState) {
	// Create new button obj
	newButton := &Button{
		game:           bm.game,
		startPos:       startPos,
		text:           text,
		font:           font,
		actionFunction: actionFunction,
	}

	// tmp
	newButton.endPos = [2]int{
		newButton.startPos[0] + 90,
		newButton.startPos[0] + 90}

	// Append to list
	bm.buttonArrayMap[state] = append(bm.buttonArrayMap[state], newButton)

}

// // Button ////
type Button struct {
	game           *Game
	startPos       [2]int
	endPos         [2]int
	text           string
	font           font.Face
	actionFunction ActionFunc
	// TODO: Need limits and init method bound to button manager
}

func (bu *Button) cursorWithin(curX, curY int) bool {
	if curX >= bu.startPos[0] && curX <= bu.endPos[0] {
		if curY >= bu.startPos[1] && curY <= bu.endPos[1] {
			return true
		}
	}
	return false
}

func (bu *Button) onTrigger() {
	bu.actionFunction(bu.game.input)
}
