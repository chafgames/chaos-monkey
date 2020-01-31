package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
	zoogamestate "github.com/mattmulhern/game-off-2019-scratch/zoogamestate"
)

var myState *zoogamestate.GameState

// func myConnChecker(r *http.Request) (http.Header, error) {
//  return http.Header{"zoo-status": []string{"accepted"}}, fmt.Errorf("ID: %s already in use", "hello")
// 	myHeader := http.Header{"zoo-status": []string{"accepted"}}
// 	return myHeader, nil
// }

//RunServer - Server entrypoint
func RunServer() {
	myState = zoogamestate.NewGameState()
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

		server.BroadcastToRoom("party", "reply", playerName+" registered!")
		server.BroadcastToRoom("party", playerName, "psst... you're called "+playerName+" !")
		newPlayerObject := zoogamestate.NewObjectState(playerName)
		myState.Players[playerName] = newPlayerObject
		myState.Players[playerName].Active = true
		if broadCastErr := broadcastGameState(server); broadCastErr != nil {
			log.Printf("Failed to broadcast state: %s", broadCastErr)
		}
	})

	server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		if broadCastErr := broadcastGameState(server); broadCastErr != nil {
			log.Printf("Failed to broadcast state: %s", broadCastErr)
		}
	})

	server.OnEvent("/", "bye", func(s socketio.Conn, playerName string) {
		if player, playerfound := myState.Players[playerName]; playerfound {
			log.Printf("bye: killing player %+v", player)
			player.Active = false
			// s.Close()

			if broadCastErr := broadcastGameState(server); broadCastErr != nil {
				log.Printf("Failed to broadcast state: %s", broadCastErr)
			}
			return
		}
		log.Printf("bye: Could not find player %s", playerName)
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

func broadcastGameState(server *socketio.Server) error {
	payload, encodingErr := json.Marshal(myState)
	if encodingErr != nil {
		return fmt.Errorf("Err encoding state: %s", encodingErr)
	}
	server.BroadcastToRoom("party", "update", string(payload))
	return nil
}
