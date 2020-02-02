package server

import (
	"log"
	"strconv"
	"strings"
	"time"

	gosocketio "github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
)

// Channel - //TODO
type Channel struct {
	Channel string `json:"channel"`
}

//Message - //TODO
type Message struct {
	ID      int    `json:"id"`
	Channel string `json:"channel"`
	Text    string `json:"text"`
}

func newSIOServer() *gosocketio.Server {
	server := gosocketio.NewServer(transport.GetDefaultWebsocketTransport())

	server.On(gosocketio.OnConnection, func(c *gosocketio.Channel) {
		log.Println("Connected")

		c.Emit("/message", Message{10, "main", "Let the chaos reign!"})

		c.Join("main")
	})
	server.On(gosocketio.OnDisconnection, func(c *gosocketio.Channel) {
		deadRole, foundDeadRole := myConnSids[c.Id()]
		if foundDeadRole == true {
			if deadRole == "player" {
				myState.Player.Active = false
				log.Printf("%s disconnected, killing off %s", c.Id(), deadRole)
				printLives()
				if ok := broadcastState(c); ok != true {
					log.Printf("failed to braodcast state")
				}
				return
			} else if strings.HasPrefix(deadRole, "monkey") {
				monkeyIDxStr := strings.TrimLeft(deadRole, "monkey")
				monkeyIdx, serr := strconv.Atoi(monkeyIDxStr)
				if serr != nil {
					log.Printf("Could not figure out who to kill from %s", deadRole)
				}
				myState.Monkeys[monkeyIdx].Active = false
				log.Printf("%s disconnected, killed monkey at index %d", c.Id(), monkeyIdx)
				printLives()
				if ok := broadcastState(c); ok != true {
					log.Printf("failed to braodcast state")
				}
				return

			} else {
				log.Printf("Could not figure out who to kill from %s", deadRole)
				return
			}
		}
		return
	})

	server.On("/join", func(c *gosocketio.Channel, channel Channel) string {
		time.Sleep(2 * time.Second)
		log.Println("Client joined to ", channel.Channel)
		return "joined to " + channel.Channel
	})

	return server
}
