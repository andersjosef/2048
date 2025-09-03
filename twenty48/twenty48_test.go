package twenty48

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReset(t *testing.T) {
	g, err := NewGame(Deps{})
	assert.NoError(t, err)
	cmds := NewCommands(g)

	g.score = 1000

	cmds.ResetGame()

	g.EventBus.Dispatch()

	assert.Equal(t, 0, g.score)
}
