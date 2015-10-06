package chatroom

import (
	"core/client"
	"example/user"
	"fmt"
	"strings"
)

type ChatRoom struct {
	clients []*client.Client
}

var ChatRoomList map[string]*ChatRoom

func init() {
	ChatRoomList = make(map[string]*ChatRoom, 10)
	ChatRoomList["all"] = &ChatRoom{}
}

func Route(url string, client *client.Client) {
	url = strings.TrimSpace(url)
	i := strings.Index(url, " ")
	action := url[:i]
	switch action {
	case "join":
		ChatRoomList["all"].clients = append(ChatRoomList["all"].clients, client)
		client.PutOutgoing("you have joined <all> chatroom")
	case "send":
		for _, c := range ChatRoomList["all"].clients {
			c.PutOutgoing(fmt.Sprintf("%s says: %s", user.GetUserName(client), strings.TrimSpace(url[i:])))
		}
	}
}
