package chatroom

import (
	"fmt"
	"log"
	"strings"

	"example/user"

	"github.com/sctlee/tcpx"
	"github.com/sctlee/utils"
)

var CHATROMMS = []string{}

type Chatroom struct {
	ct      *Chatrooms
	clients []*tcpx.Client
}

var ChatroomList map[string]*Chatroom

var userChatList map[*tcpx.Client]string

func init() {
	ChatroomList = make(map[string]*Chatroom, 10)
	userChatList = make(map[*tcpx.Client]string)
}

func initChatrooms() {
	if len(CHATROMMS) == 0 {
		log.Println("Init Chatrooms")
		ctList, err := ListChatrooms()
		if err != nil {
			log.Println("Service error")
		} else {
			for _, ct := range ctList {
				ChatroomList[ct.Name] = &Chatroom{
					ct: ct,
				}
				CHATROMMS = append(CHATROMMS, ct.Name)
			}
		}
	}
}

func List(client *tcpx.Client) {
	initChatrooms()
	client.PutOutgoing(fmt.Sprintf("You can choose one chatroom to join:\n%s",
		strings.Join(CHATROMMS, "\t")))
}

func Join(client *tcpx.Client, paramString string) {
	initChatrooms()
	params := strings.Fields(paramString)
	if len(params) != 1 {
		client.PutOutgoing("You can only input one param")
		return
	}
	ctName := params[0]
	if utils.StringInSlice(ctName, CHATROMMS) != -1 {
		Exit(client)
		userChatList[client] = ctName
		ChatroomList[ctName].clients = append(ChatroomList[ctName].clients, client)
		client.PutOutgoing(fmt.Sprintf("you have joined <%s> chatroom", ctName))
	}
}

func Exit(client *tcpx.Client) {
	if k, ok := userChatList[client]; ok {
		for i, c := range ChatroomList[k].clients {
			if c == client {
				ChatroomList[k].clients = append(ChatroomList[k].clients[:i],
					ChatroomList[k].clients[i+1:]...)
				client.PutOutgoing(fmt.Sprintf("you have exited <%s> chatroom", k))
				return
			}
		}
	}
}

func SendMsg(client *tcpx.Client, paramString string) {
	msg := paramString
	if ctName, ok := userChatList[client]; ok {
		for _, c := range ChatroomList[ctName].clients {
			c.PutOutgoing(fmt.Sprintf("%s says: %s",
				user.GetUserName(client),
				msg),
			)
		}
	}
}
