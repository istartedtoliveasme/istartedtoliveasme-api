package structures

import (
	"api/constants"
	"errors"
)

type User struct {
	Id        int
	FirstName string
	LastName  string
	Email     string
	Password  string `json:"-"` // ignoring property
}

func (u User) ComparePassword(pass string) bool {
	return pass == u.Password
}

func (u *User) SetFirstName(firstName string) error {
	if len(firstName) <= 0 {
		return errors.New(constants.RequiredFirstName)
	}

	u.FirstName = firstName
	return nil
}

func (u *User) SetLastName(lastName string) error {
	if len(lastName) <= 0 {
		return errors.New(constants.RequiredLastName)
	}

	u.LastName = lastName
	return nil
}

func (u *User) SetEmail(email string) error {
	if len(email) <= 0 {
		return errors.New(constants.RequiredEmail)
	}

	u.Email = email
	return nil
}

func (u User) SetPassword(password string) error {
	if len(password) <= 0 {
		return errors.New(constants.RequiredPassword)
	}

	u.Password = password
	return nil
}
