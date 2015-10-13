package main

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"example/chatroom"
	"example/db"
	"example/user"

	"github.com/sctlee/tcpx"
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
		db.StartPool()
		startServer()
	default:
		fmt.Println("error")
	}
}

func startServer() {
	fmt.Println("server")
	server := tcpx.CreateServer()
	// Register Router
	server.Router.RouteList["chatroom"] = chatroom.Route
	server.Router.RouteList["user"] = user.Route
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

	ic := &tcpx.TCPClient{
		Conn: c,
	}
	client := tcpx.CreateClient(ic)

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	go func() {
		for msg := range client.GetIncoming() {
			out.WriteString(msg + "\n")
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
