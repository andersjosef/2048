package app

type Updater interface{ Update() error }

type updaterFunc func() error

func (f updaterFunc) Update() error { return f() }
