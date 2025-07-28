package api

import game "github.com/nlively/gosnake/common/api"

type GoSnakeAPI interface {
	CreateGame(options CreateGameOptions) (*game.Game, error)
	LeaveGame()
}

type CreateGameOptions struct {
	RequestedGridWidth  int
	RequestedGridHeight int
}
