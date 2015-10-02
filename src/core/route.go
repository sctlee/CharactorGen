package core

import (
	"fmt"
	"strings"
	// "secret/chatroom"
)

var routeList map[string]routeSet

type routeSet interface {
	Route(url string, c *Client)
}

func init() {
	routeList = make(map[string]routeSet)
	// routeList["chatroom"] = &chatroom.ChatRoom{}
}

func RegisterRuoter(key string, value routeSet) {
	routeList[key] = value
}

func MsgRoute(client *Client, msg string) bool {
	fmt.Printf("route %v msg:%s", client, msg)
	i := strings.Index(msg, " ")
	fmt.Println(i)
	if i != -1 {
		command := msg[:i]
		routeList[command].Route(msg[i:], client)
		return true
	}
	return false
}
