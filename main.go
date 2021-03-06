package main

import (
	"fmt"
	"os"

	client "github.com/chafgames/chaos-monkey/client" // our game client
	//  dog's TMX stuff
	server "github.com/chafgames/chaos-monkey/server" // our game server
)

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) == 0 {
		fmt.Printf("running client...\n")
		client.Run(argsWithoutProg)
	}
	switch argsWithoutProg[0] {
	case "server":
		fmt.Println("server")
		server.StartServer(argsWithoutProg)
	case "client":
		fmt.Println("client")
		client.Run(argsWithoutProg)
	default:
		client.Run(argsWithoutProg)
	}
}
