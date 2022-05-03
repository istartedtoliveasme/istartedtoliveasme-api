package signin

import (
	"api/constants"
	"api/constants/jwt"
	zeroValues "api/constants/zero-values"
	"api/database/models/typings"
	userModel "api/database/models/user-model"
	"api/database/structures"
	errorHelper "api/helpers/error-helper"
	"api/helpers/httpHelper"
	jwtHelper "api/helpers/jwt-helper"
	"api/helpers/serializers"
	"errors"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"os"
)

func signInUser(userRecord structures.UserRecord) (httpHelper.JSON, errorHelper.CustomError) {
	var claims httpHelper.JSON
	var emptyUserRecord structures.UserRecord
	var err errorHelper.CustomError
	var serializedRecord serializers.UserSerializer

	if userRecord == emptyUserRecord {
		return nil, typings.RecordError{
			Message: constants.GetRecordFailed,
			Err:     errors.New(constants.GetRecordFailed),
		}
	}

	err = httpHelper.JSONParse(userRecord, &claims)
	if err != nil {
		return nil, typings.RecordError{
			Message: constants.FailedSerializeRecord,
			Err:     err,
		}
	}

	accessToken, err := generateAccessToken(jwtHelper.JWTClaim(claims))
	if err != nil {
		return nil, err
	}

	err = httpHelper.JSONParse(userRecord, &serializedRecord)
	if err != nil {
		return nil, typings.RecordError{
			Message: constants.FailedSerializeRecord,
			Err:     err,
		}
	}

	return httpHelper.JSON{
		"accessToken": accessToken,
		"profile":     serializedRecord,
	}, nil

}

func generateAccessToken(claim jwtHelper.JWTClaim) (string, errorHelper.CustomError) {
	var secret = []byte(os.Getenv(jwt.JwtSecret))

	if claim == nil {
		return zeroValues.ZeroString, jwtHelper.JWTError{
			Message: constants.EmptyJWTClaim,
			Err:     nil,
		}
	}

	signClaim, err := claim.SignClaim(secret)
	if err != nil {
		return zeroValues.ZeroString, jwtHelper.JWTError{
			Message: constants.FailedSignClaim,
			Err:     err,
		}
	}
	return signClaim, nil
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
