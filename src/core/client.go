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

type IClient interface {
	Read() (s string, err error)
	Write(outgoing chan string)
}

type Client struct {
	client IClient
	// Conn     net.Conn
	Group    int
	Xt       Xtime
	Incoming chan string
	Outgoing chan string
}

func CreateClient(ic IClient) (client *Client) {
	client = &Client{
		client: ic,
		// Conn:     conn,
		Incoming: make(chan string),
		Outgoing: make(chan string),
	}

	go client.client.Read()
	go client.client.Write(client.Outgoing)

	return
}

// func (client *Client) TRead() {
// 	reader := bufio.NewReader(client.Conn)
// 	for {
// 		if line, _, err := reader.ReadLine(); err == nil {
// 			client.Incoming <- string(line)
// 		} else {
// 			fmt.Printf("Read error: %s\n", err)
// 			return
// 		}
//
// 	}
// }
//
// func (client *Client) TWrite() {
// 	writer := bufio.NewWriter(client.Conn)
// 	for data := range client.Outgoing {
// 		writer.WriteString(data + "\n")
// 		// q: why flush is necessary? a:using buf mean: it won't send immedicately until buf is full
// 		writer.Flush()
// 	}
// }
