package action

import (
	"fmt"
	"strings"

	. "features/auth/model"

	"github.com/sctlee/tcpx"
	"github.com/sctlee/utils"
)

type UserAction struct {
	userList map[*tcpx.Client]*UserModel
}

func NewUserAction() *UserAction {
	return &UserAction{
		userList: make(map[*tcpx.Client]*UserModel),
	}
}

func (self *UserAction) GetUserName(client *tcpx.Client) string {
	s := self.userList[client]
	if s != nil {
		return s.Name
	} else {
		return "匿名"
	}
}

func (self *UserAction) SetUserName(client *tcpx.Client, params map[string]string) tcpx.IMessage {
	if !utils.IsExistInMap(params, "name") {
		return tcpx.NewMessage(client, "Please input name")
	}
	name := params["name"]

	self.userList[client] = &UserModel{
		Name: name,
	}
	return tcpx.NewMessage(client, fmt.Sprintf("Hello, %s", name))
}

func (self *UserAction) Login(client *tcpx.Client, params map[string]string) tcpx.IMessage {
	/*
		use postgresql
	*/
	if !utils.IsExistInMap(params, "username", "password") {
		return tcpx.NewMessage(client, "params error")
	}
	username := params["username"]
	password := params["password"]

	user, err := Exists(username, password)
	if err != nil {
		return tcpx.NewMessage(client, "Username or password error!")
	} else {
		self.userList[client] = user

		// save login status in client.sharedPreferences
		sp := client.GetSharedPreferences("Auth")
		sp.Set("username", user.Name)

		return tcpx.NewMessage(client, "Login Success!")
	}
}

func (self *UserAction) Logout(client *tcpx.Client) tcpx.IMessage {
	if _, ok := self.userList[client]; ok {
		delete(self.userList, client)
		return tcpx.NewMessage(client, "Logout success!")
	} else {
		return tcpx.NewMessage(client, "Please login first!")
	}
}

func (self *UserAction) Signup(client *tcpx.Client, params map[string]string) tcpx.IMessage {
	if !utils.IsExistInMap(params, "username", "password", "confitm") {
		return tcpx.NewMessage(client, "params error")
	}
	username := params["username"]
	password := params["password"]
	confirm := params["confirm"]

	if strings.EqualFold(password, confirm) {
		user := &UserModel{
			Name:     username,
			Password: password,
		}
		if err := user.Save(); err != nil {
			return tcpx.NewMessage(client, "Signup error!")
		} else {
			return tcpx.NewMessage(client, "Signup success! Now you can login with your account!")
		}
	} else {
		return tcpx.NewMessage(client, "confirm is not equal to password")
	}
}
