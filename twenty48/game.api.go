package twenty48

import (
	co "github.com/andersjosef/2048/twenty48/constants"
	"github.com/andersjosef/2048/twenty48/eventhandler"
	"github.com/andersjosef/2048/twenty48/theme"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

/*
	Functions that are needed for interfaces
*/

// // GameProvider ////
func (g *Game) GetScore() int {
	return g.score
}

func (g *Game) GetState() co.GameState {
	return g.state
}

func (g *Game) GetPreviousState() co.GameState {
	return g.previousState
}

func (g *Game) GetCurrentTheme() theme.Theme {
	return g.currentTheme
}

func (g *Game) GetFontSet() theme.FontSet {
	return *g.fontSet
}

func (g *Game) GetScale() float64 {
	return g.screenControl.GetScale()
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
	return g.screenControl.GetActualSize()
}

func (g *Game) IsFullScreen() bool {
	return g.screenControl.IsFullScreen()
}

// // RendererProvider ////
func (g *Game) DrawDoubleText(screen *ebiten.Image, message string, xpos int, ypos int, offset float64, fontUsed *text.GoTextFace, isCentered bool) {
	g.renderer.DrawDoubleText(screen, message, xpos, ypos, offset, fontUsed, isCentered)
}

// // bus ////
func (g *Game) Register(eventType eventhandler.EventType, handler func(eventhandler.Event)) {
	g.eventBus.Register(eventType, handler)
}

func (g *Game) Emit(event eventhandler.Event) {
	g.eventBus.Emit(event)
}

func (g *Game) Dispatch() {
	g.eventBus.Dispatch()
}
