package api

type Player struct {
	ID        string
	Name      string
	IPAddress string
	Port      int
	Score     int
	Lives     int
	Snake     *Snake
	Connected bool
}
