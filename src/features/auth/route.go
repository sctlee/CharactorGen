package auth

import (
	"fmt"

	. "features/auth/action"

	"github.com/sctlee/tcpx"
)

var userAction *UserAction

func init() {
	userAction = NewUserAction()
}

func Route(params map[string]string, client *tcpx.Client) tcpx.IMessage {
	var responseMsg tcpx.IMessage
	switch params["command"] {
	case "setName":
		responseMsg = userAction.SetUserName(client, params)
	case "login":
		responseMsg = userAction.Login(client, params)
	case "logout":
		responseMsg = userAction.Logout(client)
	case "signup":
		responseMsg = userAction.Signup(client, params)
	default:
		return tcpx.NewMessage(client, fmt.Sprintf("no '%s' command", params["command"]))
	}
	return responseMsg
}
