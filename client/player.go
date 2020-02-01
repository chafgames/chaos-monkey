package client

import (
	"encoding/json"
	"fmt"
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
	if p.State.Active {
		p.submitUpdate()
		// log.Printf("Drawing %s at %+v", p.ID, p.State.IdentityMatrix)
		p.Sprites[0].Draw(win, p.State.IdentityMatrix)
		// p.Sprites[0].Draw(win, pixel.IM.Moved(win.Bounds().Center())) //puts manual monkley butt at centre of screen
	}
}
func (p *player) submitUpdate() {
	topicString := fmt.Sprintf("PLAYER-UPDATE")
	update := gamestate.PlayerUpdate{ID: p.ID, State: p.State}
	payload, jsonErr := json.Marshal(update)
	log.Printf("submitting  %+v", string(payload))

	if jsonErr != nil {
		log.Printf("ERROR: Failed to Marshal State : %s", jsonErr)
		return
	}
	client.SocketioClient.Emit(topicString, string(payload))
}
