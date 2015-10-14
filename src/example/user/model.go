package user

import (
	"example/db"
	"log"
)

type User struct {
	id       int32
	Name     string
	Password string
}

func (self *User) Save() error {
	_, err := db.Pool.Exec("insert into account(name, password) values($1, $2)",
		self.Name, self.Password)
	return err
}

func Exists(name string, password string) (user *User, err error) {
	// user = nil
	// err = nil
	user = &User{}

	row := db.Pool.QueryRow("select id from account where name = $1 and password = $2", name, password)
	err = row.Scan(&user.id)
	if err == nil {
		user.Name = name
	} else {
		log.Println(err)
	}
	return
}
