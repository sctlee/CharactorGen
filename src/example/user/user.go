package user

import (
	"core/client"
	"fmt"
)

func GetUserName(client *client.Client) string {
	s := userList[client]
	if s != nil {
		return s.Name
	} else {
		return "匿名"
	}
}

func SetUserName(client *client.Client, paramString string) {
	name := paramString
	userList[client] = &User{
		Name: name,
	}
	client.PutOutgoing(fmt.Sprintf("Hello, %s", name))
}
