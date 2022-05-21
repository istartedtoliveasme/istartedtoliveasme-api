package profile

import (
	"api/configs"
	"api/constants"
	userModel "api/database/models/user-model"
	"api/helpers/responses"
	"api/helpers/serializers"
	helperTypes "api/helpers/typings"
	"github.com/gin-gonic/gin"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func Handler(c *gin.Context) {
	httpResponse := responses.HttpResponse[serializers.UserResponse]{
		Message: constants.Success,
	}
	var userSerialized serializers.UserResponse
	var params Params
	_, session := configs.StartNeo4jDriver()
	defer session.Close()

	if bindError := c.ShouldBindUri(&params); bindError != nil {
		httpResponse.Err = responses.BindError{
			Message: constants.FailedToBindRequestBody,
			Err:     bindError,
		}
		c.AbortWithStatusJSON(httpResponse.BadRequest())
	}

	var user userModel.User
	err := user.GetByEmail(params.GetByEmailController(session))
	if err != nil {
		httpResponse.Err = err
		c.AbortWithStatusJSON(httpResponse.BadRequest())
	}

	userJson := helperTypes.Json[serializers.UserResponse]{
		Payload: userSerialized,
	}
	if err = userJson.ParsePayload(user); err != nil {
		httpResponse.Err = err
		c.AbortWithStatusJSON(httpResponse.BadRequest())
	}

	httpResponse.Payload = userSerialized

	switch c.IsAborted() || httpResponse.Err != nil {
	case true:
		c.AbortWithStatusJSON(httpResponse.BadRequest())
	default:
		c.JSON(httpResponse.OkRequest())
	}
}

type Params struct {
	Email string `uri:"email" binding:"required"`
}

func (p Params) GetByEmailController(session neo4j.Session) userModel.GetByEmailProps {
	return userModel.GetByEmailProps{
		GetSession: func() neo4j.Session {
			return session
		},
		Email: p.Email,
	}
}
