package user

import (
	"core"
	"fmt"
	"strings"
)

var userList map[*core.Client]*User

func init() {
	userList = make(map[*core.Client]*User)
}

func Route(url string, client *core.Client) {
	url = strings.TrimSpace(url)
	fmt.Println(url)
	i := strings.Index(url, " ")
	action := url[:i]
	switch action {
	case "setName":
		SetUserName(client, strings.TrimSpace(url[i:]))
	}
}
