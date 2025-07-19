package game

import (
	"fmt"
	"math/rand"
	"time"
)

type Snake struct {
	x          int
	y          int
	length     int // This is the current length of the snake, and should be consisten with what gets drawn
	nextLength int // When the snake earns growth, NextLength allows us to increment the growth overtime
	heading    Heading
	segments   []Point // Starting from the tip and representing the trail of recently drawn pixels, up to Length
	// Speed      Speed
}

func (s *Snake) Move() {
	nextX := s.x
	nextY := s.y
	const moveFactor = 1

	switch s.heading {
	case HeadingRight:
		nextX += moveFactor
	case HeadingDown:
		nextY += moveFactor
	case HeadingLeft:
		nextX -= moveFactor
	case HeadingUp:
		nextY -= moveFactor
	default:
		panic(fmt.Sprintf("unrecognized heading %s", s.heading))
	}

	// Grow the snake incrementally if growth is in order
	if s.nextLength > s.length {
		s.length++
	}

	// Point operates like a stack.  We move the snake by pushing
	// onto the end of the stack and, if necessary, shifting off
	// from the beginning of the stack
	s.segments = append(s.segments, Point{X: nextX, Y: nextY})

	segmentsToDelete := len(s.segments) - s.length
	if segmentsToDelete > 0 {
		// Shave `segmentsToDelete` items from the front of the slice
		s.segments = s.segments[segmentsToDelete:] // sic?
	}

	s.x = nextX
	s.y = nextY

	fmt.Printf("Advancing snake %s to %d,%d. Length is %d, next length is %d\n", s.heading, nextX, nextY, s.length, s.nextLength)
}

func (s *Snake) SetHeading(heading Heading) {
	s.heading = heading
}

func (s *Snake) GetTip() Point {
	return Point{s.x, s.y}
}

func (s *Snake) HasCollisionWithSelf() bool {
	// Obtain the "tip" of the snake
	// Iterate through the other points and return true if any of them match
	tip := s.GetTip()
	if s.length < 5 {
		return false
	}
	for i := 0; i < len(s.segments)-1; i++ {
		if s.segments[i] == tip {
			return true
		}
	}

	return false
}

func (s *Snake) Grow() {
	s.nextLength++
}

func NewSnake(startX int, startY int) (*Snake, error) {

	snake := &Snake{
		x:          startX,
		y:          startY,
		length:     3,
		nextLength: 3,
		heading:    HeadingRight,
		segments:   make([]Point, 0),
	}

	return snake, nil
}

type Dot struct {
	X     int
	Y     int
	Level DotLevel
}

func NewRandomDot(point Point) *Dot {
	// randomize level
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	level := DotLevel(r.Intn(4) + 1)

	return &Dot{point.X, point.Y, level}
}
