package api

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"

	game "github.com/nlively/gosnake/common/api"
	"github.com/nlively/gosnake/server/internal/server"
)

type APIClient struct {
	conn net.Conn
}

func NewClient(conn net.Conn) *APIClient {
	return &APIClient{
		conn: conn,
	}
}

func Connect(host string, port int) (net.Conn, error) {
	// TODO: establish TCP/IP connection
	return nil, nil
}

func (c *APIClient) doAPISend(command server.ServerCommand) (interface{}, error) {
	data, err := json.Marshal(command)
	if err != nil {
		return nil, err
	}

	data = append(data, '\n')
	_, err = c.conn.Write(data)
	if err != nil {
		return nil, err
	}

	reader := bufio.NewReader(c.conn)
	response, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *APIClient) CreateGame(options CreateGameOptions) (*game.Game, error) {
	command := server.NewCreateGameCommand(options.RequestedGridWidth, options.RequestedGridHeight)

	result, err := c.doAPISend(command)
	if err != nil {
		return nil, err
	}

	// TODO: translate result into typed response object
	fmt.Printf("CreateGame API result: %v\n", result)

	return nil, nil
}
