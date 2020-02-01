package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

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
		time.Sleep(2 * time.Second)
		log.Println("Broadcasting State on request from ", channel.Channel)
		payload, encodingErr := json.Marshal(myState)
		if encodingErr != nil {
			log.Printf("ERROR: err encoding state: %s", encodingErr)
			return ""
		}
		// c.BroadcastTo("main", "/stateupdate", Message{10, "main", string(payload)})
		c.BroadcastTo("updatestate", "/updatestate", Message{99, "updatestate", string(payload)})

		return "SENT"
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

	serveMux.Handle("/socket.io/", server)

	log.Println("Starting server...")
	log.Panic(http.ListenAndServe(":3811", serveMux))
}

// //RunServer - Server entrypoint
// func RunServer() {

// 	server.OnEvent("/", "register", func(s socketio.Conn, playerName string) {
// 		freeMonkeyIdx, monkeyAvailable := getFreeMonkey()
// 		if myState.Player.Active == false {
// 			myState.Player.Active = true
// 			server.BroadcastToRoom("party", playerName+"-register", "PLAYER-REGISTERED:"+myState.Player.ID)
// 		} else if monkeyAvailable {
// 			myState.Monkeys[freeMonkeyIdx].Active = true
// 			server.BroadcastToRoom("party", playerName+"-register", "MONKEY-REGISTERED:"+strconv.Itoa(freeMonkeyIdx))
// 		} else {
// 			server.BroadcastToRoom("party", playerName+"-register", "MONKEY-ENGAGED-SIGNAL")
// 		}
// 		if broadCastErr := broadcastGameState(server); broadCastErr != nil {
// 			log.Printf("Failed to broadcast state: %s", broadCastErr)
// 		}

// 	})

// 	server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
// 		if broadCastErr := broadcastGameState(server); broadCastErr != nil {
// 			log.Printf("Failed to broadcast state: %s", broadCastErr)
// 		}
// 	})
// 	server.OnEvent("/", "PLAYER-UPDATE", func(s socketio.Conn, msg string) {
// 		playerUpdate := gamestate.PlayerUpdate{}

// 		jsonErr := json.Unmarshal([]byte(msg), &playerUpdate)

// 		if jsonErr != nil {
// 			log.Printf("ERROR: failed to unmarshal player update: %s", jsonErr)
// 			log.Printf("ERROR: player update msg: %s", msg)
// 			log.Printf("ERROR: player update: %+v", playerUpdate)
// 		}
// 		if playerUpdate.ID == "onhands" {
// 			myState.Player = *playerUpdate.State
// 		} else {
// 			//TODDO: sort it out for monkeys!
// 		}

// 		if broadCastErr := broadcastGameState(server); broadCastErr != nil {
// 			log.Printf("Failed to broadcast state: %s", broadCastErr)
// 		}
// 	})

// 	server.OnEvent("/", "bye", func(s socketio.Conn, playerName string) {
// 		if playerName == "player" {
// 			myState.Player.Active = false
// 		} else if strings.HasPrefix(playerName, "monkey") {
// 			monkeyIdx, converr := strconv.Atoi(strings.TrimPrefix(playerName, "monkey"))
// 			if converr != nil {
// 				log.Printf("ERROR: Could not get index for %s", playerName)
// 				return
// 			}
// 			myState.Monkeys[monkeyIdx].Active = false
// 		} else {
// 			log.Printf("ERROR: Cannot kill unknown player: %s", playerName)
// 			return
// 		}
// 		if broadCastErr := broadcastGameState(server); broadCastErr != nil {
// 			log.Printf("Failed to broadcast state: %s", broadCastErr)
// 		}
// 		return
// 	})

// 	server.OnError("/", func(e error) {
// 		fmt.Println("meet error:", e)
// 	})

// 	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
// 		fmt.Printf("%s closed: %s", s.ID(), reason)
// 	})
// 	go server.Serve()
// 	defer server.Close()

// 	http.Handle("/socket.io/", server)
// 	log.Println("Serving at localhost:8000...")

// 	log.Fatal(http.ListenAndServe(":8000", nil))
// }

// func broadcastGameState(server *socketio.Server) error {
// 	payload, encodingErr := json.Marshal(myState)
// 	if encodingErr != nil {
// 		return fmt.Errorf("Err encoding state: %s", encodingErr)
// 	}
// 	server.BroadcastToRoom("party", "update", string(payload))
// 	return nil
// }
