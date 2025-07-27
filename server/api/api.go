package api

type GoSnakeAPI interface {
	CreateGame(options CreateGameOptions) error
	LeaveGame()
}

type CreateGameOptions struct {
	RequestedGridWidth  int
	RequestedGridHeight int
}
