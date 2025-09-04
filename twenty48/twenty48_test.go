package twenty48

import (
	"testing"

	co "github.com/andersjosef/2048/twenty48/constants"
	"github.com/stretchr/testify/assert"
)

type MockFSM struct{}

func (MockFSM) Current() co.GameState {
	return co.StateMainMenu
}
func (MockFSM) Previous() co.GameState {
	return co.StateInstructions
}
func (MockFSM) Has(co.GameState) bool {
	return true
}
func (MockFSM) Switch(co.GameState) {}

func TestReset(t *testing.T) {
	g, err := NewRouter(Deps{
		FSM: MockFSM{},
	})
	assert.NoError(t, err)
	cmds := NewCommands(g)

	g.Core.SetScore(1000)

	cmds.ResetGame()

	g.EventBus.Dispatch()

	assert.Equal(t, 0, g.Core.Score())
}
