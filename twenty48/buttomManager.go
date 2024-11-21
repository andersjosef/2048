package twenty48

type ButtonManager struct {
	game        *Game
	buttonArray []Button
}

func InitButtonManager(g *Game) *ButtonManager {
	var bm *ButtonManager = &ButtonManager{
		game: g,
	}

	return bm
}

type Button struct {
	game *Game
	// TODO: Need limits and init method bound to button manager
}

func (bu *Button) onTrigger(actionFunction ActionFunc) {
	actionFunction(bu.game.input)

}
