package server

import (
	"log"
	"net"
	"os"

	. "client"
)

const (
	MAXCLIENTS = 50
)

var logger *log.Logger

type Message chan string
type Token chan int
type ClientTable map[net.Conn]*Client

type Server struct {
	listener net.Listener
	clients  ClientTable
	tokens   Token
	pending  chan net.Conn
	quiting  chan net.Conn
	incoming Message
	outgoing Message
}

func init() {
	var file *os.File
	filename := "gen.log"

	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		file, _ = os.Create(filename)
	} else {
		file, _ = os.OpenFile(filename, os.O_APPEND, 0666)
	}

	logger = log.New(file, "logger: ", log.Lshortfile)
}

func CreateServer() {
	l, _ := net.Listen("tcp", ":9000")
	defer l.Close()
	for {
		log.Println("haha")
		if conn, err := l.Accept(); err == nil {
			log.Println("hehe")
			go func(c net.Conn) {
				buf := make([]byte, 1024)
				for {
					cn, err := c.Read(buf)
					if err == nil {
						log.Println(cn, string(buf[:cn]))
					} else {
						c.Close()
						break
					}
				}
			}(conn)
		}
	}
}
