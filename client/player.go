package client

import (
	"log"

	"github.com/chafgames/chaos-monkey/gamestate"
	"github.com/faiface/pixel"
)

// player
type player struct {
	ID          string
	State       *gamestate.ObjectState
	Sprites     []*pixel.Sprite
	IsMonkey    bool
	MonkeyIndex int
	Score       int
	Health      int
}

func (p *player) draw() {
	// if present {
	if p.State.Active {
		log.Printf("Drawing %s", p.ID)
		p.Sprites[0].Draw(win, p.State.IdentityMatrix)
	} else {
		log.Printf("not drawing inactive player %s", p.ID)
	}
	// }
}

// func newPlayer() *player {

// 	myObject := newObject()

// 	return &player{Object: myObject, Score: 0, Health: 100}
// }
