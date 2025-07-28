package server

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	game "github.com/nlively/gosnake/common/api"
	"github.com/nlively/gosnake/server/internal/config"
)

type CreateGameOptions struct {
	RequestedGridWidth  int
	RequestedGridHeight int
}

func NewServer(config config.ServerConfig) *GamePool {
	s := &GamePool{
		Games:    make(map[string]*ManagedGameInstance, 0),
		Status:   ServerStateInitializing,
		MaxGames: config.MaxGames,
	}

	return s
}

func (s *GamePool) Run() error {
	// Run periodic cleanup in a separate thread
	go func() {
		for true {
			time.Sleep(time.Minute * 5)
			s.CleanUpGames()
		}
	}()

	s.Status = ServerStateReady
	return nil
}

func (s *GamePool) CreateGame(options CreateGameOptions) (*ManagedGameInstance, error) {
	if s.Status != ServerStateReady {
		return nil, fmt.Errorf("server is not ready. status: %s", s.Status)
	}
	if len(s.Games) >= s.MaxGames {
		return nil, fmt.Errorf("server already has max games %d", s.MaxGames)
	}

	game := game.NewGame(options.RequestedGridWidth, options.RequestedGridHeight)

	id := uuid.New().String()
	managed := &ManagedGameInstance{
		id:        id,
		game:      game,
		createdAt: time.Now(),
		state:     ManagedGameStatePending,
	}

	// Add the game to the managed pool
	s.Games[id] = managed

	return managed, nil
}

func (s *GamePool) JoinGame(game *ManagedGameInstance) error {
	switch game.state {
	case ManagedGameStateActive, ManagedGameStatePending:
		// TODO: make sure player is not already in game
		// TODO: put player into game
		return nil
	case ManagedGameStateAbandoned:
		return fmt.Errorf("game is no longer active")
	default:
		return fmt.Errorf("unrecognized game state: %s", game.state)
	}
}

func (s *GamePool) LeaveGame(game *ManagedGameInstance) error {
	// find player in game

	return nil
}

// Run this periodically
func (s *GamePool) CleanUpGames() {
	// build a new array of games, excluding any
	// that are abandoned
	for i := range s.Games {
		game := s.Games[i]
		include := false
		switch game.state {
		case ManagedGameStateActive:
			include = true
		case ManagedGameStatePending:
			cutoffTime := time.Now().Add(time.Minute * -10)
			if game.createdAt.After(cutoffTime) {
				include = true
			}
		}
		// Delete the game if we haven't determined it
		// should be kept
		if !include {
			delete(s.Games, i)
		}
	}
}
