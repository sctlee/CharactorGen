package chatroom

import (
	. "features/chatroom/action"

	"github.com/sctlee/tcpx"
)

//TODO: redefine struct function, then move the usage from here to example
var chatroomAction *ChatroomAction

func init() {
	chatroomAction = NewChatroomAction()
}

func Route(params map[string]string, client *tcpx.Client) {
	switch params["command"] {
	case "list":
		chatroomAction.List(client)
	case "view":
		chatroomAction.View(client, params)
	case "join":
		chatroomAction.Join(client, params)
	case "exit":
		chatroomAction.Exit(client)
	case "send":
		chatroomAction.Send(client, params)
	}
}
