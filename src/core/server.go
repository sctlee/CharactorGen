package core

import (
	"fmt"
	"log"
	"net"
)

const (
	MAXCLIENTS = 50
)

type ClientTable map[net.Conn]*Client

type Server struct {
	listener net.Listener
	clients  ClientTable
	pending  chan net.Conn
	quiting  chan net.Conn
	incoming chan string
	outgoing chan string
}

func CreateServer() (server *Server) {
	server = &Server{
		clients:  make(ClientTable),
		pending:  make(chan net.Conn),
		quiting:  make(chan net.Conn),
		incoming: make(chan string),
		outgoing: make(chan string),
	}
	return
}

func (server *Server) Listen() {
	go func() {
		for {
			select {
			case msg := <-server.incoming:
				log.Println(msg)
			case conn := <-server.pending:
				server.Join(conn)
				//
				// case conn := <-server.quiting:
			}
		}
	}()
}

func (server *Server) Join(conn net.Conn) {
	client := CreateClient(conn)
	server.clients[conn] = client

	client.Listen()
	logger.Printf("%v", conn)
	LogClose()

	go func() {
		for msg := range client.Incoming {
			// package msg whish conn
			// msg = fmt.Sprintf("format string", a ...interface{})

			server.incoming <- msg
		}
	}()
}

func (server *Server) Start(port string) {
	l, _ := net.Listen("tcp", fmt.Sprintf(":%s", port))
	server.listener = l
	defer server.listener.Close()

	// chan listen
	server.Listen()

	for {
		if conn, err := server.listener.Accept(); err == nil {
			server.pending <- conn
			// go func(c net.Conn) {
			// 	buf := make([]byte, 1024)
			// 	for {
			// 		cn, err := c.Read(buf)
			// 		if err != nil {
			// 			c.Close()
			// 			break
			// 		}
			// 		log.Println(cn, string(buf[:cn]))
			// 	}
			// }(conn)
		}
	}
}
