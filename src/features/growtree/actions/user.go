package actions

import (
	"fmt"
	"strings"

	. "features/growtree/model"

	"github.com/sctlee/tcpx"
	"github.com/sctlee/utils"
)

var userList map[*tcpx.Client]*User

func init() {
	userList = make(map[*tcpx.Client]*User)
}

func GetUserName(client *tcpx.Client) string {
	s := userList[client]
	if s != nil {
		return s.Name
	} else {
		return "匿名"
	}
}

func SetUserName(client *tcpx.Client, params map[string]string) {
	if !utils.IsExistInMap(params, "name") {
		client.PutOutgoing("Please input name")
		return
	}
	name := params["name"]

	userList[client] = &User{
		Name: name,
	}
	client.PutOutgoing(fmt.Sprintf("Hello, %s", name))
}

func Login(client *tcpx.Client, params map[string]string) {
	/*
		use postgresql
	*/
	if !utils.IsExistInMap(params, "username", "password") {
		client.PutOutgoing("params error")
		return
	}
	username := params["username"]
	password := params["password"]

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
	if !utils.IsExistInMap(params, "username", "password", "confitm") {
		client.PutOutgoing("params error")
		return
	}
	username := params["username"]
	password := params["password"]
	confirm := params["confirm"]

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
