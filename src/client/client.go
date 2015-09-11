package client

import (
	"bufio"
	"log"
	"net"
	"os"
)

type Client struct {
	conn net.Conn
}

func CreateConn() {
	conn, err := net.Dial("tcp", ":9000")
	message := make(chan string)
	in := bufio.NewReader(os.Stdin)

	defer conn.Close()

	if err == nil {
		go func(c net.Conn, m chan string) {
			for data := range m {
				cn, err := c.Write([]byte(data))
				log.Println(cn, err)
			}
		}(conn, message)
	} else {
		log.Println(err)
	}
	for {
		line, _, _ := in.ReadLine()
		message <- string(line)
	}
}
