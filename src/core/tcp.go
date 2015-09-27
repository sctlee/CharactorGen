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
	self.listener, _ = net.Listen("tcp", fmt.Sprintf(":%s", port))
	defer self.listener.Close()
}

func (self *TCPServer) Accept() (tc *TCPClient) {
	if conn, err := self.listener.Accept(); err == nil {
		tc.Conn = conn
	}
	return
}

func (self *TCPClient) Read() (s string, err error) {
	reader := bufio.NewReader(self.Conn)
	for {
		if line, _, err := reader.ReadLine(); err == nil {
			s = string(line)
		} else {
			fmt.Printf("Read error: %s\n", err)
		}
		return
	}
}

func (self *TCPClient) Write(outgoing chan string) {
	writer := bufio.NewWriter(self.Conn)
	for data := range outgoing {
		writer.WriteString(data + "\n")
		// q: why flush is necessary? a:using buf mean: it won't send immedicately until buf is full
		writer.Flush()
	}
}
