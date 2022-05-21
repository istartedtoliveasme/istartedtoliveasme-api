package signin

import (
	"api/configs"
	"api/constants"
	userModel "api/database/models/user-model"
	"api/helpers/httpHelper"
	"api/helpers/responses"
	"github.com/gin-gonic/gin"
)

func Handler(c *gin.Context) {
	var body Body
	var response httpHelper.JSON
	_, session := configs.StartNeo4jDriver()
	defer session.Close()

	bindError := c.ShouldBind(&body)
	if bindError != nil {
		c.AbortWithStatusJSON(responses.BadRequest(constants.FailedAuthentication, responses.BindError{
			Message: constants.FailedToBindRequestBody,
			Err:     bindError,
		}))
	}

	userRecord, err := userModel.GetByEmail(getByEmailFactory(session, body))

	if err != nil {
		c.AbortWithStatusJSON(responses.BadRequest(constants.FailedAuthentication, err))
	}

	response, err = signInUser(userRecord)

	if err != nil {
		c.AbortWithStatusJSON(responses.BadRequest(constants.FailedAuthentication, err))
	}

	if !c.IsAborted() && err != nil {
		c.AbortWithStatusJSON(responses.BadRequest(constants.ExistEmail, err))
	}

	if !c.IsAborted() {
		c.JSON(responses.OkRequest(constants.Authorized, response))
	}
}
