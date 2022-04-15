package userModel

import (
	"api/constants"
	"api/database/structures"
	"errors"
	"fmt"
	"math/rand"
)

type UserModel interface {
	GetByUserName(username string) (structures.User, error)
}

// GetByUserName :: Creating a mock response
func GetByUserName(username string) (structures.User, error) {
	if message := fmt.Sprintf("%s!", constants.InvalidUsername); username != "istartedtoliveasme" {
		return structures.User{}, errors.New(message)
	}
	return structures.User{
		Id:        rand.Int(),
		FirstName: "John",
		LastName:  "Doe",
		Email:     fmt.Sprintf("%s@gmail.com", username),
		Password:  "123456",
	}, nil
}
