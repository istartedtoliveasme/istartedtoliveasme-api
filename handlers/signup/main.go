package signup

import (
	"api/configs"
	"api/constants"
	"api/constants/jwt"
	userModel "api/database/models/user-model"
	"api/helpers/httpHelper"
	jwtHelper "api/helpers/jwt-helper"
	"api/helpers/responses"
	"github.com/gin-gonic/gin"
	"os"
)

func Handler(c *gin.Context) {
	var body Body
	var jwtClaims httpHelper.JSON

	_, session := configs.StartNeo4jDriver()
	defer session.Close()

	if bindError := c.ShouldBindJSON(&body); bindError != nil {
		c.AbortWithStatusJSON(responses.BadRequest(constants.RequiredMissingFields, responses.BindError{
			Message: constants.FailedToBindRequestBody,
			Err:     bindError,
		}))
	}

	createUserProps, err := createUserFactory(session, body)
	if err != nil {
		c.AbortWithStatusJSON(responses.BadRequest(err.Error(), err))
	}

	userRecord, err := userModel.Create(createUserProps)
	if err != nil {
		c.AbortWithStatusJSON(responses.BadRequest(constants.ExistEmail, err))
	}

	serializedRecord, err := getRecordSerializer(userRecord)
	err = httpHelper.JSONParse(serializedRecord, &jwtClaims)
	if err != nil {
		c.AbortWithStatusJSON(responses.BadRequest(constants.FailedParseClaim, err))
	}

	claim := jwtHelper.JWTClaim(jwtClaims)

	accessToken, err := claim.SignClaim([]byte(os.Getenv(jwt.JwtSecret)))
	if err != nil {
		c.AbortWithStatusJSON(responses.BadRequest(constants.FailedParseClaim, err))
	}

	if !c.IsAborted() && err != nil {

		c.AbortWithStatusJSON(responses.BadRequest(constants.ExistEmail, err))
	}

	if !c.IsAborted() {
		c.JSON(responses.OkRequest(constants.RegisteredSuccess, httpHelper.JSON{
			"accessToken": accessToken,
			"profile":     serializedRecord,
		}))
	}
}
