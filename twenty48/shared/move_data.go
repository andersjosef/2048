package shared

type MoveData struct {
	ScoreGain  int
	IsGameOver bool
	MoveDeltas []MoveDelta
	Dir        string
	NewBoard   [4][4]int
}

// Describes a single tile moving from one tile to another
type MoveDelta struct {
	FromRow, FromCol int
	ToRow, ToCol     int
	ValueMoved       int
	Merged           bool
}
