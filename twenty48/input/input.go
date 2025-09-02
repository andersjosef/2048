package input

import (
	"math"

	co "github.com/andersjosef/2048/twenty48/constants"
	"github.com/hajimehoshi/ebiten/v2"
)

const MOVE_THRESHOLD = 100 // Delta distance needed to trigger a move

type Input struct {
	d        Deps
	keyboard *KeyboardInput
	mouse    *MouseInput

	touchInput *TouchInput

	isHidingMouse      bool
	lastPosMouse       [2]int
	showMouseThreshold float64 // If mouse is moved beyond this show again
}

func New(d Deps) *Input {
	var i = &Input{
		d: d,
	}

	i.keyboard = NewKeyboardInput(KeyboardDeps{
		GetState:                   func() co.GameState { return d.GetState() },
		SetState:                   func(gs co.GameState) { d.SetState(gs) },
		CheckForMakingCursorHidden: func() { i.checkForMakingCursorHidden() },
		cmds:                       d.Cmds,
	})
	i.mouse = NewMouseInput(MouseInputDeps{
		GetState:                    func() co.GameState { return d.GetState() },
		SetState:                    func(gs co.GameState) { d.SetState(gs) },
		CheckForMakingCursorVisible: func() { i.checkForMakingCursorVisible() },
		IsGameOver:                  func() bool { return d.IsGameOver() },
		Cmds:                        d.Cmds,
	})

	i.touchInput = newTouchInput(i)
	i.showMouseThreshold = 20 // Set how much the mouse has to move to reappear

	return i
}

func (i *Input) GiveButtons(b Buttons) {
	i.d.Buttons = b
}

func (i *Input) UpdateInput() error {
	// Keyboard and Mouse input handling
	if i.d.Buttons.CheckButtons() {
		return nil
	}
	i.keyboard.Update()
	i.mouse.Update()
	i.touchInput.TouchUpdate()
	return nil
}

func (i *Input) SelectMoveDelta(dx, dy int) {
	if i.d.IsGameOver() {
		return
	}
	if math.Abs(float64(dx)) > math.Abs(float64(dy)) { // X-axis largest
		if dx > 0 {
			i.d.Cmds.MoveRight()
		} else {
			i.d.Cmds.MoveLeft()
		}
	} else { // Y-axis largest
		if dy > 0 {
			i.d.Cmds.MoveDown()
		} else {
			i.d.Cmds.MoveUp()
		}
	}

}

///// Utilities //////

// Helper functions for toggeling mouse being displayed or not
func (i *Input) checkForMakingCursorVisible() {
	if i.isHidingMouse {
		lastX := float64(i.lastPosMouse[0])
		lastY := float64(i.lastPosMouse[1])

		x, y := ebiten.CursorPosition()

		if math.Abs(lastX-float64(x)) > i.showMouseThreshold ||
			math.Abs(lastY-float64(y)) > i.showMouseThreshold {
			ebiten.SetCursorMode(ebiten.CursorModeVisible)
			i.isHidingMouse = false
		}
	}
}

func (i *Input) checkForMakingCursorHidden() {
	if !i.isHidingMouse {
		i.lastPosMouse[0], i.lastPosMouse[1] = ebiten.CursorPosition()
		ebiten.SetCursorMode(ebiten.CursorModeHidden)
		i.isHidingMouse = true
	}
}
