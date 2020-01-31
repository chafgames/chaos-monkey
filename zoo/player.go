package zoo

import (
	"log"

	"github.com/faiface/pixel"
)

// player
type player struct {
	ID      string
	Sprites []*pixel.Sprite
	Score   int
	Health  int
}

func (p *player) draw() {
	myState, present := state.Players[p.ID]
	if present {
		if myState.Active {
			log.Printf("Drawing %s", p.ID)
			p.Sprites[0].Draw(win, myState.IdentityMatrix)
		} else {
			log.Printf("not drawing inactive player %s", p.ID)
		}
	}

}

// func newPlayer() *player {

// 	myObject := newObject()

// 	return &player{Object: myObject, Score: 0, Health: 100}
// }
