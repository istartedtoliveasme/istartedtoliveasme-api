package mood

import (
	"api/configs"
	"api/constants"
	moodModel "api/database/models/mood-model"
	"api/helpers/responses"
	"github.com/gin-gonic/gin"
	"math/rand"
	"strconv"
	"time"
)

func CreateHandler(c *gin.Context) {
	httpResponse := responses.HttpResponse[moodModel.MoodWithUserRecord]{
		Message: constants.Success,
	}
	var body Body
	_, session := configs.StartNeo4jDriver()
	defer session.Close()

	bindError := c.ShouldBind(&body)

	if bindError != nil {
		httpResponse.Err = responses.BindError{
			Message: constants.FailedBindParams,
			Err:     bindError,
		}
		c.AbortWithStatusJSON(httpResponse.BadRequest())
	}

	mood := moodModel.Mood{
		Id:          strconv.Itoa(rand.Int()),
		Icon:        body.Icon,
		Title:       body.Title,
		Description: body.Description,
		CreatedAt:   time.Now().UTC(),
	}

	httpResponse.Payload, httpResponse.Err = mood.Create(body.CreateMoodPropertyFactory(session))

	switch c.IsAborted() || httpResponse.Err != nil {
	case true:
		c.AbortWithStatusJSON(httpResponse.BadRequest())
	default:
		c.JSON(httpResponse.OkRequest())
	}
}
