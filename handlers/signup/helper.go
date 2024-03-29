package signup

import (
	"api/constants"
	"api/database/models/typings"
	userModel "api/database/models/user-model"
	"api/database/structures"
	"api/helpers/error-helper"
	"api/helpers/httpHelper"
	"api/helpers/serializers"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"math/rand"
	"strconv"
)

type Body struct {
	FirstName string `form:"firstName" json:"firstName" binding:"required"`
	LastName  string `form:"lastName" json:"lastName" binding:"required"`
	Email     string `form:"email" json:"email" binding:"required"`
	Password  string `form:"password" json:"password" binding:"required"`
}

func getRecordSerializer(userRecord structures.UserRecord) (serializers.UserSerializer, errorHelper.CustomError) {
	serializer := serializers.UserSerializer{}

	err := httpHelper.JSONParse(userRecord, &serializer)
	if err != nil {
		return serializer, typings.RecordError{
			Message: constants.FailedSerializeRecord,
			Err:     err,
		}
	}
	return serializer, nil

}

func createUserFactory(s neo4j.Session, body Body) (userModel.CreateProps, errorHelper.CustomError) {
	getSession := func() neo4j.Session {
		return s
	}

	run := func(cypherText string, params httpHelper.JSON) (neo4j.Result, error) {
		return getSession().Run(cypherText, params)
	}

	getUserData := func() (structures.UserRecord, error) {
		var userRecord structures.UserRecord
		props := userModel.GetByEmailProps{
			Run: run,
			GetEmail: func() string {
				return body.Email
			},
		}
		userRecord, err := userModel.GetByEmail(props)
		if err != nil {
			return userRecord, typings.RecordError{
				Message: constants.GetRecordFailed,
				Err:     err,
			}
		}

		return userRecord, nil
	}

	getUserInput := func() structures.UserRecord {
		return structures.UserRecord{
			Id:        strconv.Itoa(rand.Int()),
			FirstName: body.FirstName,
			LastName:  body.LastName,
			Email:     body.Email,
			Password:  body.Password,
		}
	}

	return userModel.CreateProps{
		GetUserData:  getUserData,
		GetUserInput: getUserInput,
		Run:          run,
	}, nil
}
