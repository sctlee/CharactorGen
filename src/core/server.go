package core

import (
	"log"
	"net"
)

const (
	MAXCLIENTS = 50
)

type IServer interface {
	Listen(port string)
	Accept() (tc *TCPClient)
}

type ClientTable map[IClient]*Client

type Server struct {
	server   IServer
	listener net.Listener
	clients  ClientTable
	pending  chan IClient
	// quiting  chan net.Conn
	incoming chan string
	outgoing chan string
}

func CreateServer() (server *Server) {
	server = &Server{
		server:  &TCPServer{},
		clients: make(ClientTable),
		pending: make(chan IClient),
		// quiting:  make(chan net.Conn),
		incoming: make(chan string),
		outgoing: make(chan string),
	}
	return
}

func (self *Server) Listen() {
	go func() {
		for {
			select {
			case msg := <-self.incoming:
				log.Println(msg)
			case conn := <-self.pending:
				self.Join(conn)
				//
				// case conn := <-server.quiting:
			}
		}
	}()
}

func (self *Server) Join(ic IClient) {
	client := CreateClient(ic)
	self.clients[ic] = client

	LogClose()

	go func() {
		for msg := range client.Incoming {
			// package msg whish conn
			// msg = fmt.Sprintf("format string", a ...interface{})
			if !MsgRoute(client, msg) {
				self.incoming <- msg
			}
		}
	}()
}

func (self *Server) Start(port string) {

	self.server.Listen(port)
	// l, _ := net.Listen("tcp", fmt.Sprintf(":%s", port))
	// self.listener = l
	// defer self.listener.Close()

	// chan listen
	self.Listen()

	for {
		self.pending <- self.server.Accept()
		// if conn, err := self.listener.Accept(); err == nil {
		// 	self.pending <- conn
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
		// }
	}
}
