package client

import (
	"bufio"
	"fmt"
	"net"
	"os"

	game "github.com/nlively/gosnake/common/api"
	api "github.com/nlively/gosnake/server/api"
)

const (
	GridWidth  = 320
	GridHeight = 240
)

type GameClient struct {
	Config         *ClientConfig
	APIClient      api.GoSnakeAPI
	Game           *game.Game
	PlayerSessions []*PlayerSession
}

func NewGameClient() *GameClient {

	return nil
}

func (c *GameClient) CreateGame() error {
	game := c.APIClient.CreateGame(api.CreateGameOptions{
		RequestedGridWidth:  GridWidth,
		RequestedGridHeight: GridHeight,
	})

	return nil
}

func (c *GameClient) Run() error {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", c.APIClient.Host, c.APIClient.Port))
	if err != nil {
		return err
	}
	defer conn.Close()

	c.conn = conn

	// Listen for messages from the server
	go func() {
		reader := bufio.NewReader(conn)
		for {
			message, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Connection closed by server.")
				os.Exit(0)
			}
			fmt.Printf("Server says: %s\n", message)
		}
	}()

	err = c.CreateGame()
	if err != nil {
		return err
	}

	return nil
}
