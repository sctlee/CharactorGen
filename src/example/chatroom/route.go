package chatroom

import (
	"strings"

	"github.com/sctlee/tcpx"
)

func Route(url string, client *tcpx.Client) {
	var action string
	url = strings.TrimSpace(url)
	i := strings.Index(url, " ")
	if i == -1 {
		action = url[:]
	} else {
		action = url[:i]
	}
	switch action {
	case "list":
		List(client)
	case "join":
		Join(client, strings.TrimSpace(url[i:]))
	case "send":
		SendMsg(client, strings.TrimSpace(url[i:]))
	}
}
