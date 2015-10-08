package user

import (
	"core/client"
	"fmt"
	"strings"
)

func init() {
	userList = make(map[*client.Client]*User)
}

func Route(url string, client *client.Client) {
	var action string
	url = strings.TrimSpace(url)
	i := strings.Index(url, " ")
	if i == -1 {
		action = url[:]
	} else {
		action = url[:i]
	}
	switch action {
	case "setName":
		SetUserName(client, strings.TrimSpace(url[i:]))
	case "login":
		Login(client, strings.TrimSpace(url[i:]))
	case "logout":
		fmt.Println("1")
		Logout(client)
	}
}
