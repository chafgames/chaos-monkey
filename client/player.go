package client

import (
	"encoding/json"
	"fmt"
	"log"
	"path/filepath"

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

func (p *player) loadPlayerSheet() {
	playerSheet, err := loadPicture(filepath.Join(binPath, "assets/monkey.png"))
	if err != nil {
		panic(err)
	}

	playerPics = []*pixel.Sprite{
		pixel.NewSprite(playerSheet, spritePos(0, 0)),
	}
}

func (p *player) draw() {
	if p.State.Active {
		p.submitUpdate()
		p.Sprites[0].Draw(win, p.State.IdentityMatrix)
	}
}

func (p *player) collisionBox() pixel.Rect {
	centre := cam.Unproject(win.Bounds().Center().Sub(playerSize))

	return pixel.R(
		centre.X,
		centre.Y,
		centre.X+playerSize.X,
		centre.Y+playerSize.Y,
	)
}

func (p *player) submitUpdate() {
	topicString := fmt.Sprintf("PLAYER-UPDATE")
	update := gamestate.PlayerUpdate{ID: p.ID, State: p.State}
	payload, jsonErr := json.Marshal(update)
	// log.Printf("submitting  %+v", string(payload))

	if jsonErr != nil {
		log.Printf("ERROR: Failed to Marshal State : %s", jsonErr)
		return
	}
	client.SocketioClient.Emit(topicString, string(payload))
}
