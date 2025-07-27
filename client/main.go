package client

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"noahlively.com/snakegame/config"
	"noahlively.com/snakegame/game"
)

/**

func (p *Player) Listen() {
	fmt.Printf("Player.Listen(). IP address %s, port %d\n", p.IPAddress, p.Port)
	addr := net.UDPAddr{
		Port: p.Port,
		IP:   net.ParseIP("0.0.0.0"),
	}

	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Printf("error listening to udp port: %w\n", err)
		panic(err)
	}
	defer conn.Close()

	buf := make([]byte, 1024)

	fmt.Printf("Listening on %s:%d\n", p.IPAddress, p.Port)

	for {
		n, remoteAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}
		fmt.Printf("Received message from %v: %s\n", remoteAddr, string(buf[:n]))
	}
}

func (p *Player) SendMessage(to *Player) {
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		Port: to.Port,
		IP:   net.ParseIP(to.IPAddress),
	})
	if err != nil {
		fmt.Printf("error messaging udp address %s using port %d: %w\n", to.IPAddress, to.Port, err)
		panic(err)
	}
	defer conn.Close()

	msg := fmt.Sprintf("Hello from %s", p.Name)
	_, err = conn.Write([]byte(msg))
	if err != nil {
		fmt.Println("Write error: ", err)
		os.Exit(1)
	}

	fmt.Printf("Message sent to %s:%d: %s\n", to.IPAddress, to.Port, msg)
}

**/

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyM) {
		g.ThisPlayer.SendMessage(g.Player2)
	}
	switch g.State {
	case GameStateSetup:
		g.DelegatedSetup()
	case GameStatePaused:
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.SetState(GameStatePlaying)
		}
	case GameStateOver:
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			g.Initialize()
			g.StartPlaying()
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
		pixels[index] = 0x00 // white, for now
		pixels[index+1] = 0xFF
		pixels[index+2] = 0x00
		pixels[index+3] = 0x00
	}
	// fmt.Printf("Dump of pixels: %v\n", pixels)

	screen.WritePixels(pixels)
}

func (g *game.Game) Draw(screen *ebiten.Image) {
	switch g.State {
	case game.GameStateIntro:
		ebitenutil.DebugPrintAt(screen, "Press enter to start", 1, 1)
	case game.GameStatePlaying:
		g.DrawDotsAndSnake(screen)
		ebitenutil.DebugPrintAt(screen, "Game in progress", 1, 1)
	case game.GameStateOver:
		ebitenutil.DebugPrintAt(screen, "Congrats, you won :)", 1, 1)
		ebitenutil.DebugPrintAt(screen, "Game over :(", 1, 1)
		ebitenutil.DebugPrintAt(screen, "Press enter to play again", 1, 16)
	case game.GameStatePaused:
		g.DrawDotsAndSnake(screen)
		ebitenutil.DebugPrintAt(screen, "Game paused", 1, 1)
		ebitenutil.DebugPrintAt(screen, "Press space to resume", 1, 16)
	}

	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Score: %d points", g.Score), 200, 1)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.gridWidth, g.gridHeight
}

func main() {
	fmt.Println("Main!")
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %w", err)
		fmt.Println("error")
	}

	client := NewGameClient()

	fmt.Printf("Config: %v\n", *cfg)

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
