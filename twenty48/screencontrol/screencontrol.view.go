package screencontrol

import "github.com/andersjosef/2048/twenty48/eventhandler"

type Deps struct {
	EventHandler
}

type EventHandler interface {
	Emit(event eventhandler.Event)
}
