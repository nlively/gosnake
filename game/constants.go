package game

type Heading string

type DotLevel int

type Speed int

type GameState string

const (
	HeadingLeft      Heading   = "LEFT"
	HeadingRight     Heading   = "RIGHT"
	HeadingUp        Heading   = "UP"
	HeadingDown      Heading   = "DOWN"
	DotLevelTiny     DotLevel  = 1
	DotLevelSmall    DotLevel  = 2
	DotLevelMedium   DotLevel  = 3
	DotLevelLarge    DotLevel  = 4
	DotLevelHuge     DotLevel  = 5
	SpeedVerySlow    Speed     = 1
	SpeedSlow        Speed     = 2
	SpeedMedium      Speed     = 3
	SpeedFast        Speed     = 4
	SpeedVeryFast    Speed     = 5
	GameStateIntro   GameState = "GAME_STATE_INTRO"
	GameStatePlaying GameState = "GAME_STATE_PLAYING"
	GameStateWon     GameState = "GAME_STATE_WON"
	GameStateLost    GameState = "GAME_STATE_LOST"
	GameStatePaused  GameState = "GAME_STATE_PAUSED"
)
