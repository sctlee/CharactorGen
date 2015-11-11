package chatroom

import (
	"fmt"

	"features/auth"
	. "features/chatroom/action"

	"github.com/sctlee/tcpx"
)

//TODO: redefine struct function, then move the usage from here to example
var chatroomAction *ChatroomAction

func init() {
	chatroomAction = NewChatroomAction()
}

func Route(params map[string]string, client *tcpx.Client) tcpx.IMessage {
	var responseMsg tcpx.IMessage
	switch params["command"] {
	case "list":
		responseMsg = chatroomAction.List(client)
	case "view":
		responseMsg = chatroomAction.View(client, params)
	case "join":
		responseMsg = chatroomAction.Join(client, params)
	case "exit":
		responseMsg = chatroomAction.Exit(client)
	case "send":
		f := auth.PermissionDecorate(client, chatroomAction.Send, auth.IsLogin)
		responseMsg = f(client, params)
	default:
		return tcpx.NewMessage(client, fmt.Sprintf("no '%s' command", params["command"]))
	}
	return responseMsg
}
