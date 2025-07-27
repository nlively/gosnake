package server

type GamePoolStatus string
type ManagedGameState string
type ServerMode string

const (
	ServerStateInitializing   GamePoolStatus   = "SERVER_STATE_INITIALIZING"
	ServerStateReady          GamePoolStatus   = "SERVER_STATE_READY"
	ServerStateTerminated     GamePoolStatus   = "SERVER_STATE_TERMINATED"
	ManagedGameStatePending   ManagedGameState = "MANAGED_GAME_STATE_PENDING"
	ManagedGameStateActive    ManagedGameState = "MANAGED_GAME_STATE_ACTIVE"
	ManagedGameStateAbandoned ManagedGameState = "MANAGED_GAME_STATE_ABANDONED"
	ServerModeLocal           ServerMode       = "SERVER_MODE_LOCAL"
	ServerModeNetwork         ServerMode       = "SERVER_MODE_NETWORK"
	MaxGames                                   = 5
	InitTimeoutSeconds                         = 300
)
