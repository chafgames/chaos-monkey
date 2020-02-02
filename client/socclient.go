package client

import (
	"log"
	"runtime"
	"time"

	gosocketio "github.com/graarh/golang-socketio"
	transport "github.com/graarh/golang-socketio/transport"
)

type Channel struct {
	Channel string `json:"channel"`
}

type Message struct {
	Id      int    `json:"id"`
	Channel string `json:"channel"`
	Text    string `json:"text"`
}

func sendJoin(c *gosocketio.Client) {
	log.Println("Acking /join")
	result, err := c.Ack("/join", Channel{"main"}, time.Second*5)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Ack result to /join: ", result)
	}
}
func sendRegister(c *gosocketio.Client) (string, bool) {
	log.Println("Acking /register")
	result, err := c.Ack("/register", Channel{"main"}, time.Second*5)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Ack result to /register: ", result)
		return result, true
	}
	return "", false
}
func sendBye(c *gosocketio.Client, playerID string) (string, bool) {
	log.Println("Acking /bye")
	result, err := c.Ack("/bye", Message{Id: 0, Channel: "main", Text: playerID}, time.Second*5)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Ack result to /bye: ", result)
		return result, true
	}
	return "", false
}

func sendUpdateRequest(c *gosocketio.Client) (string, bool) {
	c.Emit("/updatestate", Message{Id: 0, Channel: "main", Text: "come on!"})
	return "", false
}

func newSIOClient() (*gosocketio.Client, error) {
	runtime.GOMAXPROCS(runtime.NumCPU())

	c, err := gosocketio.Dial(
		gosocketio.GetUrl("localhost", 3811, false),
		transport.GetDefaultWebsocketTransport())
	if err != nil {
		log.Fatal(err)
	}

	err = c.On("/message", func(h *gosocketio.Channel, args Message) {
		log.Println("/message: ", args)
	})
	if err != nil {
		log.Fatal(err)
	}

	err = c.On(gosocketio.OnDisconnection, func(h *gosocketio.Channel) {
		log.Fatal("Disconnected")
	})
	if err != nil {
		log.Fatal(err)
	}

	err = c.On(gosocketio.OnConnection, func(h *gosocketio.Channel) {
		log.Println("Connected")
	})
	if err != nil {
		log.Fatal(err)
	}

	return c, nil
}
