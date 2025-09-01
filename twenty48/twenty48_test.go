package twenty48

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReset(t *testing.T) {
	g, err := NewGame()
	assert.NoError(t, err)

	g.score = 1000

	ResetGame(g.input)
	g.eventBus.Dispatch()

	assert.Equal(t, 0, g.score)
}
