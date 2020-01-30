package zoo

import "github.com/mattmulhern/game-off-2019-scratch/zoogamestate"

// player
type player struct {
	Object *object
	Score  int
	Health int
}

func newPlayer(objStatePtr *zoogamestate.ObjectState) *player {
	myObject := newObject(objStatePtr)

	return &player{Object: myObject, Score: 0, Health: 100}
}
