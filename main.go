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
	"log"

	"noahlively.com/snakegame/game"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	GridWidth  = 320
	GridHeight = 240
)

func main() {
	game := game.NewGame(GridWidth, GridHeight)

	game.Initialize()

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Snake")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
