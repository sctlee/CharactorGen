package chatroom

import (
	"example/db"
)

type Chatrooms struct {
	id   int32
	Name string
	// MaxClients int16
	Class string
}

func ListChatrooms() (chatroomList []*Chatrooms, err error) {
	chatroomList = make([]*Chatrooms, 0)

	rows, _ := db.Pool.Query("select id, name, class from chatrooms")

	for rows.Next() {
		ct := &Chatrooms{}
		err := rows.Scan(&ct.id, &ct.Name, &ct.Class)
		if err == nil {
			chatroomList = append(chatroomList, ct)
		}
	}
	err = rows.Err()
	return

}
