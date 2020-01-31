package main

import (
	"fmt"
	"os"

	client "github.com/chafgames/chaos-monkey/client" // our game client
	dog "github.com/chafgames/chaos-monkey/dog"       //  dog's TMX stuff
	server "github.com/chafgames/chaos-monkey/server" // our game server
)

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) == 0 {
		os.Exit(1)
	}
	switch argsWithoutProg[0] {
	case "server":
		fmt.Println("server")
		server.RunServer()
	case "client":
		fmt.Println("client")
		client.Run()
	case "dog":
		fmt.Println("dog")
		dog.Run()
	default:
		fmt.Printf("bad opt! [%s]\n", argsWithoutProg[0])
		os.Exit(1)
	}
}
