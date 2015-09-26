package core

import (
	"bufio"
	"fmt"
	"net"
)

type Xtime struct {
	isExist  bool
	question string
}

type Client struct {
	Conn     net.Conn
	Group    int
	Xt       Xtime
	Incoming chan string
	Outgoing chan string
}

func CreateClient(conn net.Conn) (client *Client) {
	client = &Client{
		Conn:     conn,
		Incoming: make(chan string),
		Outgoing: make(chan string),
	}
	return
}

func (client *Client) Listen() {
	go client.Read()
	go client.Write()
	// go func(){
	// 	for {
	// 		select {
	// 		case receive := <- client
	// 			log.Println()
	// 		}
	// 	}
	// }()
}

func (client *Client) Read() {
	reader := bufio.NewReader(client.Conn)
	for {
		if line, _, err := reader.ReadLine(); err == nil {
			client.Incoming <- string(line)
		} else {
			fmt.Printf("Read error: %s\n", err)
			return
		}

	}
}

func (client *Client) Write() {
	writer := bufio.NewWriter(client.Conn)
	for data := range client.Outgoing {
		writer.WriteString(data + "\n")
		// q: why flush is necessary? a:using buf mean: it won't send immedicately until buf is full
		writer.Flush()
	}
}
