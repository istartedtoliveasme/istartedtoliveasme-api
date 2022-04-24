package signin

import (
	"api/configs"
	"api/constants"
	"api/database/models/user-model"
	"api/helpers"
	"api/helpers/httpHelper"
	"api/helpers/responses"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func Handler(context *gin.Context) {
	var body Body
	var code int
	var response httpHelper.JSON
	_, session := configs.Neo4jDriver()

	err := context.ShouldBind(&body)
	if err != nil {
		code, response = responses.BadRequest(constants.FailedAuthentication, []error{err})
	}

	userRecord, err := userModel.GetByEmail(session)(body.Email)

	if err != nil {
		code, response = responses.BadRequest(constants.FailedAuthentication, []error{err})
	}

	response, err = signInUser(userRecord)

	if err != nil {
		code, response = responses.BadRequest(constants.FailedAuthentication, []error{err})
	}

	context.JSON(code, response)
}

func signInUser(userRecord neo4j.Result) (httpHelper.JSON, error) {
	singleRecord, err := getSingleRecord(userRecord)

	if err != nil {
		return nil, err
	}

	data := helpers.GetSingleNodeProps(*singleRecord)
	accessToken, err := generateAccessToken(data)

	if err != nil {
		return nil, err
	}

	// TODO :: bind profile to a struct that hides field password
	return httpHelper.JSON{
		"accessToken": accessToken,
		"profile":     data,
	}, nil

}

func getSingleRecord(userRecord neo4j.Result) (*neo4j.Record, error) {
	return userRecord.Single()
}

func generateAccessToken(response interface{}) ([]byte, error) {
	return json.Marshal(response)
}

type Body struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
