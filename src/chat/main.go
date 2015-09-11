package main

import (
	"fmt"
	"os"

	"client"
	"server"
)

func main() {
	fmt.Println("Hello, Secret!")

	args := os.Args

	if args == nil || len(args) < 2 {
		fmt.Println("error")
		return
	}

	switch args[1] {
	case "client":
		startClient()
	case "server":
		startServer()
	default:
		fmt.Println("error")
	}
}

func startServer() {
	fmt.Println("server")
	server.CreateServer()
}

func startClient() {
	fmt.Println("client")
	client.CreateConn()
}
