package main

import (
	"bufio"
	"core"
	"example/chatroom"
	"fmt"
	"net"
	"os"
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
	server := core.CreateServer()
	// Register Router
	core.RegisterRouter("chatroom", chatroom.Route)
	// End Register
	server.Start("9000")
}

func startClient() {
	fmt.Println("client")
	c, err := net.Dial("tcp", ":9000")
	if err != nil {
		fmt.Println("hahah")
		return
	}

	ic := &core.TCPClient{
		Conn: c,
	}
	client := core.CreateClient(ic)

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	go func() {
		for {
			out.WriteString(client.GetIncoming() + "\n")
			out.Flush()
		}
	}()
	// go func(c net.Conn, m chan string) {
	// 	for data := range m {
	// 		cn, err := c.Write([]byte(data))
	// 		log.Println(cn, err)
	// 	}
	// }(client.Conn, message)

	for {
		line, _, _ := in.ReadLine()
		client.PutOutgoing(string(line))
	}
}
