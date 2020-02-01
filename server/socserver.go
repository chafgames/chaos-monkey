package server

import (
	"log"
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

		c.Join("test")
		c.BroadcastTo("test", "/message", Message{10, "main", "some dick joined our channel!"})
	})
	server.On(gosocketio.OnDisconnection, func(c *gosocketio.Channel) {
		log.Println("Disconnected")
	})

	server.On("/join", func(c *gosocketio.Channel, channel Channel) string {
		time.Sleep(2 * time.Second)
		log.Println("Client joined to ", channel.Channel)
		return "joined to " + channel.Channel
	})

	return server
}
