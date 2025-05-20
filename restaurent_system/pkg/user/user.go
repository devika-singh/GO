package user

import (
	"fmt"
	"os"
)

type User struct {
	Name string
	ID   int
	Role string
}

func (u *User) CreateUser(name string, id int, role string) {
	u.Name = name
	u.ID = id
	u.Role = role
}

func (u *User) UserRow() string {
	return fmt.Sprintf("%-25v %-2v %-10v\n", u.Name, u.ID, u.Role)
}

func (u *User) SaveUser() {
	f, err := os.OpenFile("../../user/users.txt", os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.WriteString(u.UserRow())
	if err != nil {
		panic(err)
	}
}
