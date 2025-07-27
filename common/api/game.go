package api

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

type Game struct {
	GridWidth  int
	GridHeight int
	State      GameState
	Dots       *DoublyLinkedList
	ThisPlayer *Player
	AllPlayers []*Player
	Winner     *Player

	// Private / helper members
	fullGrid *Grid
	dotGrid  *Grid

	rng *rand.Rand
}

func (g *Game) SetState(newState GameState) {
	g.State = newState
}

func NewGame(gridWidth int, gridHeight int) *Game {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	gameGrid := NewGrid(gridWidth, gridHeight)
	dotGrid := NewGrid(gridWidth, gridHeight)

	return &Game{
		GridWidth:  gridWidth,
		GridHeight: gridHeight,
		fullGrid:   gameGrid,
		dotGrid:    dotGrid,
		rng:        rng,
	}
}

// Find a random place on the grid to house a rectangle with the given dimensions
func (g *Game) placeNewRectInFreeSpace(width int, height int, xOffset int, yOffset int) (int, int) {
	const MaxAttempts = 20

	// Random X and Y coords within the grid
	for attempts := 0; attempts < MaxAttempts; attempts++ {
		x := g.rng.Intn(g.GridWidth-(xOffset*2)-width) + xOffset
		y := g.rng.Intn(g.GridHeight-(yOffset*2)-height) + yOffset

		// Loop through grid at given position and figure out whether we're colliding with something already there
		for i := x; i < x+width; i++ {
			for j := y; j < y+height; j++ {
				if g.fullGrid.IsPointFilled(Point{X: i, Y: j}) {
					continue
				}
			}
		}

		return x, y
	}

	// if we can't find something in a fixed set of attempts, panic
	panic(fmt.Sprintf("could not place item in %d attempts", MaxAttempts))
}

func (g *Game) AddNewPlayer(name string) *Player {
	// initial dimensions of snake are 3w x 1h
	x, y := g.placeNewRectInFreeSpace(3, 1, 40, 40)

	snake, err := NewSnake(x, y)
	if err != nil {
		log.Fatalf("error creating snake: %v\n", err)
	}
	fmt.Printf("Starting snake at %d,%d\n", x, y)

	g.fullGrid.PlotPoints(snake.segments)

	return &Player{
		Name:  name,
		Score: 0,
		Lives: 3,
		Snake: snake,
	}
}

func (g *Game) Initialize() {
	initialDots := &DoublyLinkedList{}
	const dotOffset = 3
	const totalDots = 300
	count := 0
	for count < totalDots {
		x, y := g.placeNewRectInFreeSpace(1, 1, dotOffset, dotOffset)
		point := Point{x, y}
		initialDots.InsertAtEnd(NewRandomDot(point))
		g.fullGrid.PlotPoint(point)
		g.dotGrid.PlotPoint(point)
		count++
	}

	g.State = GameStateIntro
	g.Dots = initialDots
}

func (g *Game) StartPlaying() {
	g.SetState(GameStatePlaying)
}
