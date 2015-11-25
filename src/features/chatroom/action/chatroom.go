package action

import (
	"fmt"
	"log"
	"strings"
	// "time"

	. "features/chatroom/model"

	"github.com/sctlee/tcpx"
	"github.com/sctlee/tcpx/daemon/message"
	"github.com/sctlee/utils"
)

var CHATROMMS = []string{}

type ChatroomAction struct {
	ChatroomList map[string]*Chatroom // string is ctName
	UserChatList map[string]*Chatroom // string is ClientID
}

type Chatroom struct {
	ct   *ChatroomModel
	cids []string
}

func NewChatroomAction() *ChatroomAction {
	return &ChatroomAction{
		ChatroomList: make(map[string]*Chatroom, 10),
		UserChatList: make(map[string]*Chatroom),
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
				fmt.Println(ct)
				self.ChatroomList[ct.Name] = &Chatroom{
					ct: ct,
				}
				CHATROMMS = append(CHATROMMS, ct.Name)
			}
		}
	}
}

func (self *ChatroomAction) List(msg *message.Message) {
	self.initChatrooms()

	response := message.NewSimpleMessage(msg.Src,
		fmt.Sprintf("You can choose one chatroom to join:\n%s",
			strings.Join(CHATROMMS, "\t")))
	tcpx.SendMessage(response)
}

func (self *ChatroomAction) View(msg *message.Message) {
	if !utils.IsExistInMap(msg.Params, "ctName") {
		response := message.NewSimpleMessage(msg.Src,
			"Please input ctName")
		tcpx.SendMessage(response)
		return
	}
	ctName := msg.Params["ctName"]

	var response *message.Message
	if utils.StringInSlice(ctName, CHATROMMS) != -1 {
		response = message.NewSimpleMessage(msg.Src,
			fmt.Sprintf("%d", len(self.ChatroomList[ctName].cids)))
	} else {
		response = message.NewSimpleMessage(msg.Src,
			"the chatroom is not existed")
	}
	tcpx.SendMessage(response)
}
func (self *ChatroomAction) Join(msg *message.Message) {
	self.initChatrooms()

	if !utils.IsExistInMap(msg.Params, "ctName") {
		tcpx.SendMessage(message.NewSimpleMessage(msg.Src,
			"Please input ctName"))
		return
	}
	ctName := msg.Params["ctName"]

	if utils.StringInSlice(ctName, CHATROMMS) != -1 {
		self.Exit(msg)

		// TODO: event manager
		// cid.SetOnCloseListener(self)

		self.UserChatList[msg.Src] = self.ChatroomList[ctName]
		self.ChatroomList[ctName].cids = append(self.ChatroomList[ctName].cids, msg.Src)
		tcpx.SendMessage(message.NewSimpleMessage(msg.Src,
			fmt.Sprintf("you have joined <%s> chatroom", ctName)))
	} else {
		tcpx.SendMessage(message.NewSimpleMessage(msg.Src,
			fmt.Sprintf("<%s> chatroom is not existed", ctName)))
	}
}

func (self *ChatroomAction) Exit(msg *message.Message) {
	if chatroom, ok := self.UserChatList[msg.Src]; ok {
		for i, c := range chatroom.cids {
			if c == msg.Src {
				chatroom.cids = append(chatroom.cids[:i],
					chatroom.cids[i+1:]...)
				// cid.PutOutgoing(fmt.Sprintf("you have exited <%s> chatroom", chatroom.ct.Name))
				break
			}
		}
		delete(self.UserChatList, msg.Src)
		// self.SendMsg(chatroom, GetUserName(cid), "has exited")
		self.SendMsg(chatroom, "haha", "has exited")
	}
}

func (self *ChatroomAction) Send(msg *message.Message) {
	if !utils.IsExistInMap(msg.Params, "msg") {
		tcpx.SendMessage(message.NewSimpleMessage(msg.Src,
			"Please input msg"))
		return
	}
	if chatroom, ok := self.UserChatList[msg.Src]; ok {
		// self.SendMsg(chatroom, GetUserName(cid), params["msg"])
		self.SendMsg(chatroom, "haha", msg.Params["msg"])
	} else {
		tcpx.SendMessage(message.NewSimpleMessage(msg.Src,
			"You have not joined a chatroom"))
	}
}

func (self *ChatroomAction) SendMsg(chatroom *Chatroom, username string, msg string) {
	// fatalCids := []string{"1.0", "0.0"}
	response := message.NewSimpleBoardMessage(
		chatroom.cids,
		fmt.Sprintf("%s says: %s",
			username,
			msg))
	tcpx.SendMessage(response)

}

// func (self *ChatroomAction) OnClose(cid daemon.ClientID) {
// 	self.Exit(cid, nil)
// }

// func GetUserName(cid daemon.ClientID) (username string) {
// username, ok := cid.GetSharedPreferences("Auth").Get("username")
// if !ok {
// 	username = "匿名"
// }

// }

//check whether a closed cid joined in a chatroom. if has, clean it.
// func cleanChatroomModel() {
// 	for {
// 		select {
// 		case <-time.After(time.Second * 2):
// 			for k, _ := range userChatList {
// 				if k.State == daemon.cid_STATE_CLOSE {
// 					Exit(k)
// 				}
// 			}
// 		}
// 	}
// }
