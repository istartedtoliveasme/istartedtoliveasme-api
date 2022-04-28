package signup

import (
	"api/configs"
	"api/constants"
	userModel "api/database/models/user-model"
	"api/helpers/responses"
	"github.com/gin-gonic/gin"
)

func Handler(c *gin.Context) {
	var body Body
	_, session := configs.Neo4jDriver()

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(responses.BadRequest(constants.RequiredMissingFields, []error{err}))
	}

	userRecord, err := userModel.Create(createUserFactory(session, body))

	if !c.IsAborted() && err != nil {
		c.AbortWithStatusJSON(responses.BadRequest(constants.ExistUserName, []error{err}))
	}

	serializedRecord, err := getRecordSerializer(userRecord)

	if !c.IsAborted() && err != nil {
		c.AbortWithStatusJSON(responses.BadRequest(constants.ExistUserName, []error{err}))
	}

	if !c.IsAborted() {
		c.JSON(responses.OkRequest(constants.RegisteredSuccess, serializedRecord))
	}
}
