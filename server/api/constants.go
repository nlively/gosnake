package api

type ManagedGameState string
type ServerMode string
type GamePoolStatus string

const (
	ManagedGameStatePending   ManagedGameState = "MANAGED_GAME_STATE_PENDING"
	ManagedGameStateActive    ManagedGameState = "MANAGED_GAME_STATE_ACTIVE"
	ManagedGameStateAbandoned ManagedGameState = "MANAGED_GAME_STATE_ABANDONED"
	ServerModeLocal           ServerMode       = "SERVER_MODE_LOCAL"
	ServerModeNetwork         ServerMode       = "SERVER_MODE_NETWORK"
	ServerStateInitializing   GamePoolStatus   = "SERVER_STATE_INITIALIZING"
	ServerStateReady          GamePoolStatus   = "SERVER_STATE_READY"
	ServerStateTerminated     GamePoolStatus   = "SERVER_STATE_TERMINATED"
	MaxGames                                   = 5
	InitTimeoutSeconds                         = 300
)
