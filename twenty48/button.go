package twenty48

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

// // Button ////
type Button struct {
	game           *Game
	startPos       [2]int
	endPos         [2]int
	identifier     string
	text           string
	font           font.Face
	offset         int
	actionFunction ActionFunc
}

// Use when wanting to move a button
func (bu *Button) UpdatePos(posX, posY int) {
	dx, dy, err := bu.GetDimentions()
	if err != nil {
		log.Fatal(err)
	}

	scale := bu.game.scale

	// Scale delta down since they use actual size
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

}

// For making the buttons text dynamic, should be called before update pos
func (bu *Button) UpdateText(newText string) {
	bu.text = newText

}

func (bu *Button) CursorWithin(curX, curY int) bool {
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

func (bu *Button) GetDimentions() (int, int, error) {
	if bu.font == nil {
		return -1, -1, fmt.Errorf("cant get dimentions, font is not set")
	}
	var x int = text.BoundString(bu.font, bu.text).Dx()
	var y int = text.BoundString(bu.font, bu.text).Dy()

	return x, y, nil
}

func (bu *Button) OnTrigger() {
	bu.actionFunction(bu.game.input)
}
