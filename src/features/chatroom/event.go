package chatroom

import (
	"github.com/sctlee/tcpx"
	"github.com/sctlee/tcpx/daemon"
)

func (self *ChatroomAction) OnClientQuit(cid string) {
	msg := daemon.NewSimpleMessage("", "")
	msg.Src = cid
	msg.Des = "chatroom"
	msg.Command = "exit"
	msg.Params = make(map[string]string)
	msg.Type = daemon.MESSAGE_TYPE_TOSERVICE
	tcpx.SendMessage(msg)
}
