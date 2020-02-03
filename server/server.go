package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/chafgames/chaos-monkey/gamestate"
	gosocketio "github.com/graarh/golang-socketio"
)

var (
	myState    *gamestate.GameState
	myConnSids = make(map[string]string)
)

func printLives() {
	log.Printf(
		"P:%t/M0:%t/M1:%t/M2:%t/M3:%t/M4:%t/M5:%t/M6:%t/M17%t",
		myState.Player.Active,
		myState.Monkeys[0].Active,
		myState.Monkeys[1].Active,
		myState.Monkeys[2].Active,
		myState.Monkeys[3].Active,
		myState.Monkeys[4].Active,
		myState.Monkeys[5].Active,
		myState.Monkeys[6].Active,
		myState.Monkeys[7].Active,
	)
}

func broadcastState(c *gosocketio.Channel) bool {
	payload, encodingErr := json.Marshal(myState)
	if encodingErr != nil {
		log.Printf("ERROR: err encoding state: %s", encodingErr)
		return false
	}
	c.BroadcastTo("main", "/updatestate", Message{99, "main", string(payload)})
	return true
}

var mySIOAddr string

//StartServer - entry point for module
func StartServer(args []string) {
	if len(args) < 2 {
		log.Printf("ERROR: expected at least two args")
		log.Printf("e.g. ./chaos-monkey server 127.0.0.1:3811")
		os.Exit(1)
	}
	var addrReg = regexp.MustCompile(`^[0-9\.]*:[0-9]*$`)
	if addrReg.Match([]byte(args[1])) == false {
		log.Printf("Error: %s not of expected form e.g. 127.0.0.1:3811", args[1])
		os.Exit(1)
	}

	mySIOAddr = args[1]

	myState = gamestate.NewGameState()

	serveMux := http.NewServeMux()
	server := newSIOServer()

	server.On("/updatestate", func(c *gosocketio.Channel, channel Channel) string {
		if ok := broadcastState(c); ok != true {
			log.Printf("failed to braodcast state")
			return "NOTSENT"
		}
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

		return "OK"
	})

	server.On("/register", func(c *gosocketio.Channel, channel Channel) string {
		if myState.Player.Active == false {
			myState.Player.Active = true
			myConnSids[c.Id()] = "player"
			log.Printf("assigning player to %s", c.Id())
			printLives()
			if ok := broadcastState(c); ok != true {
				log.Printf("failed to braodcast state")
			}
			return "player"

		}
		freeMonkeyIdx, monkeyAvailable := getFreeMonkey()
		if monkeyAvailable {
			myState.Monkeys[freeMonkeyIdx].Active = true
			myConnSids[c.Id()] = fmt.Sprintf("monkey%d", freeMonkeyIdx)
			log.Printf("assigning monkey%d to %s", freeMonkeyIdx, c.Id())
			printLives()
			if ok := broadcastState(c); ok != true {
				log.Printf("failed to braodcast state")
			}
			return fmt.Sprintf("monkey%d", freeMonkeyIdx)
		}
		return ""
	})
	server.On("/bye", func(c *gosocketio.Channel, msg Message) string {
		playerName := msg.Text
		if playerName == "onhands" {
			myState.Player.Active = false
			if ok := broadcastState(c); ok != true {
				log.Printf("failed to braodcast state")
			}
			return "byenow"
		} else if strings.HasPrefix(playerName, "monkey") {
			monkeyIdx, converr := strconv.Atoi(strings.TrimPrefix(playerName, "monkey"))
			if converr != nil {
				log.Printf("ERROR: Could not get index for %s", playerName)
				return fmt.Sprintf("ERROR: Could not get index for %s", playerName)
			}
			myState.Monkeys[monkeyIdx].Active = false
			if ok := broadcastState(c); ok != true {
				log.Printf("failed to braodcast state")
			}
			return "byenow"
		}
		return fmt.Sprintf("ERROR: Cannot kill unknown player: %s", playerName)
	})

	serveMux.Handle("/socket.io/", server)

	log.Println("Starting server...")
	log.Panic(http.ListenAndServe(mySIOAddr, serveMux))
}
