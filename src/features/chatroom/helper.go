package chatroom

import (
	"time"

	"github.com/sctlee/hazel"
	"github.com/sctlee/hazel/daemon"
	"github.com/sctlee/hazel/daemon/message"
)

func GetUserName(cid string, originalMsg *message.Message) (username string) {
	request := message.NewMessage(nil, "chatroom", "auth",
		map[string]string{"command": "getusername", "cid": cid},
		daemon.MESSAGE_TYPE_TOSERVICE)
	message.CopySession(originalMsg, request)
	hazel.SendMessage(request)
	select {
	case msg := <-request.Response:
		return msg.Params["username"]
	case <-time.After(time.Second * 2):
		return "匿名(auth time out)"
	}
}
