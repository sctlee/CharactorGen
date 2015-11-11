package auth

import (
	"reflect"

	"github.com/sctlee/tcpx"
)

const (
	IsLogin = "IsLogin"
)

type Permission struct {
	auth tcpx.SharedPreferences
}

func (self Permission) IsLogin(client *tcpx.Client) bool {
	self.auth = client.GetSharedPreferences("Auth")
	if _, ok := self.auth.Get("name"); ok {
		return true
	}
	return false
}

type responseFunc func(client *tcpx.Client, params map[string]string) tcpx.IMessage

var permission = new(Permission)

func PermissionDecorate(client *tcpx.Client, f responseFunc, permissions ...string) responseFunc {
	for _, methodName := range permissions {
		method := reflect.ValueOf(permission).MethodByName(methodName)
		if method.Interface().(func(client *tcpx.Client) bool)(client) {
			return f
		}
	}
	nf := func(client *tcpx.Client, params map[string]string) tcpx.IMessage {
		return tcpx.NewMessage(client, "Permission refused")
	}
	return nf
}
