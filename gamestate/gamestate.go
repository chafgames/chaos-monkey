package gamestate

// GameState - //TODO
type GameState struct {
	// ID      int                     `json:"id"`
	Players map[string]*ObjectState `json:"players"`
	Animals map[string]*ObjectState `json:"animals"`
	Cages   map[string]*ObjectState `json:"cages"`
	Walls   map[string]*ObjectState `json:"walls"`
}

//NewGameState - convenience func for new game state
func NewGameState() *GameState {
	players := make(map[string]*ObjectState)
	animals := make(map[string]*ObjectState)
	cages := make(map[string]*ObjectState)
	walls := make(map[string]*ObjectState)
	myState := GameState{
		// ID:      0,
		Players: players,
		Animals: animals,
		Cages:   cages,
		Walls:   walls,
	}
	return &myState
}
