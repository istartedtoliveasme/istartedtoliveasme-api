package mood

import (
	"api/configs"
	"api/constants"
	moodModel "api/database/models/mood-model"
	"api/helpers/responses"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetAllHandler(c *gin.Context) {
	_, session := configs.Neo4jDriver()

	records, err := moodModel.GetMoods(session)

	if err != nil {
		fmt.Println("no error")
		c.AbortWithStatusJSON(responses.BadRequest(constants.GetRecordFailed, []error{err}))
	}

	if !c.IsAborted() {
		c.JSON(responses.OkRequest(constants.Success, records))
	}

}
