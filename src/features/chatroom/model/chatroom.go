package model

import (
	"github.com/sctlee/hazel/db"
)

type ChatroomModel struct {
	id   int32
	Name string
	// MaxClients int16
	Class string
}

func ListChatroomModel() (chatroomList []*ChatroomModel, err error) {
	chatroomList = make([]*ChatroomModel, 0)

	rows, _ := db.Pool.Query("select id, name, class from Chatrooms")

	for rows.Next() {
		ct := &ChatroomModel{}
		err := rows.Scan(&ct.id, &ct.Name, &ct.Class)
		if err == nil {
			chatroomList = append(chatroomList, ct)
		}
	}
	err = rows.Err()
	return

}
