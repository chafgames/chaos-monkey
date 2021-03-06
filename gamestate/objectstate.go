package gamestate

import "github.com/faiface/pixel"

// ObjectState - Shared Physics information needed for any object
type ObjectState struct {
	ID      string `json:"id"`
	Active  bool   `json:"active"`
	CurAnim string // key in Player Sprites Map to animate

	IdentityMatrix pixel.Matrix `json:"identity_matrix"`
}

// NewObjectState - convenience func for new object state
func NewObjectState(id string) *ObjectState {
	myState := ObjectState{
		ID:             id,
		Active:         false,
		IdentityMatrix: pixel.IM,
		CurAnim:        "S",
	}
	return &myState
}
