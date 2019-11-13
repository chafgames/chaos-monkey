package main

import (
	"fmt"
	"os"

	client "github.com/mattmulhern/game-off-2019-scratch/client"
	server "github.com/mattmulhern/game-off-2019-scratch/server"
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

	default:
		fmt.Printf("bad opt! [%s]\n", argsWithoutProg[0])
		os.Exit(1)

	}
}
