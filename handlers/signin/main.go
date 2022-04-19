package signin

import (
	"api/configs"
	"api/constants"
	"api/database/models"
	"api/helpers/httpHelper"
	"api/helpers/responses"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func Handler(context *gin.Context) {
	var body Body
	var code int
	var response = httpHelper.JSON{}
	_, session := configs.Neo4jDriver()

	err := context.ShouldBindJSON(&body)
	if err != nil {
		code, response = responses.BadRequest(constants.FailedAuthentication, []error{err})
	}

	code, response = signInUser(session, body, response)

	context.JSON(code, response)
}

func signInUser(session neo4j.Session, body Body, response httpHelper.JSON) (int, httpHelper.JSON) {
	record, err := userModel.GetByEmail(session)(body.Email)
	if err != nil {
		return responses.BadRequest(constants.FailedAuthentication, []error{err})
	}

	singleRecord, err := record.Single()
	if singleRecord == nil || err != nil {
		return responses.BadRequest(constants.FailedAuthentication, []error{err})
	}

	// Generate Access Token
	accessToken, err := json.Marshal(response)

	if err != nil {
		return responses.BadRequest(constants.FailedAuthentication, []error{err})
	}

	// TODO :: bind profile to a struct that hides field password
	return responses.OkRequest(constants.Success, httpHelper.JSON{
		"accessToken": accessToken,
		"profile":     httpHelper.GetJsonKey(response, "data"),
	})
}

type Body struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
