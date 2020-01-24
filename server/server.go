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

//RunServer - Server entrypoint
func RunServer() {
	players := make(map[string]*zoogamestate.Player)
	animals := make(map[string]*zoogamestate.Animal)
	cages := make(map[string]*zoogamestate.Cage)
	walls := make(map[string]*zoogamestate.Wall)
	myState = &zoogamestate.GameState{
		ID:      0,
		Players: players,
		Animals: animals,
		Cages:   cages,
		Walls:   walls,
	}
	server, err := socketio.NewServer(nil)

	if err != nil {
		log.Fatal(err)
	}
	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("/")
		server.JoinRoom("party", s)

		fmt.Println("connected:", s.ID())
		newplayerID := s.ID()
		newPlayer := zoogamestate.Player{ID: newplayerID, Active: true}
		// myState.Players = append(myState.Players, &newPlayer)
		myState.Players[newplayerID] = &newPlayer //TODO: check for existing id's and barf!
		return nil
	})
	server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		fmt.Println("notice:", msg)
		server.JoinRoom("party", s)
		// myState.ID++
		payload, encodingErr := json.Marshal(myState)
		log.Printf("SENDING: %+v", string(payload))
		if encodingErr != nil {
			log.Printf("Err encoding state: %s", encodingErr)
			return
		}
		// payload := fmt.Sprintf("Total updates sent: %d", myState.ID)
		server.BroadcastToRoom("party", "update", string(payload))
	})

	server.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)

		// myState.Players = removePlayer(myState.Players)
		if player, playerfound := myState.Players[s.ID()]; playerfound {
			log.Printf("bye: deleting player %+v", player)
			player.Active = false
			// delete(myState.Players, player.ID)
		}
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

func removePlayer(slice []*zoogamestate.Player, s int) []*zoogamestate.Player {
	return append(slice[:s], slice[s+1:]...)
}
