package mood

import (
	"api/configs"
	"api/constants"
	moodModel "api/database/models/mood-model"
	"api/helpers/httpHelper"
	"api/helpers/responses"
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
)

func Handler(context *gin.Context) {
	var body Body
	var code int
	var response httpHelper.JSON
	_, session := configs.Neo4jDriver()

	err := context.ShouldBind(&body)

	if err != nil {
		code, response = responses.BadRequest(constants.FailedCreateMood, []error{err})
	}

	record := moodModel.Mood{
		Id: rand.Int(),
		Icon: body.Icon,
		Title: body.Title,
		Description: body.Description,
		CreatedAt: time.Now().UTC(),
	}
	_, err = moodModel.CreateMood(session, record)

	if err != nil {
		code, response = responses.BadRequest(constants.FailedCreateMood, []error{err})
	}

	code, response = responses.OkRequest(constants.Success, record)

	context.JSON(code, response)
}

type Body struct {
	Icon        string `form:"icon" json:"icon" binding:"required"`
	Title       string `form:"title" json:"title" binding:"required"`
	Description string `form:"description" json:"description" binding:"required"`
}
