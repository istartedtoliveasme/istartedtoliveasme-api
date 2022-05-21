package signin

import (
	"api/constants"
	"api/constants/jwt"
	zeroValues "api/constants/zero-values"
	"api/database/models/typings"
	userModel "api/database/models/user-model"
	jwtHelper "api/helpers/jwt-helper"
	"api/helpers/serializers"
	"api/helpers/types"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"os"
)

type UserHandler userModel.User

func (user *UserHandler) SignIn() (jwtHelper.JWTClaim, types.CustomError) {
	var jsonClaims types.Json[jwtHelper.JWTClaim]
	var err types.CustomError
	var serializedRecord serializers.UserResponse

	if err = jsonClaims.Parse(user); err != nil {
		return jsonClaims.Payload, typings.RecordError{
			Message: constants.FailedSerializeRecord,
			Err:     err,
		}
	}

	accessToken, err := user.GenerateAccessToken(jsonClaims.Payload)
	if err != nil {
		return jsonClaims.Payload, err
	}

	jsonUser := types.Json[serializers.UserResponse]{
		Payload: serializedRecord,
	}
	if err = jsonUser.Parse(user); err != nil {
		return nil, typings.RecordError{
			Message: constants.FailedSerializeRecord,
			Err:     err,
		}
	}

	return jwtHelper.JWTClaim{
		"accessToken": accessToken,
		"profile":     serializedRecord,
	}, nil

}

func (user *UserHandler) GenerateAccessToken(claim jwtHelper.JWTClaim) (string, types.CustomError) {
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

func (body *Body) GetByEmailFactory(s neo4j.Session) userModel.GetByEmailProps {
	getSession := func() neo4j.Session {
		return s
	}

	return userModel.GetByEmailProps{
		GetSession: getSession,
		Email:      body.Email,
	}
}
