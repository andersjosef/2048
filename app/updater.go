package app

import "github.com/hajimehoshi/ebiten/v2"

type Updater interface{ Update() error }

type Drawer interface{ Draw(screen *ebiten.Image) }
