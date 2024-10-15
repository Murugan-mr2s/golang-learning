package model

import (
	"encoding/json"
	"errors"
	"io"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var userList = []*User{
	&User{1, "Murugan", "muru2@gmail.com"},
	&User{2, "Siddhart", "sid@gmail.com"},
	&User{3, "Sanjana", "sajna@gmail.com"},
}

func (u *User) fromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(u)
}

type Users []*User

func (u *Users) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}

func GetUsers() Users {
	return userList
}

func GetUserById(id int) (*User, error) {
	for _, user := range userList {

		if user.ID == id {
			return user, nil
		}
	}
	return nil, errors.New("User not found")
}

func AddUser(u *User) {
	userList = append(userList, u)
}

func DeleteUserById(id int) error {
	for index, user := range userList {
		if user.ID == id {
			userList = append(userList[:index], userList[index+1:]...)
			return nil
		}
	}
	return errors.New("User not found")
}
