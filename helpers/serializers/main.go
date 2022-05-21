package serializers

import (
	"api/constants"
	"api/database/models/typings"
	userModel "api/database/models/user-model"
	"api/helpers/types"
)

type UserResponse struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

func (userResponse *UserResponse) GetRecord(user userModel.User) types.CustomError {
	jsonUserResponse := types.Json[UserResponse]{
		Payload: *userResponse,
	}
	err := jsonUserResponse.Parse(user)
	if err != nil {
		return typings.RecordError{
			Message: constants.FailedSerializeRecord,
			Err:     err,
		}
	}
	return nil

}
