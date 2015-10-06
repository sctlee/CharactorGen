package route

import (
	"core/client"
	"fmt"
	"strings"
	// "secret/chatroom"
)

type route func(url string, c *client.Client)

var routeList map[string]route

func init() {
	routeList = make(map[string]route)
}

func RegisterRouter(key string, value route) {
	routeList[key] = value
}

func MsgRoute(client *client.Client, msg string) bool {
	fmt.Printf("route %v msg:%s", client, msg)
	i := strings.Index(msg, " ")
	fmt.Println(i)
	if i != -1 {
		command := msg[:i]
		fmt.Println(msg[i:])
		routeList[command](msg[i:], client)
		return true
	}
	return false
}
