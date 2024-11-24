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
