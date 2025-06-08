package screencontrol

import "github.com/andersjosef/2048/twenty48/eventhandler"

type View interface {
	GetBusHandler() *eventhandler.EventBus
}
