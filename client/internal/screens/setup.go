package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type GameSetup struct {
	inputPrompt string
	ticker      int
	runes       []rune
	playerName  string
	game        *Game
}

func repeatingKeyPressed(key ebiten.Key) bool {
	const (
		delay    = 30
		interval = 3
	)
	d := inpututil.KeyPressDuration(key)
	if d == 1 {
		return true
	}
	if d >= delay && (d-delay)%interval == 0 {
		return true
	}
	return false
}

func (s *GameSetup) DelegatedUpdate() {
	// Add runesthat are input by the user by AppendInputChars
	// Note that AppendInputchars result changes everyf rame, so you need
	// to call this every frame
	s.runes = ebiten.AppendInputChars(s.runes[:0])
	s.playerName += string(s.runes)

	// Adjust the string to be at most 20 characters
	if len(s.playerName) > 20 {
		s.playerName = s.playerName[0:20]
	}

	// handle backspace
	if repeatingKeyPressed(ebiten.KeyBackspace) {
		if len(s.playerName) >= 1 {
			s.playerName = s.playerName[:len(s.playerName)-1]
		}
	}

	s.inputPrompt = fmt.Sprintf("Enter your name: %s", s.playerName)
}

func (s *GameSetup) DelegatedDraw(screen *ebiten.Image) {
	t := s.inputPrompt
	// Blink 50% of the time
	if s.ticker%60 < 30 {
		t += "_"
	}
	ebitenutil.DebugPrint(screen, t)
}
