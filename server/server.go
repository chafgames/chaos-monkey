package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	socketio "github.com/googollee/go-socket.io"
	zoogamestate "github.com/mattmulhern/game-off-2019-scratch/zoogamestate"
)

var myState *zoogamestate.GameState

//RunServer - Server entrypoint
func RunServer() {
	myState = &zoogamestate.GameState{
		ID: 0,
	}

	server, err := socketio.NewServer(nil)

	if err != nil {
		log.Fatal(err)
	}
	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("/")
		server.JoinRoom("party", s)

		fmt.Println("connected:", s.ID())
		newplayerID, _ := strconv.Atoi(s.ID()) //TODO errorchecking for atoi? check what socket.ID can be in general
		newPlayer := zoogamestate.Player{ID: newplayerID}
		myState.Players = append(myState.Players, &newPlayer)
		// server.BroadcastToRoom("party", "reply", ""+s.ID()+" joined!")

		return nil
	})
	server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		fmt.Println("notice:", msg)
		server.JoinRoom("party", s)
		myState.ID++
		payload, encodingErr := json.Marshal(myState)
		log.Printf("SENDING: %+v", payload)
		if encodingErr != nil {
			log.Printf("Err encoding state: %s", encodingErr)
			return
		}
		// payload := fmt.Sprintf("Total updates sent: %d", myState.ID)
		server.BroadcastToRoom("party", "update", string(payload))
	})

	server.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})
	server.OnError("/", func(e error) {
		fmt.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})
	go server.Serve()
	defer server.Close()

	http.Handle("/socket.io/", server)
	log.Println("Serving at localhost:8000...")

	log.Fatal(http.ListenAndServe(":8000", nil))
}
