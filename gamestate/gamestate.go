package gamestate

// GameState - //TODO
type GameState struct {
	// ID      int                     `json:"id"`

	Player  ObjectState   `json:"players"`
	Monkeys []ObjectState `json:"animals"`
	Servers []ObjectState `json:"servers"`
	Walls   []ObjectState `json:"walls"`
}

//NewGameState - convenience func for new game state
func NewGameState() *GameState {
	monkey0state := ObjectState{ID: "monkey0"}
	monkey1state := ObjectState{ID: "monkey1"}
	monkey2state := ObjectState{ID: "monkey2"}
	monkey3state := ObjectState{ID: "monkey3"}
	monkey4state := ObjectState{ID: "monkey4"}
	monkey5state := ObjectState{ID: "monkey5"}
	monkey6state := ObjectState{ID: "monkey6"}
	monkey7state := ObjectState{ID: "monkey7"}
	monkeys := []ObjectState{monkey0state, monkey1state, monkey2state, monkey3state, monkey4state, monkey5state, monkey6state, monkey7state}

	myState := GameState{
		// ID:      0,
		Player:  ObjectState{ID: "player"},
		Monkeys: monkeys,
		Servers: []ObjectState{},
		Walls:   []ObjectState{},
	}
	return &myState
}
