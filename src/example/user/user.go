package user

import (
	"core/client"
	"fmt"
	"strings"
)

var userList map[*client.Client]*User

type UserAuth string

func (s UserAuth) String() string {
	switch s {
	case "lhc":
		return "lhc"
	case "hc":
		return "hc"
		// ...
	default:
		return ""
	}
}

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

func Login(client *client.Client, paramString string) {
	userInfo := strings.Fields(paramString)
	if len(userInfo) != 2 {
		client.PutOutgoing("params number error: Please input correct number of params")
		return
	}
	password := userInfo[1]
	auth := UserAuth(userInfo[0])
	if auth != "" && strings.EqualFold(string(auth), password) {
		userList[client] = &User{
			Name: userInfo[0],
		}
		client.PutOutgoing("Login Success!")
	}
}

func Logout(client *client.Client) {
	if _, ok := userList[client]; ok {
		client.PutOutgoing("Logout success!")
	} else {
		client.PutOutgoing("Please login first!")
	}
}
