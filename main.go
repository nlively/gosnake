package main

/***

	Snake, by Noah
	July, 2025

	Objective:

	Achieve the highest score by growing the snake without crashing.

	Rules:

	* Player loses if snake collides with a wall or with itself
	* Snake grows by consuming dots
	* The longer the snake, the higher the score

	Controls:

	- Move the snake with the arrow keys
	- Pause / Unpause the game using the space bar

 ***/

import (
	"fmt"
	"log"

	"noahlively.com/snakegame/config"
	"noahlively.com/snakegame/game"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	GridWidth  = 320
	GridHeight = 240
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Errorf("failed to load config: %w", err)
	}

	player1 := &game.Player{
		ID:        "player1",
		Name:      "Abbot",
		IPAddress: "127.0.0.1",
		Port:      3000,
	}

	player2 := &game.Player{
		ID:        "player2",
		Name:      "Costello",
		IPAddress: cfg.PeerAddress,
		Port:      cfg.PeerPort,
	}

	player1.Listen()

	player1.SendMessage(player2)

	game := game.NewGame(GridWidth, GridHeight)

	game.Initialize()

	ebiten.SetWindowSize(1024, 768)
	ebiten.SetWindowTitle("Snake")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
