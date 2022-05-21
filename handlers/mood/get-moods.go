package mood

import (
	"api/configs"
	"api/constants"
	moodModel "api/database/models/mood-model"
	"api/helpers/responses"
	"github.com/gin-gonic/gin"
)

func GetAllHandler(c *gin.Context) {
	_, session := configs.StartNeo4jDriver()
	defer session.Close()

	records, err := moodModel.GetMoods(session)

	if err != nil {
		c.AbortWithStatusJSON(responses.BadRequest(constants.GetRecordFailed, err))
	}

	if !c.IsAborted() {
		c.JSON(responses.OkRequest(constants.Success, records))
	}

}
