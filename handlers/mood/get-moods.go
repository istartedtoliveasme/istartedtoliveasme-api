package mood

import (
	"api/configs"
	"api/constants"
	moodModel "api/database/models/mood-model"
	"api/helpers/httpHelper"
	"api/helpers/responses"
	"github.com/gin-gonic/gin"
)

func GetMoodHandler(context *gin.Context) {
	var code int
	var response httpHelper.JSON
	_, session := configs.Neo4jDriver()

	records, err := moodModel.GetMoods(session)

	if err != nil {
		code, response = responses.BadRequest(constants.GetRecordFailed, []error{err})
	}

	response = httpHelper.JSON{
		"message": constants.Success,
		"data": records,
	}

	context.JSON(code, response)
}
