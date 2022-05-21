package structures

import (
	"api/constants"
	"errors"
)

type UserRecord struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"_"` // Ignore password from JSON parse
}

func (u UserRecord) ComparePassword(pass string) bool {
	return pass == u.Password
}

func (u *UserRecord) SetFirstName(firstName string) error {
	if len(firstName) <= 0 {
		return errors.New(constants.RequiredFirstName)
	}

	u.FirstName = firstName
	return nil
}

func (u *UserRecord) SetLastName(lastName string) error {
	if len(lastName) <= 0 {
		return errors.New(constants.RequiredLastName)
	}

	u.LastName = lastName
	return nil
}

func (u *UserRecord) SetEmail(email string) error {
	if len(email) <= 0 {
		return errors.New(constants.RequiredEmail)
	}

	u.Email = email
	return nil
}

func (u UserRecord) SetPassword(password string) error {
	if len(password) <= 0 {
		return errors.New(constants.RequiredPassword)
	}

	u.Password = password
	return nil
}
