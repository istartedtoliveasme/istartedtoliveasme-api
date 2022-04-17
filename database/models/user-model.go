package userModel

import (
	"api/constants"
	"api/database/structures"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
)

type UserModel interface {
	GetByUserName(username string) (structures.User, error)
	Create(user structures.User) (structures.User, error)
}

// GetByUserName :: Creating a mock response
func GetByUserName(username string) (structures.User, error) {
	if message := fmt.Sprintf("%s!", constants.InvalidUserName); username != "istartedtoliveasme" {
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

func Create(user structures.User) (structures.User, error) {
	if user.Email == "istartedtoliveasme@gmail.com" {
		return structures.User{}, errors.New(constants.ExistUserName)
	}
	return structures.User{
		Id:        rand.Int(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  strconv.Itoa(rand.Int()),
	}, nil
}
