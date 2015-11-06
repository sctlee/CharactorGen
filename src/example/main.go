package main

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"features/chatroom"
	"features/growtree"

	"github.com/sctlee/tcpx"
	"github.com/sctlee/tcpx/db"
)

func main() {
	fmt.Println("Hello, Secret!")

	var cf tcpx.Config
	args := os.Args

	if args == nil || len(args) < 2 {
		fmt.Println("error")
		return
	}
	if len(args) == 2 {
		cf = tcpx.LoadConfig("config.yml")
	} else if len(args) == 3 {
		cf = tcpx.LoadConfig(args[2])
	}
	fmt.Println(cf)

	switch args[1] {
	case "client":
		startClient(":" + cf.Port)
	case "server":
		db.StartPool(cf.Db)
		startServer(cf.Port)
	default:
		fmt.Println("error")
	}
}

func startServer(port string) {
	fmt.Println("server")
	server := tcpx.CreateServer()
	// Register Router
	server.Routers["chatroom"] = chatroom.Route
	server.Routers["growtree"] = growtree.Route
	// End Register
	server.Start("9000")
}

func startClient(ip string) {
	fmt.Println("client")
	fmt.Println(ip)
	c, err := net.Dial("tcp", ip)
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
