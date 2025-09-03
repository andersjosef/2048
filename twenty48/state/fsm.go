package state

import "github.com/hajimehoshi/ebiten/v2"

type State interface {
	Enter()
	Exit()
	Update() error
	Draw(screen *ebiten.Image)
}

type FSM[T comparable] struct {
	current T
	prev    T
	states  map[T]State
}

func NewFSM[T comparable]() *FSM[T] { return &FSM[T]{states: make(map[T]State)} }

func (f *FSM[T]) Register(id T, s State) { f.states[id] = s }

func (f *FSM[T]) Has(id T) bool { _, ok := f.states[id]; return ok }

func (f *FSM[T]) Current() T  { return f.current }
func (f *FSM[T]) Previous() T { return f.current }

// Set the initial state to be
func (f *FSM[T]) Start(id T) {
	next, ok := f.states[id]
	if !ok {
		return
	}
	f.current = id
	next.Enter()
}

func (f *FSM[T]) Switch(id T) {
	if f.current == id {
		return
	}
	next, ok := f.states[id]
	if !ok {
		return
	}
	if current, ok := f.states[id]; ok {
		current.Exit()
	}
	f.prev = f.current
	f.current = id
	next.Enter()
}

func (f *FSM[T]) Update() error {
	if s, ok := f.states[f.current]; ok {
		return s.Update()
	}
	return nil // TODO: reevaluate
}

func (f *FSM[T]) Draw(screen *ebiten.Image) {
	if s, ok := f.states[f.current]; ok {
		s.Draw(screen)
	}
}
