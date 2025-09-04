package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScore(t *testing.T) {
	c := NewCore()
	assert.NotNil(t, c)
	assert.Equal(t, 0, c.Score())

	expect := 1000
	c.AddScore(expect / 2)
	c.AddScore(expect / 2)

	assert.Equal(t, expect, c.Score())
}

func TestSetScore(t *testing.T) {
	c := NewCore()
	assert.NotNil(t, c)
	assert.Equal(t, 0, c.Score())
	c.AddScore(1000)
	assert.Equal(t, 1000, c.Score())

	expect := 0

	c.SetScore(expect)

	assert.Equal(t, expect, c.Score())
}
