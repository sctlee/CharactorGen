package user

import (
	"github.com/sctlee/tcpx"
)

func init() {
	userList = make(map[*tcpx.Client]*User)
}

func Route(params map[string]string, client *tcpx.Client) {
	switch params["command"] {
	case "setName":
		SetUserName(client, params)
	case "login":
		Login(client, params)
	case "logout":
		Logout(client)
	case "signup":
		Signup(client, params)
	}
}
