package screencontrol

type GameView interface {
	GameProvider
}

type GameProvider interface {
	GetScale() float64
}
