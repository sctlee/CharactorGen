package chatroom

import (
	"fmt"
	"log"
	"strings"
	"time"

	. "features/chatroom/model"
	"features/user"

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
	go cleanChatrooms()
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

//check whether a closed client joined in a chatroom. if has, clean it.
func cleanChatrooms() {
	for {
		select {
		case <-time.After(time.Second * 2):
			for k, _ := range userChatList {
				if k.State == tcpx.CLIENT_STATE_CLOSE {
					Exit(k)
				}
			}
		}
	}
}

func List(client *tcpx.Client) {
	initChatrooms()
	client.PutOutgoing(fmt.Sprintf("You can choose one chatroom to join:\n%s",
		strings.Join(CHATROMMS, "\t")))
}

func View(client *tcpx.Client, params map[string]string) {
	if !utils.IsExistInMap(params, "ctName") {
		client.PutOutgoing("Please input ctName")
		return
	}
	ctName := params["ctName"]

	if utils.StringInSlice(ctName, CHATROMMS) != -1 {
		client.PutOutgoing(fmt.Sprintf("%d", len(ChatroomList[ctName].clients)))
	} else {
		client.PutOutgoing("the chatroom is not existed")
	}
}

func Join(client *tcpx.Client, params map[string]string) {
	initChatrooms()

	if !utils.IsExistInMap(params, "ctName") {
		client.PutOutgoing("Please input ctName")
		return
	}
	ctName := params["ctName"]

	if utils.StringInSlice(ctName, CHATROMMS) != -1 {
		Exit(client)
		userChatList[client] = ctName
		ChatroomList[ctName].clients = append(ChatroomList[ctName].clients, client)
		client.PutOutgoing(fmt.Sprintf("you have joined <%s> chatroom", ctName))
	} else {
		client.PutOutgoing(fmt.Sprintf("<%s> chatroom is not existed", ctName))
	}
}

func Exit(client *tcpx.Client) {
	if k, ok := userChatList[client]; ok {
		for i, c := range ChatroomList[k].clients {
			if c == client {
				ChatroomList[k].clients = append(ChatroomList[k].clients[:i],
					ChatroomList[k].clients[i+1:]...)
				client.PutOutgoing(fmt.Sprintf("you have exited <%s> chatroom", k))
				break
			}
		}
		delete(userChatList, client)
		SendMsg(k, user.GetUserName(client), "has exited")
	}
}

func Send(client *tcpx.Client, params map[string]string) {
	if !utils.IsExistInMap(params, "ctName") {
		client.PutOutgoing("Please input msg")
		return
	}
	if ctName, ok := userChatList[client]; ok {
		SendMsg(ctName, user.GetUserName(client), params["msg"])
	}
}

func SendMsg(ctName string, username string, msg string) {
	for _, c := range ChatroomList[ctName].clients {
		c.PutOutgoing(fmt.Sprintf("%s says: %s",
			username,
			msg),
		)
	}
}
