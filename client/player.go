package client

import (
	"encoding/json"
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
	// centre := cam.Unproject(win.Bounds().Center().Sub(playerSize))
	centre := playerVec

	fmt.Println("CCB:", centre)
	// fmt.Println("PSX:", centre.X+playerSize.X)
	// fmt.Println("PSY:", centre.Y+playerSize.Y)
	return pixel.R(
		playerVec.X-6,
		playerVec.Y-6,
		playerVec.X+playerSize.X-6,
		playerVec.Y+playerSize.Y-6,
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
