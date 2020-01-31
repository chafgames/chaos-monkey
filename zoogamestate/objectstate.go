package zoogamestate

import "github.com/faiface/pixel"

// ObjectState - Shared Physics information needed for any object
type ObjectState struct {
	ID     string `json:"id"`
	Active bool   `json:"active"`

	// Gravity   float64 `json:"gravity"`
	// RunSpeed  float64 `json:"runspeed"`
	// JumpSpeed float64 `json:"jumpspeed"`

	// Vel            *pixel.Vec    `json:"vel"`
	// Position       *pixel.Vec    `json:"position"`
	IdentityMatrix pixel.Matrix `json:"identity_matrix"`
}

// NewObjectState - convenience func for new object state
func NewObjectState(id string) *ObjectState {
	myState := ObjectState{
		ID:             id,
		Active:         false,
		IdentityMatrix: pixel.IM,
	}
	return &myState
}
