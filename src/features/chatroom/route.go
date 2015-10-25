package chatroom

import (
	"github.com/sctlee/tcpx"
)

func Route(params map[string]string, client *tcpx.Client) {
	switch params["command"] {
	case "list":
		List(client)
	case "join":
		Join(client, params)
	case "send":
		SendMsg(client, params)
	}
}
