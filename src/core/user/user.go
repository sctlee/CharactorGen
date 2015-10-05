package user

import (
	"core"
	"fmt"
	"strings"
)

var UserList map[*core.Client]string

func init() {
	UserList = make(map[*core.Client]string)
	core.RegisterRouter("user", Route)
}

func User(client *core.Client) string {
	s := UserList[client]
	if s != "" {
		return s
	} else {
		return "匿名"
	}
}

func Route(url string, client *core.Client) {
	url = strings.TrimSpace(url)
	fmt.Println(url)
	i := strings.Index(url, " ")
	action := url[:i]
	fmt.Println(url)
	switch action {
	case "setName":
		name := strings.TrimSpace(url[i:])
		UserList[client] = name
		client.PutOutgoing(fmt.Sprintf("Hello, %s", name))
	}
}
