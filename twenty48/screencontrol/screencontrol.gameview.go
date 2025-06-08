package screencontrol

import "github.com/andersjosef/2048/twenty48/eventhandler"

type GameView interface {
	GetBusHandler() *eventhandler.EventBus
}

// type GameProvider interface {
// 	GetScale() float64
// }
