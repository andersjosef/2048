package twenty48

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type FontType int

const (
	FontMini FontType = iota
	FontSmaller
	FontNormal
	FontBig
)

// // Button ////
type Button struct {
	game           *Game
	startPos       [2]int
	endPos         [2]int
	identifier     string
	text           string
	font           *text.GoTextFace
	fontType       FontType
	offset         float64
	actionFunction ActionFunc
}

// Use when wanting to move a button
func (bu *Button) UpdatePos(posX, posY int) {
	dx, dy, err := bu.GetDimentions()
	if err != nil {
		log.Fatal(err)
	}

	var textLengt int = (dx / 2)
	var textHeight int = (dy / 2)

	bu.startPos = [2]int{
		posX - textLengt,
		posY - textHeight,
	}
	bu.endPos = [2]int{
		posX + textLengt,
		posY + textHeight,
	}

}

// For making the buttons text dynamic, should be called before update pos
func (bu *Button) UpdateText(newText string) {
	bu.text = newText

}

func (bu *Button) CursorWithin(curX, curY int) bool {
	curX = int(float64(curX))
	curY = int(float64(curY))

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

	textLength, textHeight := text.Measure(bu.text, bu.font, 0)

	return int(textLength), int(textHeight), nil
}

func (bu *Button) OnTrigger() {
	bu.actionFunction(bu.game.input)
}
