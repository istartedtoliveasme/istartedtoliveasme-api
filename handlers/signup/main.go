package signup

import (
	"api/configs"
	"api/constants"
	"api/constants/jwt"
	userModel "api/database/models/user-model"
	jwtHelper "api/helpers/jwt-helper"
	"api/helpers/responses"
	"api/helpers/serializers"
	helperTypes "api/helpers/typings"
	"github.com/gin-gonic/gin"
	"net/mail"
	"os"
)

func Handler(c *gin.Context) {
	httpResponse := responses.HttpResponse[jwtHelper.JWTClaim]{
		Message: constants.RegisteredSuccess,
	}
	var body Body
	var jsonUserResponse helperTypes.Json[serializers.UserResponse]
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

	if _, emailError := mail.ParseAddress(body.Email); emailError != nil {
		httpResponse.Message = constants.InvalidEmail
		httpResponse.Err = responses.BindError{
			Message: constants.InvalidEmail,
			Err:     emailError,
		}
	}

	user.FirstName = body.FirstName
	user.LastName = body.LastName
	user.Email = body.Email
	user.Password = body.Password
	err := user.Create(body.CreateUserFactory(session))
	if err != nil {
		httpResponse.Err = err
		httpResponse.Message = constants.ExistEmail
		c.AbortWithStatusJSON(httpResponse.BadRequest())
	}

	err = jsonUserResponse.ParsePayload(user)
	if !c.IsAborted() && err != nil {
		httpResponse.Message = constants.FailedParseClaim
		httpResponse.Err = err
		c.AbortWithStatusJSON(httpResponse.BadRequest())
	}

	profile := jwtHelper.JWTClaim{"profile": jsonUserResponse.Payload}
	switch accessToken, err := profile.SignClaim([]byte(os.Getenv(jwt.JwtSecret))); err != nil {
	case true:
		httpResponse.Message = constants.FailedParseClaim
		httpResponse.Err = err

	default:
		httpResponse.Payload = jwtHelper.JWTClaim{
			"accessToken": accessToken,
			"profile":     jsonUserResponse.Payload,
		}
	}

	switch c.IsAborted() || httpResponse.Err != nil {
	case true:
		c.AbortWithStatusJSON(httpResponse.BadRequest())
	default:

		c.JSON(httpResponse.OkRequest())
	}
}
