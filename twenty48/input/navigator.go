package input

import (
	co "github.com/andersjosef/2048/twenty48/constants"
)

type Navigator struct {
	SwitchState func(co.GameState)
}

func (i *Input) SetNavigator(s func(co.GameState)) {
	i.nav.SwitchState = s
}

func (n *Navigator) Switch(s co.GameState) { n.SwitchState(s) }
