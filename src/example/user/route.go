package user

import (
	"core/client"
	"fmt"
	"strings"
)

var userList map[*client.Client]*User

func init() {
	userList = make(map[*client.Client]*User)
}

func Route(url string, client *client.Client) {
	url = strings.TrimSpace(url)
	fmt.Println(url)
	i := strings.Index(url, " ")
	action := url[:i]
	switch action {
	case "setName":
		SetUserName(client, strings.TrimSpace(url[i:]))
	}
}
