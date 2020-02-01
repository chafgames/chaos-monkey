package client

import (
	"encoding/json"
	"log"

	"fmt"
	"github.com/chafgames/chaos-monkey/gamestate"
	"github.com/faiface/pixel"
)

// player
type player struct {
	ID          string
	State       *gamestate.ObjectState
	Sprites     map[string][]*pixel.Sprite // map of list of sprites representing an animation each
	LastAnimIdx int                        // index of last drawn sprite in list referenced by p.Sprites[p.State.CurAnim][CurAnimIdx]
	IsMonkey    bool
	Score       int
	Health      int
}

func (p *player) draw() {
	if p.State.Active {
		// loop to next idx of anim to draw
		animLen := len(p.Sprites[p.State.CurAnim])
		if p.LastAnimIdx == animLen-1 {
			p.LastAnimIdx = 0
		} else {
			p.LastAnimIdx++
		}
		p.submitUpdate()
		p.Sprites[p.State.CurAnim][p.LastAnimIdx].Draw(win, p.State.IdentityMatrix)
	}
}

func (p *player) collisionBox() pixel.Rect {
	// centre := cam.Unproject(win.Bounds().Center().Sub(playerSize))
	centre := playerVec

	fmt.Println("CCB:", centre)
	// fmt.Println("PSX:", centre.X+playerSize.X)
	// fmt.Println("PSY:", centre.Y+playerSize.Y)
	return pixel.R(
		playerVec.X+6,
		playerVec.Y+6,
		playerVec.X+42,
		playerVec.Y+22,
	)
}

func (p *player) submitUpdate() {
	update := gamestate.PlayerUpdate{ID: p.ID, State: p.State}
	payload, jsonErr := json.Marshal(update)
	// log.Printf("submitting  %+v", string(payload))

	if jsonErr != nil {
		log.Printf("ERROR: Failed to Marshal State : %s", jsonErr)
		return
	}
	mySIOClient.Emit("/updateobject", Message{Id: 0, Channel: "main", Text: string(payload)})

	// mySIOClient.Emit(topicString, string(payload))
}
