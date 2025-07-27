package server

type GamePool struct {
	Games    map[string]*ManagedGameInstance
	Status   GamePoolStatus
	MaxGames int
}
