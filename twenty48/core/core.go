package core

type Core struct {
	score int
}

func NewCore() *Core { return &Core{} }

func (c *Core) AddScore(n int) { c.score += n }

func (c *Core) SetScore(n int) { c.score = n }

func (c *Core) Score() int { return c.score }

func (c *Core) Reset() { c.SetScore(0) }
