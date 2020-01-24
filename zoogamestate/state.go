package zoogamestate

import "github.com/faiface/pixel"

// GameState - //TODO
type GameState struct {
	ID      int       `json:"id"`
	Players []*Player `json:"players"`
	Animals []*Animal `json:"animals"`
	Cages   []*Cage   `json:"cages"`
	Walls   []*Wall   `json:"walls"`
}

// Animal - //TODO
type Animal struct {
	ID  int `json:"id"`
	X   int `json:"x"`
	Y   int `json:"y"`
	Rot int `json:"rot"`
}

// Wall - //TODO
type Wall struct {
	ID  int `json:"id"`
	X   int `json:"x"`
	Y   int `json:"y"`
	Rot int `json:"rot"`
}

// Cage - ///TODO
type Cage struct {
	ID int `json:"id"`
}

// Player - //TODO
type Player struct {
	ID        int     `json:"id"`
	gravity   float64 `json:"gravity"`
	runSpeed  float64 `json:"runspeed"`
	jumpSpeed float64 `json:""`

	rect   pixel.Rect `json:"rect"`
	vel    pixel.Vec  `json:"vel"`
	ground bool       `json:"ground"`
}
