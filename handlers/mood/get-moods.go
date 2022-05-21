package mood

import (
	"api/configs"
	"api/constants"
	moodModel "api/database/models/mood-model"
	"api/helpers/responses"
	"github.com/gin-gonic/gin"
)

func GetAllHandler(c *gin.Context) {
	httpResponse := responses.HttpResponse[[]moodModel.MoodWithUserRecord]{
		Message: constants.Success,
	}
	_, session := configs.StartNeo4jDriver()
	defer session.Close()

	var mood moodModel.Mood
	records, err := mood.GetMoods(session)
	httpResponse.Payload = records
	if err != nil {
		httpResponse.Err = err
	}

	switch c.IsAborted() || httpResponse.Err != nil {
	case true:
		c.AbortWithStatusJSON(httpResponse.BadRequest())
	default:
		c.JSON(httpResponse.OkRequest())
	}

}
