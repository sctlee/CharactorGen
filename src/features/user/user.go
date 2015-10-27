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

func SetUserName(client *tcpx.Client, params map[string]string) {
	name, ok := params["name"]
	if !ok {
		client.PutOutgoing("Please input name")
		return
	}
	userList[client] = &User{
		Name: name,
	}
	client.PutOutgoing(fmt.Sprintf("Hello, %s", name))
}

func Login(client *tcpx.Client, params map[string]string) {
	/*
		use postgresql
	*/
	username, ok1 := params["username"]
	password, ok2 := params["password"]
	if !ok1 || !ok2 {
		client.PutOutgoing("params error")
		return
	}
	user, err := Exists(username, password)
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

func Signup(client *tcpx.Client, params map[string]string) {
	username, ok1 := params["username"]
	password, ok2 := params["password"]
	confirm, ok3 := params["confirm"]

	if !ok1 || !ok2 || !ok3 {
		client.PutOutgoing("params error")
		return
	}

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
	} else {
		client.PutOutgoing("confirm is not equal to password")
	}
}
