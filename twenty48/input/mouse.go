package input

import (
	"math"

	"github.com/andersjosef/2048/twenty48/commands"
	co "github.com/andersjosef/2048/twenty48/constants"
	"github.com/hajimehoshi/ebiten/v2"
)

type MouseInputDeps struct {
	State interface {
		GetState() co.GameState
		SetState(co.GameState)
		IsGameOver() bool
	}
	Cmds   *commands.Commands
	Cursor interface{ MaybeShow() }
}

type MouseInput struct {
	d MouseInputDeps

	// Cursor positions
	startCursorPos [2]int
	endCursorPos   [2]int
	justMoved      bool // To make sure only one move is done

	onPressed map[co.GameState]func()
}

func NewMouseInput(d MouseInputDeps) *MouseInput {
	mi := &MouseInput{
		d: d,
	}

	mi.onPressed = map[co.GameState]func(){
		co.StateMainMenu: func() { mi.d.Cmds.GoToRunning() },
	}

	return mi
}

func (i *MouseInput) Update() {
	i.d.Cursor.MaybeShow()

	// Can left, right or wheel click
	var pressed bool = ebiten.IsMouseButtonPressed(ebiten.MouseButton0) ||
		ebiten.IsMouseButtonPressed(ebiten.MouseButton1) ||
		ebiten.IsMouseButtonPressed(ebiten.MouseButton2)

	// Cursor movement updates
	if pressed {
		if onPressAction, exist := i.onPressed[i.d.State.GetState()]; exist { // If specified will trigger state change
			onPressAction()
		} else { // If not in menu update only end cursor coordinate
			i.endCursorPos[0], i.endCursorPos[1] = ebiten.CursorPosition()
		}
	} else { // If not clicking: update both values
		i.resetMouseState()
	}

	// Check if delta movements is large enough to trigger move
	if i.shoulTriggerMove() && !i.justMoved {
		i.performMove()
		i.justMoved = true
	}
}

func (m *MouseInput) shoulTriggerMove() bool {
	dx := m.endCursorPos[0] - m.startCursorPos[0]
	dy := m.endCursorPos[1] - m.startCursorPos[1]

	return int(math.Abs(float64(dx))) > MOVE_THRESHOLD || int(math.Abs(float64(dy))) > MOVE_THRESHOLD
}

func (m *MouseInput) resetMouseState() {
	m.justMoved = false
	m.startCursorPos[0], m.startCursorPos[1] = ebiten.CursorPosition()
	m.endCursorPos[0], m.endCursorPos[1] = ebiten.CursorPosition()
}

func (i *MouseInput) performMove() {
	if i.d.State.IsGameOver() {
		return
	}
	dx := i.endCursorPos[0] - i.startCursorPos[0]
	dy := i.endCursorPos[1] - i.startCursorPos[1]

	i.SelectMoveDelta(dx, dy)
}

func (i *MouseInput) SelectMoveDelta(dx, dy int) {
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
