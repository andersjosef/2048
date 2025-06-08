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

func (g *Game) GetPreviousState() GameState {
	return g.previousState
}

func (g *Game) GetCurrentTheme() theme.Theme {
	return g.currentTheme
}

func (g *Game) GetFontSet() theme.FontSet {
	return *g.fontSet
}

// // ButtonManagerProvider ////
func (g *Game) ButtonExists(keyName string) (exists bool) {
	_, exists = g.buttonManager.buttonKeyMap[keyName]
	return
}

func (g *Game) UpdatePosForButton(keyName string, posX, posY int) (exists bool) {
	if button, doExist := g.buttonManager.buttonKeyMap[keyName]; doExist {
		button.UpdatePos(posX, posY)
		return true
	}
	return false
}

func (g *Game) UpdateTextForButton(keyName, newText string) (exists bool) {
	if button, doExist := g.buttonManager.buttonKeyMap[keyName]; doExist {
		button.UpdateText(newText)
		return true
	}
	return false
}

func (g *Game) GetButton(identifier string) (button *Button, exists bool) {
	if button, doExist := g.buttonManager.buttonKeyMap[identifier]; doExist {
		return button, true
	}
	return nil, false
}

// // ScreenControlProvider ////
func (g *Game) GetActualSize() (x, y int) {
	return g.screenControl.actualWidth, g.screenControl.actualHeight
}

func (g *Game) IsFullScreen() bool {
	return g.screenControl.isFullscreen
}

// // RendererProvider ////
func (g *Game) DrawDoubleText(screen *ebiten.Image, message string, xpos int, ypos int, offset float64, fontUsed *text.GoTextFace, isCentered bool) {
	g.renderer.DrawDoubleText(screen, message, xpos, ypos, offset, fontUsed, isCentered)
}
