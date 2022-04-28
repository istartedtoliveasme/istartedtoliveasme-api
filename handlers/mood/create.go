package mood

import (
	"api/configs"
	"api/constants"
	moodModel "api/database/models/mood-model"
	"api/helpers/responses"
	"github.com/gin-gonic/gin"
)

func CreateHandler(c *gin.Context) {
	var body Body
	_, session := configs.Neo4jDriver()

	err := c.ShouldBind(&body)

	if err != nil {
		c.AbortWithStatusJSON(responses.BadRequest(constants.FailedCreateMood, []error{err}))
	}

	record := createMoodRecord(body)

	_, err = moodModel.CreateMood(CreateMoodPropertyFactory(session, record))

	if !c.IsAborted() && err != nil {
		c.AbortWithStatusJSON(responses.BadRequest(constants.FailedCreateMood, []error{err}))
	}

	if !c.IsAborted() {
		c.JSON(responses.OkRequest(constants.Success, record))
	}
}
