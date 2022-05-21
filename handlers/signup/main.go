package signup

import (
	"api/configs"
	"api/constants"
	"api/constants/jwt"
	userModel "api/database/models/user-model"
	jwtHelper "api/helpers/jwt-helper"
	"api/helpers/responses"
	"api/helpers/serializers"
	"api/helpers/types"
	"github.com/gin-gonic/gin"
	"os"
)

func Handler(c *gin.Context) {
	httpResponse := responses.HttpResponse{
		Message: constants.RegisteredSuccess,
	}
	var body Body
	var jwtClaims types.Json[jwtHelper.JWTClaim]
	var user userModel.User

	_, session := configs.StartNeo4jDriver()
	defer session.Close()

	if bindError := c.ShouldBindJSON(&body); bindError != nil {
		httpResponse.Message = constants.RequiredMissingFields
		httpResponse.Err = responses.BindError{
			Message: constants.FailedToBindRequestBody,
			Err:     bindError,
		}
		c.AbortWithStatusJSON(httpResponse.BadRequest())
	}

	err := user.Create(body.CreateUserFactory(session))
	if httpResponse.Err != nil {
		httpResponse.Err = err
		httpResponse.Message = constants.ExistEmail
		c.AbortWithStatusJSON(httpResponse.BadRequest())
	}

	var userResponse serializers.UserResponse
	err = userResponse.GetRecord(user)
	if err != nil {
		httpResponse.Message = constants.FailedSerializeRecord
		httpResponse.Err = err
		c.AbortWithStatusJSON(httpResponse.BadRequest())
	}

	err = jwtClaims.Parse(userResponse)
	if err != nil {
		httpResponse.Message = constants.FailedParseClaim
		httpResponse.Err = err
		c.AbortWithStatusJSON(httpResponse.BadRequest())
	}

	claim := jwtClaims.Payload

	accessToken, err := claim.SignClaim([]byte(os.Getenv(jwt.JwtSecret)))
	if err != nil {
		httpResponse.Message = constants.FailedParseClaim
		httpResponse.Err = err
		c.AbortWithStatusJSON(httpResponse.BadRequest())
	}

	switch c.IsAborted() || httpResponse.Err != nil {
	case true:
		c.AbortWithStatusJSON(httpResponse.BadRequest())
	default:
		httpResponse.Payload = jwtHelper.JWTClaim{
			"accessToken": accessToken,
			"profile":     userResponse,
		}
		c.JSON(httpResponse.OkRequest())
	}
}
