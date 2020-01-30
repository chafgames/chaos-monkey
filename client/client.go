package client

import (
	socketio_client "github.com/zhouhui8915/go-socket.io-client"

	"fmt"
	"log"
	"sync"
	"time"
)

//Client - wrapper for socketio connection
type Client struct {
	SocketioClient *socketio_client.Client
	ID             string
}

//NewClient - connect to game server
func NewClient(clientID string) (*Client, error) {

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

	// client.On("error", func() {
	// 	log.Printf("on error\n")
	// })
	// client.On("connection", func() {
	// 	log.Printf("on connect\n")
	// })
	// client.On("message", func(msg string) {
	// 	log.Printf("on message:%v\n", msg)
	// })
	// client.On("notice", func(msg string) {
	// 	log.Printf("on notice:%v\n", msg)
	// })
	// // client.On("update", func(msg string) {
	// // 	log.Printf("update:%v\n", msg)
	// // })
	// client.On("reply", func(msg string) {
	// 	log.Printf("on reply:%v\n", msg)
	// })
	// client.On("disconnection", func() {
	// 	log.Println("disconnected from server")
	// })

	myClient := Client{ID: clientID, SocketioClient: client}
	return &myClient, nil
}

// RunTestClient - client entrypoint
func RunTestClient() {
	myClient, _ := NewClient("testclient")
	var waitgroup sync.WaitGroup
	waitgroup.Add(1)
	go myClient.sendGenericNoticeLoop()
	waitgroup.Wait()
	log.Println("done")
}

func (c *Client) sendGenericNoticeLoop() {
	for {
		log.Printf("sending")
		c.SocketioClient.Emit("notice", "hi there")
		time.Sleep(3 * time.Millisecond)
	}
}

//NOTICE -- //TODO
func (c *Client) Notice(msg string) {
	c.SocketioClient.Emit("notice", msg)
}

// Bye - //TODO
func (c *Client) Bye(msg string) {
	c.SocketioClient.Emit("bye", msg)
}
