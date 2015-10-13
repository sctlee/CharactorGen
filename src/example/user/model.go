package user

import (
	"example/db"
)

type User struct {
	Name     string
	Password string
}

func (self *User) Save() error {
	_, err := db.Pool.Exec("insert into account(name, password) values($1, $2)",
		self.Name, self.Password)
	return err
}
