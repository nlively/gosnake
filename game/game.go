package game

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	gridWidth  int
	gridHeight int
	fullGrid   *Grid
	dotGrid    *Grid
	Snake      *Snake
	State      GameState
	Dots       *DoublyLinkedList
	Score      int
}

func (g *Game) SetState(newState GameState) {
	g.State = newState
}

func NewGame(gridWidth int, gridHeight int) *Game {
	return &Game{gridWidth: gridWidth, gridHeight: gridHeight}
}

func (g *Game) Initialize() {
	gameGrid := NewGrid(g.gridWidth, g.gridHeight)
	dotGrid := NewGrid(g.gridWidth, g.gridHeight)

	// Random X and Y coords within the grid
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	const (
		xOffset = 40
		yOffset = 40
	)

	x := r.Intn(g.gridWidth-(xOffset*2)) + xOffset
	y := r.Intn(g.gridHeight-(yOffset*2)) + yOffset

	snake, err := NewSnake(x, y)
	if err != nil {
		log.Fatalf("error creating snake: %v\n", err)
	}
	fmt.Printf("Starting snake at %d,%d\n", x, y)
	gameGrid.PlotPoints(snake.segments)

	initialDots := &DoublyLinkedList{}
	const dotOffset = 3
	const totalDots = 200
	count := 0
	for count < totalDots {
		x := r.Intn(g.gridWidth-(dotOffset*2)) + dotOffset
		y := r.Intn(g.gridHeight-(dotOffset*2)) + dotOffset
		point := Point{x, y}
		if !gameGrid.IsPointFilled(point) {
			initialDots.InsertAtEnd(NewRandomDot(point))
			gameGrid.PlotPoint(point)
			dotGrid.PlotPoint(point)
			fmt.Printf("Dot generated at %d,%d\n", point.X, point.Y)
			count++
		}
	}

	g.fullGrid = gameGrid
	g.dotGrid = dotGrid
	g.Snake = snake
	g.State = GameStateIntro
	g.Score = 0
	g.Dots = initialDots
}

func (g *Game) StartPlaying() {
	g.SetState(GameStatePlaying)
}

func (g *Game) Update() error {
	switch g.State {
	case GameStatePaused:
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.SetState(GameStatePlaying)
		}
	case GameStatePlaying:
		// Respond to arrow keys for changing direction, but don't allow a direct reversal without a turn
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.SetState(GameStatePaused)
		} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) && g.Snake.heading != HeadingDown {
			g.Snake.SetHeading(HeadingUp)
		} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) && g.Snake.heading != HeadingLeft {
			g.Snake.SetHeading(HeadingRight)
		} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) && g.Snake.heading != HeadingUp {
			g.Snake.SetHeading(HeadingDown)
		} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) && g.Snake.heading != HeadingRight {
			g.Snake.SetHeading(HeadingLeft)
		}

		g.Snake.Move()

		// Detect collision of snake with itself
		if g.Snake.HasCollisionWithSelf() {
			fmt.Println("Snake has collided with itself. Game over")
			// Lose the game
			g.SetState(GameStateLost)
		}

		snakeTip := g.Snake.GetTip()

		// Detect collision of snake with walls
		if snakeTip.X < 0 || snakeTip.X > g.gridWidth || snakeTip.Y < 0 || snakeTip.Y > g.gridHeight {
			fmt.Println("Snake has collided with a wall. Game over")
			// Lose the game
			g.SetState(GameStateLost)
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
	case GameStateIntro:
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			g.Initialize()
			g.StartPlaying()
		}
	}

	return nil
}

func (g *Game) DrawDotsAndSnake(screen *ebiten.Image) {
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
		pixels[index] = 0xFF // white, for now
		pixels[index+1] = 0xFF
		pixels[index+2] = 0xFF
		pixels[index+3] = 0xFF
	}
	// fmt.Printf("Dump of pixels: %v\n", pixels)

	screen.WritePixels(pixels)
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch g.State {
	case GameStateIntro:
		ebitenutil.DebugPrint(screen, "Press any key to start")
	case GameStatePlaying:
		g.DrawDotsAndSnake(screen)
		ebitenutil.DebugPrint(screen, "Game in progress")
	case GameStateWon:
		ebitenutil.DebugPrint(screen, "Congrats, you won :)")
	case GameStateLost:
		ebitenutil.DebugPrint(screen, "Game over :/")
	case GameStatePaused:
		g.DrawDotsAndSnake(screen)
		ebitenutil.DebugPrint(screen, "Game paused. Press space to resume")
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.gridWidth, g.gridHeight
}
