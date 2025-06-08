package constants

// Gamestates Enum style
type GameState int

const (
	StateRunning GameState = iota + 1
	StateMainMenu
	StateInstructions
	StateGameOver
)

/* variables and constants */
const (
	LOGICAL_WIDTH  int = 640
	LOGICAL_HEIGHT int = 480
	BOARDSIZE      int = 4
)
