// Copyright 2021 The Ebiten Authors
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
	_ "image/jpeg"
	"math"

	co "github.com/andersjosef/2048/twenty48/constants"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// distance between points a and b.
func distance(xa, ya, xb, yb int) float64 {
	x := math.Abs(float64(xa - xb))
	y := math.Abs(float64(ya - yb))
	return math.Sqrt(x*x + y*y)
}

type touch struct {
	originX, originY int
	currX, currY     int
	duration         int
	wasPinch, isPan  bool
}

// By move the motion of swiping the buttons is meant
func (t *touch) shouldTriggerTouchMove() (bool, int, int) {
	dx := t.currX - t.originX
	dy := t.currY - t.originY

	return int(math.Abs(float64(dx))) > MOVE_THRESHOLD || int(math.Abs(float64(dy))) > MOVE_THRESHOLD, dx, dy
}

type pinch struct {
	id1, id2 ebiten.TouchID
	originH  float64
	prevH    float64
}

type pan struct {
	id ebiten.TouchID

	prevX, prevY     int
	originX, originY int
}

type tap struct {
	X, Y int
}

type TouchInput struct {
	input *Input

	x, y float64 // Used for placing the image curr not in use
	zoom float64 // curr not in use

	touchIDs []ebiten.TouchID
	touches  map[ebiten.TouchID]*touch
	pan      *pan
	taps     []tap
	tapped   bool

	canSwipe bool
}

func newTouchInput(i *Input) *TouchInput {
	ti := &TouchInput{
		input:   i,
		touches: map[ebiten.TouchID]*touch{},
	}

	return ti
}

func (g *TouchInput) TouchUpdate() error {
	// Clear the previous frame's taps.
	g.taps = g.taps[:0]

	if len(g.touches) == 0 {
		g.canSwipe = true
	}

	// What touches have just ended?
	for id, t := range g.touches {
		if inpututil.IsTouchJustReleased(id) {

			// If this one has not been touched long (30 frames can be assumed
			// to be 500ms), or moved far, then it's a tap.
			diff := distance(t.originX, t.originY, t.currX, t.currY)
			if !t.wasPinch && !t.isPan && (t.duration <= 30 || diff < 2) {
				g.taps = append(g.taps, tap{
					X: t.currX,
					Y: t.currY,
				})
			}

			delete(g.touches, id)
		}
	}

	// What touches are new in this frame?
	g.touchIDs = inpututil.AppendJustPressedTouchIDs(g.touchIDs[:0])
	for _, id := range g.touchIDs {
		x, y := ebiten.TouchPosition(id)
		g.touches[id] = &touch{
			originX: x, originY: y,
			currX: x, currY: y,
		}
	}

	g.touchIDs = ebiten.AppendTouchIDs(g.touchIDs[:0])

	// Update the current position and durations of any touches that have
	// neither begun nor ended in this frame.
	for _, id := range g.touchIDs {
		t := g.touches[id]
		t.duration = inpututil.TouchPressDuration(id)
		t.currX, t.currY = ebiten.TouchPosition(id)
		shouldTriggerMove, dx, dy := t.shouldTriggerTouchMove()

		if shouldTriggerMove && g.canSwipe {
			if g.input.game.state == co.StateMainMenu {
				g.input.game.state = co.StateRunning
			}
			g.input.SelectMoveDelta(dx, dy)
			g.canSwipe = false
		}

	}

	return nil
}

func (ti *TouchInput) checkTapped() bool {
	if len(ti.taps) == 0 {
		ti.tapped = false
		return false
	} else {
		ti.tapped = true
		return true
	}
}
