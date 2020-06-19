package model

import "fmt"

type User struct {
	ID string
	Email string
}

func (u User) String() string {
	return fmt.Sprintf("User's id : %s, email : %s", u.ID, u.Email)
}

func NewUser(id, email string) *User {
	return & User{
		ID: id,
		Email: email,
	}
}

func (u *User) GetID() string{
	return u.ID
}

func (u * User) GetEmail() string{
	return u.Email
}
