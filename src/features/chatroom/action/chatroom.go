package action

import (
	"fmt"
	"log"
	"strings"
	// "time"

	. "features/chatroom/model"

	"github.com/sctlee/tcpx"
	"github.com/sctlee/utils"
)

var CHATROMMS = []string{}

type ChatroomAction struct {
	ChatroomList map[string]*Chatroom
	UserChatList map[*tcpx.Client]*Chatroom
}

type Chatroom struct {
	ct      *ChatroomModel
	clients []*tcpx.Client
}

func NewChatroomAction() *ChatroomAction {
	return &ChatroomAction{
		ChatroomList: make(map[string]*Chatroom, 10),
		UserChatList: make(map[*tcpx.Client]*Chatroom),
	}
}

func (self *ChatroomAction) initChatrooms() {
	if len(CHATROMMS) == 0 {
		log.Println("Init ChatroomModel")
		ctList, err := ListChatroomModel()
		if err != nil {
			log.Println(err)
		} else {
			for _, ct := range ctList {
				self.ChatroomList[ct.Name] = &Chatroom{
					ct: ct,
				}
				CHATROMMS = append(CHATROMMS, ct.Name)
			}
		}
	}
}

func (self *ChatroomAction) List(client *tcpx.Client) tcpx.IMessage {
	self.initChatrooms()
	return tcpx.NewMessage(client, fmt.Sprintf("You can choose one chatroom to join:\n%s",
		strings.Join(CHATROMMS, "\t")))
}

func (self *ChatroomAction) View(client *tcpx.Client, params map[string]string) tcpx.IMessage {
	if !utils.IsExistInMap(params, "ctName") {
		return tcpx.NewMessage(client, "Please input ctName")
	}
	ctName := params["ctName"]

	if utils.StringInSlice(ctName, CHATROMMS) != -1 {
		return tcpx.NewMessage(client, fmt.Sprintf("%d", len(self.ChatroomList[ctName].clients)))
	} else {
		return tcpx.NewMessage(client, "the chatroom is not existed")
	}
}
func (self *ChatroomAction) Join(client *tcpx.Client, params map[string]string) tcpx.IMessage {
	self.initChatrooms()

	if !utils.IsExistInMap(params, "ctName") {
		return tcpx.NewMessage(client, "Please input ctName")
	}
	ctName := params["ctName"]

	if utils.StringInSlice(ctName, CHATROMMS) != -1 {
		self.Exit(client)

		client.SetOnCloseListener(self)

		self.UserChatList[client] = self.ChatroomList[ctName]
		self.ChatroomList[ctName].clients = append(self.ChatroomList[ctName].clients, client)
		return tcpx.NewMessage(client, fmt.Sprintf("you have joined <%s> chatroom", ctName))
	} else {
		return tcpx.NewMessage(client, fmt.Sprintf("<%s> chatroom is not existed", ctName))
	}
}

func (self *ChatroomAction) Exit(client *tcpx.Client) tcpx.IMessage {
	if chatroom, ok := self.UserChatList[client]; ok {
		for i, c := range chatroom.clients {
			if c == client {
				chatroom.clients = append(chatroom.clients[:i],
					chatroom.clients[i+1:]...)
				client.PutOutgoing(fmt.Sprintf("you have exited <%s> chatroom", chatroom.ct.Name))
				break
			}
		}
		delete(self.UserChatList, client)
		return self.SendMsg(chatroom, GetUserName(client), "has exited")
	}
	return tcpx.NewMessage(client, "You have not joined a chatroom")
}

func (self *ChatroomAction) Send(client *tcpx.Client, params map[string]string) tcpx.IMessage {
	if !utils.IsExistInMap(params, "msg") {
		return tcpx.NewMessage(client, "Please input msg")
	}
	if chatroom, ok := self.UserChatList[client]; ok {
		return self.SendMsg(chatroom, GetUserName(client), params["msg"])
	}
	return tcpx.NewMessage(client, "You have not joined a chatroom")
}

func (self *ChatroomAction) SendMsg(chatroom *Chatroom, username string, msg string) tcpx.IMessage {
	return tcpx.NewBoardMessage(nil,
		fmt.Sprintf("%s says: %s",
			username,
			msg),
		chatroom.clients)
}

func (self *ChatroomAction) OnClose(client *tcpx.Client) {
	self.Exit(client)
}

func GetUserName(client *tcpx.Client) (username string) {
	username, ok := client.GetSharedPreferences("Auth").Get("username")
	if !ok {
		username = "匿名"
	}
	return
}

//check whether a closed client joined in a chatroom. if has, clean it.
// func cleanChatroomModel() {
// 	for {
// 		select {
// 		case <-time.After(time.Second * 2):
// 			for k, _ := range userChatList {
// 				if k.State == tcpx.CLIENT_STATE_CLOSE {
// 					Exit(k)
// 				}
// 			}
// 		}
// 	}
// }
