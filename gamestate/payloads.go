package gamestate

// PlayerUpdate - used for clients to submit state updates
type PlayerUpdate struct {
	ID    string       `json:"id"`
	State *ObjectState `json:"state"`
}
