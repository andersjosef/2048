package twenty48

import (
	"github.com/andersjosef/2048/twenty48/theme"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

/*
	Functions that are needed for interfaces
*/

// // GameProvider ////
func (g *Game) GetState() GameState {
	return g.state
}

func (g *Game) GetCurrentTheme() theme.Theme {
	return g.currentTheme
}

func (g *Game) GetFontSet() theme.FontSet {
	return *g.fontSet
}

// // ButtonManagerProvider ////
func (g *Game) UpdatePosForButton(keyName string, posX, posY int) {
	g.buttonManager.buttonKeyMap[keyName].UpdatePos(posX, posY)
}

// // ScreenControlProvider ////
func (g *Game) GetActualSize() (x, y int) {
	return g.screenControl.actualWidth, g.screenControl.actualHeight
}

func (g *Game) GetIsFullScreen() bool {
	return g.screenControl.isFullscreen
}

// // RendererProvider ////
func (g *Game) DrawDoubleText(screen *ebiten.Image, message string, xpos int, ypos int, offset float64, fontUsed *text.GoTextFace, isCentered bool) {
	g.renderer.DrawDoubleText(screen, message, xpos, ypos, offset, fontUsed, isCentered)
}
