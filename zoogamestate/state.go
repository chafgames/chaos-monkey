package zoogamestate

import "github.com/faiface/pixel"

// GameState - //TODO
type GameState struct {
	ID      int                `json:"id"`
	Players map[string]*Player `json:"players"`
	Animals map[string]*Animal `json:"animals"`
	Cages   map[string]*Cage   `json:"cages"`
	Walls   map[string]*Wall   `json:"walls"`
}

// Animal - //TODO
type Animal struct {
	ID     int  `json:"id"`
	Active bool `json:"active"`
	X      int  `json:"x"`
	Y      int  `json:"y"`
	Rot    int  `json:"rot"`
}

// Wall - //TODO
type Wall struct {
	ID     int  `json:"id"`
	Active bool `json:"active"`
	X      int  `json:"x"`
	Y      int  `json:"y"`
	Rot    int  `json:"rot"`
}

// Cage - ///TODO
type Cage struct {
	ID     int  `json:"id"`
	Active bool `json:"active"`
}

// Player - //TODO
type Player struct {
	ID     string `json:"id"`
	Active bool   `json:"active"`

	gravity   float64 `json:"gravity"`
	runSpeed  float64 `json:"runspeed"`
	jumpSpeed float64 `json:""`

	rect   pixel.Rect `json:"rect"`
	vel    pixel.Vec  `json:"vel"`
	ground bool       `json:"ground"`
}
