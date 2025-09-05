package eventhandler

type EventType int

const (
	EventScreenChanged EventType = iota
	EventResetGame
	EventMoveMade
	EventThemeChanged
	EventScaleBoardView
)

type Event struct {
	Type EventType
	Data any // If any data should be sent with the emit
}

type EventBus struct {
	listeners map[EventType][]func(Event) // Participants
	queue     []Event                     // Events to be processed
}

func NewEventBus() *EventBus {
	return &EventBus{
		listeners: make(map[EventType][]func(Event)),
	}
}

// Register funciton that is going to be triggered by emit
func (b *EventBus) Register(eventType EventType, handler func(Event)) {
	b.listeners[eventType] = append(b.listeners[eventType], handler)
}

// Trigger an event type for all registered functions
func (b *EventBus) Emit(event Event) {
	b.queue = append(b.queue, event)
}

// Swallow away events form the queue
func (b *EventBus) Dispatch() {
	lenStart := len(b.queue)
	for _, evt := range b.queue {
		if handlers, ok := b.listeners[evt.Type]; ok {
			for _, h := range handlers {
				h(evt)
			}
		}
	}

	// In case handlers are emitted in an registered function
	if len(b.queue) > lenStart {
		b.queue = b.queue[lenStart-1:]
		b.Dispatch()
	}

	b.queue = b.queue[:0] // Clear queue after all event have been processed
}
