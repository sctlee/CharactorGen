package user

import (
	"fmt"
	"strings"

	"github.com/sctlee/tcpx"
)

var userList map[*tcpx.Client]*User

func GetUserName(client *tcpx.Client) string {
	s := userList[client]
	if s != nil {
		return s.Name
	} else {
		return "匿名"
	}
}

func SetUserName(client *tcpx.Client, paramString string) {
	name := paramString
	userList[client] = &User{
		Name: name,
	}
	client.PutOutgoing(fmt.Sprintf("Hello, %s", name))
}

func Login(client *tcpx.Client, paramString string) {
	userInfo := strings.Fields(paramString)
	if len(userInfo) != 2 {
		client.PutOutgoing("Params number error: Please input correct number of params")
		return
	}

	/*
		use postgresql
	*/
	user, err := Exists(userInfo[0], userInfo[1])
	if err != nil {
		client.PutOutgoing("Username or password error!")
	} else {
		userList[client] = user
		client.PutOutgoing("Login Success!")
	}
}

func Logout(client *tcpx.Client) {
	if _, ok := userList[client]; ok {
		delete(userList, client)
		client.PutOutgoing("Logout success!")
	} else {
		client.PutOutgoing("Please login first!")
	}
}

func Signup(client *tcpx.Client, paramString string) {
	userInfo := strings.Fields(paramString)
	if len(userInfo) != 3 {
		client.PutOutgoing("Params number error: " +
			"Please input three params(username, password, confirm)")
		return
	}
	username := userInfo[0]
	password := userInfo[1]
	confirm := userInfo[2]

	if strings.EqualFold(password, confirm) {
		user := &User{
			Name:     username,
			Password: password,
		}
		if err := user.Save(); err != nil {
			client.PutOutgoing("Signup error!")
		} else {
			client.PutOutgoing("Signup success! Now you can login with your account!")
		}
	}
}
