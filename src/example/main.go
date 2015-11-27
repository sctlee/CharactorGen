package main

import (
	"bufio"
	"fmt"
	"net"
	"os"

	// "features/auth"
	"features/chatroom"

	"github.com/sctlee/tcpx"
	"github.com/sctlee/tcpx/auth"
	"github.com/sctlee/tcpx/daemon/service"
	"github.com/sctlee/tcpx/tcpx/client"

	// "github.com/sctlee/tcpx/auth"
	// "github.com/sctlee/tcpx/db"
)

func main() {
	fmt.Println("Hello, Secret!")

	var cf *tcpx.Config
	args := os.Args

	if args == nil || len(args) < 2 {
		fmt.Println("error")
		return
	}
	cf = tcpx.LoadConfig()

	switch args[1] {
	case "client":
		startClient(":" + cf.Port)
	case "server":
		startServer(cf)
	default:
		fmt.Println("error")
	}
}

func startServer(cf *tcpx.Config) {
	fmt.Println("server")
	// Register Router
	tcpx.MainDaemon(cf,
		service.NewService("chatroom", chatroom.NewChatroomAction()),
		service.NewService("auth", auth.NewAuthAction()))
	// server.Routers["auth"] = auth.Router
	// End Register
}

func startClient(ip string) {
	fmt.Println("client")
	conn, err := net.Dial("tcp", ip)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ip)

	client := client.CreateClient(conn, "")

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	go func() {
		for {
			msg, ok := client.GetIncoming()
			if !ok {
				break
			}
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
