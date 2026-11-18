package user

import (
	"errors"
)

type User struct {
	firstname string
	lastname  string
}

func (u *User) Clear() {
	u.firstname = ""
	u.lastname = ""
}

func New(firstname, lastname string) (*User, error) {
	if firstname == "" || lastname == "" {
		return nil, errors.New("Error")
	}

	return &User{
		firstname: firstname,
		lastname:  lastname,
	}, nil
}
