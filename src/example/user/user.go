package user

import (
	"core"
	"fmt"
)

func GetUserName(client *core.Client) string {
	s := userList[client]
	if s != nil {
		return s.Name
	} else {
		return "匿名"
	}
}

func SetUserName(client *core.Client, paramString string) {
	name := paramString
	userList[client] = &User{
		Name: name,
	}
	client.PutOutgoing(fmt.Sprintf("Hello, %s", name))
}
