package ux

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	game "github.com/nlively/gosnake/common/api"
)

type GameUXState string

const (
	GameUXStateIntro   GameUXState = "GAME_UX_STATE_INTRO"
	GameUXStatePlaying GameUXState = "GAME_UX_STATE_PLAYING"
	GameUXStateOver    GameUXState = "GAME_UX_STATE_OVER"
	GameUXStatePaused  GameUXState = "GAME_UX_STATE_PAUSED"
)

type GameUXEngine struct {
	Client  *GameClient
	UXState GameUXState
}

func (g *GameUXEngine) Update() error {
	switch g.State {
	case game.GameStateSetup:
		g.DelegatedSetup()
	case game.GameStatePaused:
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.SetState(game.GameStatePlaying)
		}
	case game.GameStateOver:
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			g.Initialize()
			g.StartPlaying()
		}
	case game.GameStatePlaying:
		// Respond to arrow keys for changing direction, but don't allow a direct reversal without a turn
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.SetState(game.GameStatePaused)
		} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) && g.Snake.heading != game.HeadingDown {
			g.Snake.SetHeading(game.HeadingUp)
		} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) && g.Snake.heading != game.HeadingLeft {
			g.Snake.SetHeading(game.HeadingRight)
		} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) && g.Snake.heading != game.HeadingUp {
			g.Snake.SetHeading(game.HeadingDown)
		} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) && g.Snake.heading != game.HeadingRight {
			g.Snake.SetHeading(game.HeadingLeft)
		}

		g.Snake.Move()

		// Detect collision of snake with itself
		if g.Snake.HasCollisionWithSelf() {
			fmt.Println("Snake has collided with itself. Game over")
			// Lose the game
			g.SetState(game.GameStateLost)
		}

		snakeTip := g.Snake.GetTip()

		// Detect collision of snake with walls
		if snakeTip.X < 0 || snakeTip.X > g.gridWidth || snakeTip.Y < 0 || snakeTip.Y > g.gridHeight {
			fmt.Println("Snake has collided with a wall. Game over")
			// Lose the game
			g.SetState(game.GameStateLost)
		}

		// Detect collision of snake with a dot
		if g.dotGrid.IsPointFilled(snakeTip) {
			// Consume the dot, grow the snake, increment the score
			fmt.Printf("Snake has eaten a dot at %d,%d\n", snakeTip.X, snakeTip.Y)

			// Increment the score
			g.Score++

			// Grow the snake
			g.Snake.Grow()

			// Unplot the dot
			g.dotGrid.UnplotPoint(snakeTip)

			// Remove the dot from the array
			g.Dots.RemoveByCoordinates(snakeTip)
		}
	case game.GameStateIntro:
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			g.Initialize()
			g.StartPlaying()
		}
	}

	return nil
}

func (g *GameUXEngine) DrawDotsAndSnake(screen *ebiten.Image) {
	pixels := make([]byte, 4*g.gridWidth*g.gridHeight)
	currentDot := g.Dots.head
	count := 0
	for currentDot != nil {
		count++
		index := (currentDot.dot.X + (g.gridWidth * currentDot.dot.Y)) * 4
		pixels[index] = 0xFF // white, for now
		pixels[index+1] = 0xFF
		pixels[index+2] = 0xFF
		pixels[index+3] = 0xFF

		currentDot = currentDot.nextNode
	}

	for i := range g.Snake.segments {
		point := g.Snake.segments[i]

		index := (point.X + (g.gridWidth * point.Y)) * 4
		pixels[index] = 0x00 // white, for now
		pixels[index+1] = 0xFF
		pixels[index+2] = 0x00
		pixels[index+3] = 0x00
	}
	// fmt.Printf("Dump of pixels: %v\n", pixels)

	screen.WritePixels(pixels)
}

func (g *GameUXEngine) Draw(screen *ebiten.Image) {
	switch g.UXState {
	case GameUXStateIntro:
		ebitenutil.DebugPrintAt(screen, "Press enter to start", 1, 1)
	case GameUXStatePlaying:
		g.DrawDotsAndSnake(screen)
		ebitenutil.DebugPrintAt(screen, "Game in progress", 1, 1)
	case GameUXStateOver:
		ebitenutil.DebugPrintAt(screen, "Congrats, you won :)", 1, 1)
		ebitenutil.DebugPrintAt(screen, "Game over :(", 1, 1)
		ebitenutil.DebugPrintAt(screen, "Press enter to play again", 1, 16)
	case GameUXStatePaused:
		g.DrawDotsAndSnake(screen)
		ebitenutil.DebugPrintAt(screen, "Game paused", 1, 1)
		ebitenutil.DebugPrintAt(screen, "Press space to resume", 1, 16)
	}

	// ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Score: %d points", g.Score), 200, 1)

}

func (g *GameUXEngine) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return GridWidth, GridHeight
}
