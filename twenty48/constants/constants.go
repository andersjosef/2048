package constants

// Gamestates Enum style
type GameState int

const (
	StateRunning GameState = iota + 1
	StateMainMenu
	StateInstructions
	StateGameOver
)
