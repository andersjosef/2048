// Copyright 2022 The Ebitengine Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package twenty48

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Handle touch input

type TouchInput struct {
	input *Input

	releasedTouchIDs []ebiten.TouchID
	pressedTouchIDs  []ebiten.TouchID

	touchMap map[ebiten.TouchID]Touches
}

func newTouchInput(i *Input) *TouchInput {
	t := &TouchInput{
		input:            i,
		releasedTouchIDs: make([]ebiten.TouchID, 5),
		pressedTouchIDs:  make([]ebiten.TouchID, 5),
		touchMap:         make(map[ebiten.TouchID]Touches),
	}
	return t
}

type Touches struct {
	startPos *pos
	endPos   *pos
}

type pos struct {
	x int
	y int
}

// Creates a new touches with start position
func newTouches(startX, startY int) Touches {
	touches := Touches{}
	touches.startPos = &pos{}
	touches.startPos.x = startX
	touches.startPos.y = startY
	return touches
}

func (t *TouchInput) endTouchesAndEvaluate(id ebiten.TouchID, endX, endY int) {
	// find touches by id and set endPos
	touches := t.touchMap[id]
	touches.endPos = &pos{}
	touches.endPos.x = endX
	touches.endPos.y = endY

	shouldTrigger, dx, dy := touches.shoudTriggerTouchMove()
	if shouldTrigger {
		t.input.SelectMoveDelta(dx, dy)
		delete(t.touchMap, id)
	}

}

// Updates and evaluates
func (t *TouchInput) handleTouchInput() {

	t.pressedTouchIDs = inpututil.AppendJustPressedTouchIDs(t.pressedTouchIDs[:0])
	t.releasedTouchIDs = inpututil.AppendJustReleasedTouchIDs(t.releasedTouchIDs[:0])

	for _, id := range t.pressedTouchIDs {
		if t.input.game.state == StateMainMenu { // If in main menu click will trigger game state
			t.input.game.state = StateRunning
		}
		_, exist := t.touchMap[id]
		if !exist {
			x, y := ebiten.TouchPosition(id)
			t.touchMap[id] = newTouches(x, y)
		}
	}

	for _, id := range t.releasedTouchIDs {
		_, exist := t.touchMap[id]
		if exist {
			x, y := inpututil.TouchPositionInPreviousTick(id)
			t.endTouchesAndEvaluate(id, x, y)
			// Delete entry regardless because its released
			delete(t.touchMap, id)
		}
	}
	// Check if reached max length without release
	for _, id := range t.pressedTouchIDs {
		_, exist := t.touchMap[id]
		if exist {
			x, y := ebiten.TouchPosition(id)
			t.endTouchesAndEvaluate(id, x, y)
			// Delete entry regardless because its released
		}
	}

}

// Calculates distance of move and tells if it should trigger a move
func (to *Touches) shoudTriggerTouchMove() (bool, int, int) {
	if to.endPos == nil || to.startPos == nil {
		return false, 0, 0
	}

	dx := to.endPos.x - to.startPos.x
	dy := to.endPos.y - to.startPos.y

	return int(math.Abs(float64(dx))) > MOVE_THRESHOLD || int(math.Abs(float64(dy))) > MOVE_THRESHOLD,
		dx,
		dy
}
