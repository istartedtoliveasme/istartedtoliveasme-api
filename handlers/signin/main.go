package signin

import (
	"api/configs"
	"api/constants"
	userModel "api/database/models/user-model"
	"api/helpers/responses"
	"github.com/gin-gonic/gin"
)

func Handler(c *gin.Context) {
	httpResponse := responses.HttpResponse{
		Message: constants.Authorized,
	}
	var body Body
	_, session := configs.StartNeo4jDriver()
	defer session.Close()

	bindError := c.ShouldBind(&body)
	if bindError != nil {
		httpResponse.Message = constants.FailedAuthentication
		httpResponse.Err = responses.BindError{
			Message: constants.FailedToBindRequestBody,
			Err:     bindError,
		}
		c.AbortWithStatusJSON(httpResponse.BadRequest())
	}

	var user userModel.User
	err := user.GetByEmail(body.GetByEmailFactory(session))
	if err != nil {
		httpResponse.Message = constants.FailedAuthentication
		httpResponse.Err = err
		c.AbortWithStatusJSON(httpResponse.BadRequest())
	}

	var userHandler = UserHandler(user)
	httpResponse.Payload, err = userHandler.SignIn()
	if err != nil {
		httpResponse.Err = err
	}

	switch httpResponse.Err != nil {
	case true:
		c.AbortWithStatusJSON(httpResponse.BadRequest())
	default:
		c.JSON(httpResponse.OkRequest())
	}
}
