package userModel

import (
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
	if username != "istartedtoliveasme" {
		return structures.User{}, errors.New("Invalid username!")
	}
	return structures.User{
		Id:        rand.Int(),
		FirstName: "John",
		LastName:  "Doe",
		Email:     fmt.Sprintf("%s@gmail.com", username),
		Password:  "123456",
	}, nil
}
