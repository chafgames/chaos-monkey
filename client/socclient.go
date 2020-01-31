package client

import (
	socketio_client "github.com/zhouhui8915/go-socket.io-client"

	"fmt"
)

//SocClient - wrapper for socketio connection
type socClient struct {
	SocketioClient *socketio_client.Client
	ID             string
}

//NewClient - connect to game server
func newSocClient() (*socClient, error) {

	opts := &socketio_client.Options{
		Transport: "websocket",
		Query:     make(map[string]string),
	}
	opts.Query["user"] = "user"
	opts.Query["pwd"] = "pass"

	uri := "http://127.0.0.1:8000"

	client, err := socketio_client.NewClient(uri, opts)
	if err != nil {
		return nil, fmt.Errorf("socketio_client returned error: %v", err)
	}

	myClient := socClient{SocketioClient: client}
	return &myClient, nil
}

// Notice - send msg tto server on topic 'notice'
func (c *socClient) Notice(msg string) {
	c.SocketioClient.Emit("notice", msg)
}

// Bye - tell the server you're leaving!
func (c *socClient) Bye(msg string) {
	c.SocketioClient.Emit("bye", msg)
}
