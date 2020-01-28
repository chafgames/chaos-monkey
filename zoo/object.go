package zoo

import zoogamestate "github.com/mattmulhern/game-off-2019-scratch/zoogamestate"

//object - coupling of physics and drawing info
type object struct {
	State    *zoogamestate.ObjectState
	Graphics *ObjectGraphics
}

// newObject : instantiate an object
func newObject(statePtr *zoogamestate.ObjectState) *object {
	myGraphics := newObjectGraphics()
	return &object{State: statePtr, Graphics: myGraphics}
}
