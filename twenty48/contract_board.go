package twenty48

import (
	"github.com/andersjosef/2048/twenty48/board"
	co "github.com/andersjosef/2048/twenty48/constants"
	"github.com/andersjosef/2048/twenty48/theme"
	"github.com/hajimehoshi/ebiten/v2"
)

type Board interface {
	Draw(screen *ebiten.Image)
	DrawBackgoundBoard(screen *ebiten.Image)
	GetBoardDimentions() (x, y int)
	DrawMovingMatrix(
		screen *ebiten.Image,
		x,
		y int,
		movDistX,
		movDistY float32,
	)
	CreateBoardImage()
	ScaleBoard()
	Move(dir board.Direction)
}

func NewBoard(g *Game) Board {
	d := board.Deps{
		EventHandler:      g.eventBus,
		ScreenControl:     g.screenControl,
		SetGameOver:       func(isGameOver bool) { g.gameOver = isGameOver },
		SetGameState:      func(gs co.GameState) { g.state = gs },
		IsGameOver:        func() bool { return g.gameOver },
		GetCurrentTheme:   func() theme.Theme { return g.currentTheme },
		GetCurrentFontSet: func() theme.FontSet { return *g.fontSet },
	}

	b, err := board.New(d)
	if err != nil {
		panic(err)
	}

	return b
}
