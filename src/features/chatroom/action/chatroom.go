package action

import (
	"fmt"
	"log"
	"strings"
	// "time"

	. "features/chatroom/model"
	ga "features/growtree/action"

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

func (self *ChatroomAction) List(client *tcpx.Client) {
	self.initChatrooms()
	client.PutOutgoing(fmt.Sprintf("You can choose one chatroom to join:\n%s",
		strings.Join(CHATROMMS, "\t")))
}

func (self *ChatroomAction) View(client *tcpx.Client, params map[string]string) {
	if !utils.IsExistInMap(params, "ctName") {
		client.PutOutgoing("Please input ctName")
		return
	}
	ctName := params["ctName"]

	if utils.StringInSlice(ctName, CHATROMMS) != -1 {
		client.PutOutgoing(fmt.Sprintf("%d", len(self.ChatroomList[ctName].clients)))
	} else {
		client.PutOutgoing("the chatroom is not existed")
	}
}
func (self *ChatroomAction) OnClose(client *tcpx.Client) {
	self.Exit(client)
}

func (self *ChatroomAction) Join(client *tcpx.Client, params map[string]string) {
	self.initChatrooms()

	if !utils.IsExistInMap(params, "ctName") {
		client.PutOutgoing("Please input ctName")
		return
	}
	ctName := params["ctName"]

	if utils.StringInSlice(ctName, CHATROMMS) != -1 {
		self.Exit(client)

		client.SetOnCloseListener(self)

		self.UserChatList[client] = self.ChatroomList[ctName]
		self.ChatroomList[ctName].clients = append(self.ChatroomList[ctName].clients, client)
		client.PutOutgoing(fmt.Sprintf("you have joined <%s> chatroom", ctName))
	} else {
		client.PutOutgoing(fmt.Sprintf("<%s> chatroom is not existed", ctName))
	}
}

func (self *ChatroomAction) Exit(client *tcpx.Client) {
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
		self.SendMsg(chatroom, ga.GetUserName(client), "has exited")
	}
}

func (self *ChatroomAction) Send(client *tcpx.Client, params map[string]string) {
	if !utils.IsExistInMap(params, "msg") {
		client.PutOutgoing("Please input msg")
		return
	}
	if chatroom, ok := self.UserChatList[client]; ok {
		self.SendMsg(chatroom, ga.GetUserName(client), params["msg"])
	}
}

func (self *ChatroomAction) SendMsg(chatroom *Chatroom, username string, msg string) {
	for _, c := range chatroom.clients {
		c.PutOutgoing(fmt.Sprintf("%s says: %s",
			username,
			msg),
		)
	}
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
