package signin

import (
	"api/constants"
	userModel "api/database/models/user-model"
	"api/database/structures"
	"api/helpers/httpHelper"
	"api/serializers"
	"encoding/json"
	"errors"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func signInUser(userRecord structures.UserRecord) (httpHelper.JSON, error) {
	serializer := serializers.UserSerializer{}
	var emptyUserRecord structures.UserRecord

	if userRecord == emptyUserRecord {
		return nil, errors.New(constants.GetRecordFailed)
	}

	accessToken, err := generateAccessToken(userRecord)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(accessToken, &serializer); err != nil {
		return nil, err
	}

	return httpHelper.JSON{
		"accessToken": accessToken,
		"profile":     serializer,
	}, nil

}

func generateAccessToken(response interface{}) ([]byte, error) {
	return json.Marshal(response)
}

type Body struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func getByEmailFactory(s neo4j.Session, body Body) userModel.GetByEmailProps {
	getSession := func() neo4j.Session {
		return s
	}

	getEmail := func() string {
		return body.Email
	}

	return userModel.GetByEmailProps{
		GetSession: getSession,
		GetEmail:   getEmail,
	}
}
