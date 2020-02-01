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
		p.Sprites[0].Draw(win, p.State.IdentityMatrix)
	}
}
func (p *player) submitUpdate() {
	topicString := fmt.Sprintf("PLAYER-UPDATE")
	update := gamestate.PlayerUpdate{ID: p.ID, State: p.State}
	payload, jsonErr := json.Marshal(update)
	if jsonErr != nil {
		log.Printf("ERROR: Failed to Marshal State : %s", jsonErr)
		return
	}
	mySIOClient.Emit(topicString, string(payload))
}
