package api

type GameState string

const (
	GameStateIntro   GameState = "GAME_STATE_INTRO"
	GameStateMenu    GameState = "GAME_STATE_MENU"
	GameStateSetup   GameState = "GAME_STATE_SETUP"
	GameStatePlaying GameState = "GAME_STATE_PLAYING"
	GameStateOver    GameState = "GAME_STATE_OVER"
	GameStatePaused  GameState = "GAME_STATE_PAUSED"
)
