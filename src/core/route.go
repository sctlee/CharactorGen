package core

import (
	"fmt"
)

func MsgRoute(client *Client, msg string) bool {
	fmt.Printf("route %v msg:%s", client, msg)
	client.outgoing <- msg
	return true
}
