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
	UserChatList map[tcpx.ClientID]*Chatroom
}

type Chatroom struct {
	ct   *ChatroomModel
	cids []tcpx.ClientID
}

func NewChatroomAction() *ChatroomAction {
	return &ChatroomAction{
		ChatroomList: make(map[string]*Chatroom, 10),
		UserChatList: make(map[tcpx.ClientID]*Chatroom),
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

func (self *ChatroomAction) List(cid tcpx.ClientID, params map[string]string) {
	self.initChatrooms()
	msg := tcpx.NewMessage(cid,
		fmt.Sprintf("You can choose one chatroom to join:\n%s",
			strings.Join(CHATROMMS, "\t")))
	tcpx.SendMessage(msg)
}

func (self *ChatroomAction) View(cid tcpx.ClientID, params map[string]string) {
	if !utils.IsExistInMap(params, "ctName") {
		tcpx.SendMessage(tcpx.NewMessage(cid, "Please input ctName"))
		return
	}
	ctName := params["ctName"]

	var msg tcpx.IMessage
	if utils.StringInSlice(ctName, CHATROMMS) != -1 {
		msg = tcpx.NewMessage(cid,
			fmt.Sprintf("%d", len(self.ChatroomList[ctName].cids)))
	} else {
		msg = tcpx.NewMessage(cid, "the chatroom is not existed")
	}
	tcpx.SendMessage(msg)
}
func (self *ChatroomAction) Join(cid tcpx.ClientID, params map[string]string) {
	self.initChatrooms()

	if !utils.IsExistInMap(params, "ctName") {
		tcpx.SendMessage(tcpx.NewMessage(cid, "Please input ctName"))
		return
	}
	ctName := params["ctName"]

	if utils.StringInSlice(ctName, CHATROMMS) != -1 {
		self.Exit(cid, params)

		// TODO: event manager
		// cid.SetOnCloseListener(self)

		self.UserChatList[cid] = self.ChatroomList[ctName]
		self.ChatroomList[ctName].cids = append(self.ChatroomList[ctName].cids, cid)
		tcpx.SendMessage(tcpx.NewMessage(cid, fmt.Sprintf("you have joined <%s> chatroom", ctName)))
	} else {
		tcpx.SendMessage(tcpx.NewMessage(cid, fmt.Sprintf("<%s> chatroom is not existed", ctName)))
	}
}

func (self *ChatroomAction) Exit(cid tcpx.ClientID, params map[string]string) {
	if chatroom, ok := self.UserChatList[cid]; ok {
		for i, c := range chatroom.cids {
			if c == cid {
				chatroom.cids = append(chatroom.cids[:i],
					chatroom.cids[i+1:]...)
				// cid.PutOutgoing(fmt.Sprintf("you have exited <%s> chatroom", chatroom.ct.Name))
				break
			}
		}
		delete(self.UserChatList, cid)
		// self.SendMsg(chatroom, GetUserName(cid), "has exited")
		self.SendMsg(chatroom, "haha", "has exited")
	}
}

func (self *ChatroomAction) Send(cid tcpx.ClientID, params map[string]string) {
	if !utils.IsExistInMap(params, "msg") {
		tcpx.SendMessage(tcpx.NewMessage(cid, "Please input msg"))
		return
	}
	if chatroom, ok := self.UserChatList[cid]; ok {
		// self.SendMsg(chatroom, GetUserName(cid), params["msg"])
		self.SendMsg(chatroom, "haha", params["msg"])
	} else {
		tcpx.SendMessage(tcpx.NewMessage(cid, "You have not joined a chatroom"))
	}
}

func (self *ChatroomAction) SendMsg(chatroom *Chatroom, username string, msg string) {
	fatalCids := []tcpx.ClientID{"1.0", "0.0"}
	tcpx.SendMessage(tcpx.NewBoardMessage("",
		fmt.Sprintf("%s says: %s",
			username,
			msg),
		// chatroom.cid))
		fatalCids))
}

func (self *ChatroomAction) OnClose(cid tcpx.ClientID) {
	self.Exit(cid, nil)
}

// func GetUserName(cid tcpx.ClientID) (username string) {
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
// 				if k.State == tcpx.cid_STATE_CLOSE {
// 					Exit(k)
// 				}
// 			}
// 		}
// 	}
// }
