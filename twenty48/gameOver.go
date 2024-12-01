package twenty48

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) DrawGameOverScreen(screen *ebiten.Image) {
	scoreString := fmt.Sprintf("Score: %v", g.score)
	var realWidth, realHeight int = g.screenControl.GetRealWidthHeight()

	g.renderer.DrawDoubleText(screen, "Game Over", realWidth/2, realHeight/3, 2, g.fontSet.Big, true)
	g.renderer.DrawDoubleText(screen, scoreString, realWidth/2, realHeight/2, 2, g.fontSet.Mini, true)
	// Restart Button pos
	g.buttonManager.buttonKeyMap["R: Play again"].UpdatePos(realWidth/2, realHeight-realHeight/3)

}
