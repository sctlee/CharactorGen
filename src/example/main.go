package main

import (
	"bufio"
	"fmt"
	"net"
	"os"

	// "features/auth"
	"features/chatroom"

	"github.com/sctlee/hazel"
	"github.com/sctlee/hazel/auth"
	"github.com/sctlee/hazel/daemon/service"
	"github.com/sctlee/hazel/tcpx/client"

	// "github.com/sctlee/hazel/auth"
	// "github.com/sctlee/hazel/db"
)

func main() {
	fmt.Println("Hello, Secret!")

	args := os.Args

	switch args[1] {
	case "client":
		startClient(args[2])
	case "server":
		var cf *hazel.Config
		if args == nil || len(args) < 2 {
			fmt.Println("error")
			return
		}
		cf = hazel.LoadConfig()
		startServer(cf)
	default:
		fmt.Println("error")
	}
}

func startServer(cf *hazel.Config) {
	fmt.Println("server listen port:", cf.Port)
	// Register Router
	hazel.MainDaemon(cf,
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
