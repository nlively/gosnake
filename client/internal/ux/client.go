package ux

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"github.com/nlively/gosnake/client/internal/config"
	game "github.com/nlively/gosnake/common/api"
	api "github.com/nlively/gosnake/server/api"
)

type ConnectionState string

const (
	GridWidth                                   = 320
	GridHeight                                  = 240
	ConnectionStatePending      ConnectionState = "CONNECTION_STATE_PENDING"
	ConnectionStateReady        ConnectionState = "CONNECTION_STATE_READY"
	ConnectionStateDisconnected ConnectionState = "CONNECTION_STATE_DISCONNECTED"
)

type GameClient struct {
	Config          config.ClientConfig
	APIClient       api.GoSnakeAPI
	Game            *game.Game
	PlayerSessions  []*PlayerSession
	ConnectionState ConnectionState
}

func NewGameClient(config config.ClientConfig) *GameClient {
	client := &GameClient{
		Config: config,
	}

	return client
}

func (c *GameClient) listenToServer(conn net.Conn, readyChan chan<- bool, doneChan chan<- bool) {
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Connection closed by server.")
			os.Exit(0)
		}

		// Check if this is the "ready" message from the server
		if message == "READY\n" {
			readyChan <- true
			close(readyChan)
		}

		fmt.Printf("Server says: %s\n", message)
	}
}

func (c *GameClient) CreateGame() error {
	game, err := c.APIClient.CreateGame(api.CreateGameOptions{
		RequestedGridWidth:  GridWidth,
		RequestedGridHeight: GridHeight,
	})
	if err != nil {
		return err
	}

	c.Game = game

	return nil
}

func (c *GameClient) Run() error {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", c.Config.ServerHost, c.Config.ServerPort))
	if err != nil {
		return err
	}
	defer conn.Close()

	apiClient := api.NewClient(conn)

	readyChan := make(chan bool)
	doneChan := make(chan bool)

	// Listen for messages from the server
	go c.listenToServer(conn, readyChan, doneChan)

	// Wait for server to be "ready"
	<-readyChan
	fmt.Println("Server is ready")

	go c.CreateGame()

	<-doneChan
	fmt.Println("Server session terminated")
	return nil
}
