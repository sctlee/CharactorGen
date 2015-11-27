package chatroom

import (
	"time"

	"github.com/sctlee/tcpx"
	"github.com/sctlee/tcpx/daemon"
	"github.com/sctlee/tcpx/daemon/message"
)

func GetUserName(cid string, originalMsg *message.Message) (username string) {
	request := message.NewMessage(nil, "chatroom", "auth",
		map[string]string{"command": "getusername", "cid": cid},
		daemon.MESSAGE_TYPE_TOSERVICE)
	message.CopySession(originalMsg, request)
	tcpx.SendMessage(request)
	select {
	case msg := <-request.Response:
		return msg.Params["username"]
	case <-time.After(time.Second * 2):
		return "匿名(auth time out)"
	}
}
