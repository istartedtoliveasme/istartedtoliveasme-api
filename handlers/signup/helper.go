package signup

import (
	"api/constants"
	"api/database/models/typings"
	userModel "api/database/models/user-model"
	"api/database/structures"
	"api/helpers/error-helper"
	"api/serializers"
	"encoding/json"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"math/rand"
)

type Body struct {
	FirstName string `form:"firstName" json:"firstName" binding:"required"`
	LastName  string `form:"lastName" json:"lastName" binding:"required"`
	Email     string `form:"email" json:"email" binding:"required"`
	Password  string `form:"password" json:"password" binding:"required"`
}

func getRecordSerializer(userRecord structures.UserRecord) (serializers.UserSerializer, errorHelper.CustomError) {
	serializer := serializers.UserSerializer{}

	byteRecord, err := json.Marshal(userRecord)

	if err != nil {
		return serializer, typings.RecordError{
			Message: constants.FailedSerializeRecord,
			Err:     err,
		}
	}

	err = json.Unmarshal(byteRecord, &serializer)

	if err != nil {
		return serializer, typings.RecordError{
			Message: constants.FailedDecodeRecord,
			Err:     err,
		}
	}

	return serializer, nil

}

func createUserFactory(s neo4j.Session, body Body) userModel.CreateProps {
	getSession := func() neo4j.Session {
		return s
	}

	getUserData := func() (structures.UserRecord, error) {
		var userRecord structures.UserRecord
		props := userModel.GetByEmailProps{
			GetSession: getSession,
			GetEmail: func() string {
				return body.Email
			},
		}
		userRecord, err := userModel.GetByEmail(props)

		if err != nil {
			return userRecord, err
		}

		return userRecord, nil
	}

	getUserInput := func() structures.UserRecord {
		return structures.UserRecord{
			Id:        rand.Int(),
			FirstName: body.FirstName,
			LastName:  body.LastName,
			Email:     body.Email,
			Password:  body.Password,
		}
	}

	return userModel.CreateProps{
		GetSession:   getSession,
		GetUserData:  getUserData,
		GetUserInput: getUserInput,
	}
}
