package client

import (
	"github.com/zhouhui8915/go-socket.io-client"

	"log"
	"sync"
	"time"
)

// RunTestClient - client entrypoint
func RunTestClient() {

	opts := &socketio_client.Options{
		Transport: "websocket",
		Query:     make(map[string]string),
	}
	opts.Query["user"] = "user"
	opts.Query["pwd"] = "pass"
	uri := "http://127.0.0.1:8000"

	client, err := socketio_client.NewClient(uri, opts)
	if err != nil {
		log.Printf("NewClient error:%v\n", err)
		return
	}

	client.On("error", func() {
		log.Printf("on error\n")
	})
	client.On("connection", func() {
		log.Printf("on connect\n")
	})
	client.On("message", func(msg string) {
		log.Printf("on message:%v\n", msg)
	})
	client.On("notice", func(msg string) {
		log.Printf("on notice:%v\n", msg)
	})
	client.On("update", func(msg string) {
		log.Printf("update:%v\n", msg)
	})
	client.On("reply", func(msg string) {
		log.Printf("on reply:%v\n", msg)
	})
	client.On("disconnection", func() {
		log.Printf("on disconnect\n")
	})

	var waitgroup sync.WaitGroup
	waitgroup.Add(1)
	go sendGenericNotice(client)
	waitgroup.Wait()
	log.Println("done")
}

func sendGenericNotice(s *socketio_client.Client) {
	for {
		log.Printf("sending")
		s.Emit("notice", "hi there")
		time.Sleep(3 * time.Millisecond)
	}
}
