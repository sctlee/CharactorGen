package chatroom

import (
	"github.com/sctlee/tcpx/daemon/service"
)

func (self *ChatroomAction) GetRouteList() service.RouteList {
	return service.RouteList{
		"list": self.List,
		"view": self.View,
		"join": self.Join,
		"exit": self.Exit,
		"send": self.Send,
	}
}

//
// // f := auth.PermissionDecorate(client, chatroomAction.Send, auth.IsLogin)
// // responseMsg = f(client, params)
