package signin

import (
	"api/configs"
	"api/constants"
	userModel "api/database/models/user-model"
	jwtHelper "api/helpers/jwt-helper"
	"api/helpers/responses"
	"github.com/gin-gonic/gin"
	"net/mail"
)

func Handler(c *gin.Context) {
	httpResponse := responses.HttpResponse[jwtHelper.JWTClaim]{
		Message: constants.Authorized,
	}
	var body Body
	_, session := configs.StartNeo4jDriver()
	defer session.Close()

	if bindError := c.ShouldBind(&body); bindError != nil {
		httpResponse.Message = constants.FailedAuthentication
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

	var user userModel.User
	if err := user.GetByEmail(body.GetByEmailFactory(session)); err != nil {
		httpResponse.Message = constants.FailedAuthentication
		httpResponse.Err = err
		c.AbortWithStatusJSON(httpResponse.BadRequest())
	}

	var userHandler = UserHandler(user)
	if payload, err := userHandler.SignIn(); err != nil {
		httpResponse.Payload = payload
		httpResponse.Err = err
		c.AbortWithStatusJSON(httpResponse.BadRequest())
	}

	if !c.IsAborted() {
		c.JSON(httpResponse.OkRequest())
	}
}
