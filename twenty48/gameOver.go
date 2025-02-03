package twenty48

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) DrawGameOverScreen(screen *ebiten.Image) {
	scoreString := fmt.Sprintf("Score: %v", g.score)

	g.renderer.DrawDoubleText(screen, "Game Over", logicalWidth/2, logicalHeight/3, 2, g.fontSet.Big, true)
	g.renderer.DrawDoubleText(screen, scoreString, logicalWidth/2, logicalHeight/2, 2, g.fontSet.Mini, true)
	// Restart Button pos
	g.buttonManager.buttonKeyMap["R: Play again"].UpdatePos(logicalWidth/2, logicalHeight-logicalHeight/3)

}
