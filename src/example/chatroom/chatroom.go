package chatroom

import (
	"core/client"
	"core/utils"
	"example/user"
	"fmt"
	"strings"
)

var CHATROMMS = []string{"1", "2", "3", "4", "5"}

type ChatRoom struct {
	clients []*client.Client
}

var ChatRoomList map[string]*ChatRoom

var userChatList map[*client.Client]string

func init() {
	ChatRoomList = make(map[string]*ChatRoom, 10)
	userChatList = make(map[*client.Client]string)
	for _, k := range CHATROMMS {
		ChatRoomList[k] = &ChatRoom{}
	}
}

func List(client *client.Client) {
	client.PutOutgoing(fmt.Sprintf("You can choose one chatroom to join:\n%s",
		strings.Join(CHATROMMS, "\t")))
}

func Join(client *client.Client, paramString string) {
	params := strings.Fields(paramString)
	if len(params) != 1 {
		client.PutOutgoing("You can only input one param")
		return
	}
	ctName := params[0]
	if utils.StringInSlice(ctName, CHATROMMS) {
		userChatList[client] = ctName
		ChatRoomList[ctName].clients = append(ChatRoomList[ctName].clients, client)
		client.PutOutgoing(fmt.Sprintf("you have joined <%s> chatroom", ctName))
	}
}

func SendMsg(client *client.Client, paramString string) {
	msg := paramString
	if ctName, ok := userChatList[client]; ok {
		for _, c := range ChatRoomList[ctName].clients {
			c.PutOutgoing(fmt.Sprintf("%s says: %s",
				user.GetUserName(client),
				msg),
			)
		}
	}
}
