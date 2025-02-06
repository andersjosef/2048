package twenty48

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) DrawGameOverScreen(screen *ebiten.Image) {
	scoreString := fmt.Sprintf("Score: %v", g.score)

	g.renderer.DrawDoubleText(screen, "Game Over", g.screenControl.actualWidth/2, g.screenControl.actualHeight/3, 2, g.fontSet.Big, true)
	g.renderer.DrawDoubleText(screen, scoreString, g.screenControl.actualWidth/2, g.screenControl.actualHeight/2, 2, g.fontSet.Mini, true)
	// Restart Button pos
	g.buttonManager.buttonKeyMap["R: Play again"].UpdatePos(g.screenControl.actualWidth/2, g.screenControl.actualHeight-g.screenControl.actualHeight/3)

}
