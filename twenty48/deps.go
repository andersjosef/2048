package twenty48

import co "github.com/andersjosef/2048/twenty48/constants"

type Deps struct {
	FMS
	IsGameOver func() bool
}

type FMS interface {
	Current() co.GameState
	Previous() co.GameState
	Has(id co.GameState) bool
	Switch(id co.GameState)
}
