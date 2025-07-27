package server

type ServerCommand struct {
	Command string                 `json:"command"`
	Params  map[string]interface{} `json:"params"`
}

func NewCreateGameCommand(gridWidth int, gridHeight int) ServerCommand {
	return ServerCommand{
		Command: "create_game",
		Params: map[string]interface{}{
			"grid_width":  gridWidth,
			"grid_height": gridHeight,
		},
	}
}
