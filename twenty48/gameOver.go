package twenty48

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func (g *Game) DrawGameOverScreen(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Game over screen")
}
