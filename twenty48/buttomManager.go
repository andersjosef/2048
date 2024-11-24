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

	if !pressed {
		return false
	}

	curX, curY := ebiten.CursorPosition()

	buttonArray := bm.buttonArrayMap[bm.game.state]
	for _, button := range buttonArray {
		if button.cursorWithin(curX, curY) {
			button.onTrigger()
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

	// tmp
	dx, dy, err := newButton.getDimentions()
	if err != nil {
		log.Fatal(err)
	}
	var textLengt int = dx / 2
	var textWidth int = dy / 2

	newButton.startPos = [2]int{
		startPos[0] - textLengt,
		startPos[1] - textWidth,
	}
	newButton.endPos = [2]int{
		startPos[0] + textLengt,
		startPos[1] + textWidth,
	}

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
