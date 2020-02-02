package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/chafgames/chaos-monkey/gamestate"
	gosocketio "github.com/graarh/golang-socketio"
)

var (
	myState *gamestate.GameState
)

//StartServer - entry point for module
func StartServer() {
	myState = gamestate.NewGameState()

	serveMux := http.NewServeMux()
	server := newSIOServer()

	server.On("/updatestate", func(c *gosocketio.Channel, channel Channel) string {
		payload, encodingErr := json.Marshal(myState)
		if encodingErr != nil {
			log.Printf("ERROR: err encoding state: %s", encodingErr)
			return ""
		}
		c.BroadcastTo("main", "/updatestate", Message{99, "main", string(payload)})

		return "SENT"
	})
	server.On("/updateobject", func(c *gosocketio.Channel, msg Message) string {
		var encodingErr error
		var playerUpdate gamestate.PlayerUpdate
		encodingErr = json.Unmarshal([]byte(msg.Text), &playerUpdate)
		if encodingErr != nil {
			log.Printf("ERROR: err decoding update: %s", msg.Text)
			return ""
		}

		if playerUpdate.ID == "onhands" {
			myState.Player = *playerUpdate.State
		} else if strings.HasPrefix(playerUpdate.ID, "monkey") {
			monkeyIDxStr := strings.TrimLeft(playerUpdate.ID, "monkey")
			monkeyIdx, serr := strconv.Atoi(monkeyIDxStr)
			if serr != nil {
				log.Printf("Could not get object to update from %s", playerUpdate.ID)
				return ""
			}
			myState.Monkeys[monkeyIdx] = *playerUpdate.State
		} else {
			log.Printf("Could not get object to update from %s", playerUpdate.ID)
			return ""
		}
		//TODO: read update payload here somehow!?!?!?
		payload, encodingErr := json.Marshal(myState)
		if encodingErr != nil {
			log.Printf("ERROR: err encoding state: %s", encodingErr)
			return ""
		}

		c.BroadcastTo("main", "/updatestate", Message{99, "main", string(payload)})

		return "OK"
	})

	server.On("/register", func(c *gosocketio.Channel, channel Channel) string {
		if myState.Player.Active == false {
			myState.Player.Active = true
			return "player"
		}
		freeMonkeyIdx, monkeyAvailable := getFreeMonkey()
		if monkeyAvailable {
			myState.Monkeys[freeMonkeyIdx].Active = true
			return fmt.Sprintf("monkey%d", freeMonkeyIdx)
		}
		return ""
	})
	server.On("/bye", func(c *gosocketio.Channel, msg Message) string {
		playerName := msg.Text
		if playerName == "onhands" {
			myState.Player.Active = false
			return "byenow"
		} else if strings.HasPrefix(playerName, "monkey") {
			monkeyIdx, converr := strconv.Atoi(strings.TrimPrefix(playerName, "monkey"))
			if converr != nil {
				log.Printf("ERROR: Could not get index for %s", playerName)
				return fmt.Sprintf("ERROR: Could not get index for %s", playerName)
			}
			myState.Monkeys[monkeyIdx].Active = false
			return "byenow"
		}
		return fmt.Sprintf("ERROR: Cannot kill unknown player: %s", playerName)
	})

	serveMux.Handle("/socket.io/", server)

	log.Println("Starting server...")
	log.Panic(http.ListenAndServe(":3811", serveMux))
}
