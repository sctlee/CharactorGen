package core

import (
	"bufio"
	"fmt"
	"net"
)

type TCPServer struct {
	listener net.Listener
}

type TCPClient struct {
	Conn net.Conn
}

func (self *TCPServer) Listen(port string) {
	fmt.Println("1")

	self.listener, _ = net.Listen("tcp", fmt.Sprintf(":%s", port))
	// defer self.listener.Close()
}

func (self *TCPServer) Accept() (tc *TCPClient) {
	fmt.Println("2")
	if conn, err := self.listener.Accept(); err == nil {
		fmt.Printf("%v", conn)
		tc = &TCPClient{
			Conn: conn,
		}
	} else {
		fmt.Println(err)
	}
	return
}

func (self *TCPClient) TRead(incoming chan string) {
	reader := bufio.NewReader(self.Conn)
	for {
		fmt.Println("r")
		if line, _, err := reader.ReadLine(); err == nil {
			fmt.Println("9")
			incoming <- string(line)
		} else {
			fmt.Printf("Read error: %s\n", err)
		}
	}
}

func (self *TCPClient) TWrite(outgoing chan string) {
	writer := bufio.NewWriter(self.Conn)
	for data := range outgoing {
		fmt.Println("8")
		writer.WriteString(data + "\n")
		// q: why flush is necessary? a:using buf mean: it won't send immedicately until buf is full
		writer.Flush()
	}
}
