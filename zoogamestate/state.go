package zoogamestate

import "github.com/faiface/pixel"

// GameState - //TODO
type GameState struct {
	ID      int
	Players []*Player
	Animals []*Animal
	Cages   []*Cage
	Walls   []*Wall
}

type Animal struct {
	ID  int
	X   int
	Y   int
	Rot int
}

type Wall struct {
	ID  int
	X   int
	Y   int
	Rot int
}

type Cage struct {
	ID int
}

// Player - //TODO
type Player struct {
	ID        int
	gravity   float64
	runSpeed  float64
	jumpSpeed float64

	rect   pixel.Rect
	vel    pixel.Vec
	ground bool
}
