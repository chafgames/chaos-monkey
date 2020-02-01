package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	gamestate "github.com/chafgames/chaos-monkey/gamestate"
	socketio "github.com/googollee/go-socket.io"
)

var myState *gamestate.GameState

//RunServer - Server entrypoint
func RunServer() {
	myState = gamestate.NewGameState()
	server, err := socketio.NewServer(nil)

	if err != nil {
		log.Fatal(err)
	}
	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("/")
		server.JoinRoom("party", s)

		fmt.Println("connected:", s.ID())
		return nil
	})

	server.OnEvent("/", "register", func(s socketio.Conn, playerName string) {
		freeMonkeyIdx, monkeyAvailable := getFreeMonkey()
		if myState.Player.Active == false {
			myState.Player.Active = true
			server.BroadcastToRoom("party", playerName+"-register", "PLAYER-REGISTERED:"+myState.Player.ID)
		} else if monkeyAvailable {
			myState.Monkeys[freeMonkeyIdx].Active = true
			server.BroadcastToRoom("party", playerName+"-register", "MONKEY-REGISTERED:"+strconv.Itoa(freeMonkeyIdx))
		} else {
			server.BroadcastToRoom("party", playerName+"-register", "MONKEY-ENGAGED-SIGNAL")
		}
		if broadCastErr := broadcastGameState(server); broadCastErr != nil {
			log.Printf("Failed to broadcast state: %s", broadCastErr)
		}

	})

	server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		if broadCastErr := broadcastGameState(server); broadCastErr != nil {
			log.Printf("Failed to broadcast state: %s", broadCastErr)
		}
	})
	server.OnEvent("/", "PLAYER-UPDATE", func(s socketio.Conn, msg string) {
		playerUpdate := gamestate.PlayerUpdate{}

		jsonErr := json.Unmarshal([]byte(msg), &playerUpdate)

		if jsonErr != nil {
			log.Printf("ERROR: failed to unmarshal player update: %s", jsonErr)
			log.Printf("ERROR: player update msg: %s", msg)
			log.Printf("ERROR: player update: %+v", playerUpdate)
		}
		if playerUpdate.ID == "onhands" {
			myState.Player = *playerUpdate.State
		} else {
			//TODDO: sort it out for monkeys!
		}

		if broadCastErr := broadcastGameState(server); broadCastErr != nil {
			log.Printf("Failed to broadcast state: %s", broadCastErr)
		}
	})

	server.OnEvent("/", "bye", func(s socketio.Conn, playerName string) {
		if playerName == "player" {
			myState.Player.Active = false
		} else if strings.HasPrefix(playerName, "monkey") {
			monkeyIdx, converr := strconv.Atoi(strings.TrimPrefix(playerName, "monkey"))
			if converr != nil {
				log.Printf("ERROR: Could not get index for %s", playerName)
				return
			}
			myState.Monkeys[monkeyIdx].Active = false
		} else {
			log.Printf("ERROR: Cannot kill unknown player: %s", playerName)
			return
		}
		if broadCastErr := broadcastGameState(server); broadCastErr != nil {
			log.Printf("Failed to broadcast state: %s", broadCastErr)
		}
		return
	})

	server.OnError("/", func(e error) {
		fmt.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Printf("%s closed: %s", s.ID(), reason)
	})
	go server.Serve()
	defer server.Close()

	http.Handle("/socket.io/", server)
	log.Println("Serving at localhost:8000...")

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func getFreeMonkey() (int, bool) {
	for index, monkeyState := range myState.Monkeys {
		if monkeyState.Active == false {
			return index, true
		}
	}
	return -1, false
}

func broadcastGameState(server *socketio.Server) error {
	payload, encodingErr := json.Marshal(myState)
	if encodingErr != nil {
		return fmt.Errorf("Err encoding state: %s", encodingErr)
	}
	server.BroadcastToRoom("party", "update", string(payload))
	return nil
}
