package chatroom

import (
	"fmt"
	"log"
	"strings"
	// "time"

	. "features/chatroom/model"

	"github.com/sctlee/hazel"
	"github.com/sctlee/hazel/daemon"
	"github.com/sctlee/hazel/daemon/message"
	"github.com/sctlee/hazel/db"
	"github.com/sctlee/utils"

	"github.com/garyburd/redigo/redis"
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
	c := &ChatroomAction{
		ChatroomList: make(map[string]*Chatroom, 10),
		UserChatList: make(map[string]*Chatroom),
	}
	c.initChatrooms()
	return c
}

func (self *ChatroomAction) initChatrooms() {
	conn := db.RedisPool.Get()
	defer conn.Close()
	err := conn.Send("DEL", "chatrooms")
	if err != nil {
		fmt.Println("redis error:", err)
	}
	if len(CHATROMMS) == 0 {
		log.Println("Init ChatroomModel")
		ctList, err := ListChatroomModel()
		if err != nil {
			log.Println("example error:", err)
		} else {
			for _, ct := range ctList {
				fmt.Println(ct)
				self.ChatroomList[ct.Name] = &Chatroom{
					ct: ct,
				}

				// original memory way
				// CHATROMMS = append(CHATROMMS, ct.Name)
				// use redis
				err := conn.Send("LPUSH", "chatrooms", ct.Name)
				if err != nil {
					fmt.Println("redis error:", err)
				}
			}
		}
	}
}

func (self *ChatroomAction) List(msg *message.Message) {
	conn := db.RedisPool.Get()
	defer conn.Close()

	chatrooms, err := redis.Strings(conn.Do("LRANGE", "chatrooms", "0", "-1"))
	if err != nil {
		fmt.Println("redis error:", err)
		hazel.SendMessage(daemon.NewSimpleMessage(msg.Src,
			"service error"))
	}
	fmt.Printf("%t", chatrooms)
	response := daemon.NewSimpleMessage(msg.Src,
		fmt.Sprintf("You can choose one chatroom to join:\n%s",
			strings.Join(chatrooms, "\t")))
	//original way: CHATROOMS -> chatrooms
	hazel.SendMessage(response)
}

func (self *ChatroomAction) View(msg *message.Message) {
	if !utils.IsExistInMap(msg.Params, "ctName") {
		response := daemon.NewSimpleMessage(msg.Src,
			"Please input ctName")
		hazel.SendMessage(response)
		return
	}
	conn := db.RedisPool.Get()
	defer conn.Close()

	ctName := msg.Params["ctName"]

	chatrooms, err := redis.Strings(conn.Do("LRANGE", "chatrooms", "0", "-1"))
	if err != nil {
		fmt.Println("redis error:", err)
	}

	var response *message.Message
	if utils.StringInSlice(ctName, chatrooms) != -1 {
		if number, err := redis.Int(conn.Do("LLEN", "chatroom:"+ctName)); err != nil {
			fmt.Println("redis err:", err)
		} else {
			response = daemon.NewSimpleMessage(msg.Src,
				fmt.Sprintf("%d", number))
			// fmt.Sprintf("%d", len(self.ChatroomList[ctName].cid)))
		}
	} else {
		response = daemon.NewSimpleMessage(msg.Src,
			"the chatroom is not existed")
	}
	hazel.SendMessage(response)
}
func (self *ChatroomAction) Join(msg *message.Message) {
	if !utils.IsExistInMap(msg.Params, "ctName") {
		hazel.SendMessage(daemon.NewSimpleMessage(msg.Src,
			"Please input ctName"))
		return
	}
	ctName := msg.Params["ctName"]

	conn := db.RedisPool.Get()
	defer conn.Close()

	chatrooms, err := redis.Strings(conn.Do("LRANGE", "chatrooms", "0", "-1"))
	if err != nil {
		fmt.Println("redis error:", err)
	}

	// original CHATROOMS -> chatroom
	// now use redis
	if utils.StringInSlice(ctName, chatrooms) != -1 {
		self.Exit(msg)

		// original event manager,
		// now use interface + event trigger, see event.go
		// cid.SetOnCloseListener(self)

		// original way
		// self.UserChatList[msg.Src] = self.ChatroomList[ctName]
		// self.ChatroomList[ctName].cids = append(self.ChatroomList[ctName].cids, msg.Src)

		// now use redis
		err1 := conn.Send("HSET", msg.Src, "chatroom", ctName)
		err2 := conn.Send("LPUSH", "chatroom:"+ctName, msg.Src)
		if err1 != nil || err2 != nil {
			hazel.SendMessage(daemon.NewSimpleMessage(msg.Src,
				fmt.Sprintf("Failed")))
			return
		}
		hazel.SendMessage(daemon.NewSimpleMessage(msg.Src,
			fmt.Sprintf("you have joined <%s> chatroom", ctName)))
	} else {
		hazel.SendMessage(daemon.NewSimpleMessage(msg.Src,
			fmt.Sprintf("<%s> chatroom is not existed", ctName)))
	}
}

func (self *ChatroomAction) Exit(msg *message.Message) {
	conn := db.RedisPool.Get()
	defer conn.Close()

	/*	original way:
		delete cid in chatroom.cids by manual delete from slice
	*/
	// if chatroom, ok := self.UserChatList[msg.Src]; ok {
	// for i, c := range chatroom.cids {
	// 	if c == msg.Src {
	// 		chatroom.cids = append(chatroom.cids[:i],
	// 			chatroom.cids[i+1:]...)
	// 		// cid.PutOutgoing(fmt.Sprintf("you have exited <%s> chatroom", chatroom.ct.Name))
	// 		break
	// 	}
	// }
	// delete(self.UserChatList, msg.Src)

	/*	now use redis:
		delete easily
	*/
	exists, err := redis.Bool(conn.Do("HEXISTS", msg.Src, "chatroom"))
	if err != nil {
		fmt.Println("redis err:", err)
		return
	}
	if exists {
		ctName, err := redis.String(conn.Do("HGET", msg.Src, "chatroom"))
		err = conn.Send("LREM", "chatroom:"+ctName, "0", msg.Src)
		err = conn.Send("HDEL", msg.Src, "chatroom")
		cids, err := redis.Strings(conn.Do("LRANGE", "chatroom:"+ctName, "0", "-1"))
		if err != nil {
			fmt.Println("redis error:", err)
			hazel.SendMessage(daemon.NewSimpleMessage(msg.Src,
				fmt.Sprintf("Failed")))
			return
		}
		hazel.SendMessage(daemon.NewSimpleMessage(msg.Src,
			fmt.Sprintf("Quited Success")))

		self.SendMsg(cids, GetUserName(msg.Src, msg), "has exited")
	}
}

func (self *ChatroomAction) Send(msg *message.Message) {
	if !utils.IsExistInMap(msg.Params, "msg") {
		hazel.SendMessage(daemon.NewSimpleMessage(msg.Src,
			"Please input msg"))
		return
	}
	conn := db.RedisPool.Get()
	defer conn.Close()
	// if chatroom, ok := self.UserChatList[msg.Src]; ok {
	if ctName, err := redis.String(conn.Do("HGET", msg.Src, "chatroom")); err == nil {
		// self.SendMsg(chatroom, GetUserName(cid), params["msg"])
		cids, err := redis.Strings(conn.Do("LRANGE", "chatroom:"+ctName, "0", "-1"))
		if err != nil {
			fmt.Println("redis error", err)
			return
		}
		self.SendMsg(cids, GetUserName(msg.Src, msg), msg.Params["msg"])
		// self.SendMsg(chatroom, GetUserName(msg, msg.Response), msg.Params["msg"])
	} else {
		hazel.SendMessage(daemon.NewSimpleMessage(msg.Src,
			"You have not joined a chatroom"))
	}
}

// original way
// func (self *ChatroomAction) SendMsg(chatroom *Chatroom, username string, msg string) {
// now use redis
func (self *ChatroomAction) SendMsg(cids []string, username string, msg string) {
	// fatalCids := []string{"1.0", "0.0"}
	response := daemon.NewSimpleBoardMessage(
		cids,
		fmt.Sprintf("%s says: %s",
			username,
			msg))
	hazel.SendMessage(response)

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
