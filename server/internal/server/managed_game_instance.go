package server

import (
	"time"

	game "github.com/nlively/gosnake/common/api"
)

type ManagedGameInstance struct {
	id        string
	game      *game.Game
	createdAt time.Time
	state     ManagedGameState
}
