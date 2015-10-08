package route

import (
	"core/client"
	"fmt"
	"strings"
	// "secret/chatroom"
)

type RouteFun func(url string, c *client.Client)

type Router struct {
	RouteList map[string]RouteFun
}

func (self *Router) Route(client *client.Client, msg string) bool {
	fmt.Printf("route %v msg:%s", client, msg)
	i := strings.Index(msg, " ")
	fmt.Println(i)
	if i != -1 {
		command := msg[:i]
		fmt.Println(msg[i:])
		self.RouteList[command](msg[i:], client)
		return true
	}
	return false
}
