package input

import (
	"math"
)

const MOVE_THRESHOLD = 100 // Delta distance needed to trigger a move

type Input struct {
	d        Deps
	cursor   *CursorVisibility
	keyboard *KeyboardInput
	mouse    *MouseInput

	touchInput *TouchInput
}

func New(d Deps) *Input {
	var i = &Input{
		d:      d,
		cursor: NewCursorVisibility(20),
	}

	i.keyboard = NewKeyboardInput(KeyboardDeps{
		State:  d.State,
		cmds:   d.Cmds,
		Cursor: i.cursor,
	})
	i.mouse = NewMouseInput(MouseInputDeps{
		State:  d.State,
		Cmds:   d.Cmds,
		Cursor: i.cursor,
	})

	i.touchInput = newTouchInput(i)

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
	i.touchInput.Update()
	return nil
}

func (i *Input) SelectMoveDelta(dx, dy int) {
	if i.d.State.IsGameOver() {
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
