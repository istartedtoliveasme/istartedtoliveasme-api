package profile

import (
	"api/configs"
	"api/constants"
	userModel "api/database/models/user-model"
	"api/helpers/httpHelper"
	"api/helpers/responses"
	"api/helpers/serializers"
	"github.com/gin-gonic/gin"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func Handler(c *gin.Context) {
	var userSerialized serializers.UserSerializer
	var params Params
	_, session := configs.StartNeo4jDriver()
	defer session.Close()

	if bindError := c.ShouldBindUri(&params); bindError != nil {
		c.AbortWithStatusJSON(responses.BadRequest(constants.FailedBindParams, responses.BindError{
			Message: constants.FailedToBindRequestBody,
			Err:     bindError,
		}))
	}

	getByEmailProps := userModel.GetByEmailProps{
		GetSession: func() neo4j.Session {
			return session
		},
		GetEmail: func() string {
			return params.Email
		},
	}

	record, err := userModel.GetByEmail(getByEmailProps)
	if err != nil {
		c.AbortWithStatusJSON(responses.BadRequest(constants.NotFoundRecord, err))
	}

	if err = httpHelper.JSONParse(record, &userSerialized); err != nil {
		c.AbortWithStatusJSON(responses.BadRequest(constants.FailedSerializeRecord, err))
	}

	if !c.IsAborted() {
		c.JSON(responses.OkRequest(constants.Success, userSerialized))
	}
}

type Params struct {
	Email string `uri:"email" binding:"required"`
}
