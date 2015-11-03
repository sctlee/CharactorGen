package chatroom

import (
	. "features/chatroom/action"

	"github.com/sctlee/tcpx"
)

func Route(params map[string]string, client *tcpx.Client) {
	switch params["command"] {
	case "list":
		List(client)
	case "view":
		View(client, params)
	case "join":
		Join(client, params)
	case "exit":
		Exit(client)
	case "send":
		Send(client, params)
	}
}
