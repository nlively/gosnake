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

	"github.com/hajimehoshi/ebiten/v2"
	"noahlively.com/snakegame/config"
	"noahlively.com/snakegame/game"
)

const (
	GridWidth  = 320
	GridHeight = 240
)

func main() {
	fmt.Println("Main!")
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %w", err)
		fmt.Println("error")
	}

	fmt.Printf("Config: %v\n", *cfg)

	player1 := &game.Player{
		ID:        "player1",
		Name:      "Abbot",
		IPAddress: "127.0.0.1",
		Port:      cfg.LocalPort,
	}

	player2 := &game.Player{
		ID:        "player2",
		Name:      "Costello",
		IPAddress: cfg.PeerAddress,
		Port:      cfg.PeerPort,
	}

	fmt.Printf("Player 1: %v\n", *player1)
	fmt.Printf("Player 2: %v\n", *player2)

	go player1.Listen()

	go player1.SendMessage(player2)

	game := game.NewGame(GridWidth, GridHeight)
	game.Player1 = player1
	game.Player2 = player2

	fmt.Printf("Game: %v\n", *game)

	game.Initialize()

	ebiten.SetWindowSize(1024, 768)
	ebiten.SetWindowTitle("Snake")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
