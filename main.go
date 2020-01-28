package main

import (
	"fmt"
	"os"

	client "github.com/mattmulhern/game-off-2019-scratch/client"
	game "github.com/mattmulhern/game-off-2019-scratch/game"
	server "github.com/mattmulhern/game-off-2019-scratch/server"
	zoo "github.com/mattmulhern/game-off-2019-scratch/zoo"
	dog "github.com/mattmulhern/game-off-2019-scratch/dog"
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
		client.RunTestClient()
	case "game":
		fmt.Println("game")
		game.Run()
	case "zoo":
		fmt.Println("zoo")
		zoo.Run()
	case "dog":
		fmt.Println("dog")
		dog.Run()
	default:
		fmt.Printf("bad opt! [%s]\n", argsWithoutProg[0])
		os.Exit(1)

	}
}
