package mood

import (
	"api/configs"
	"api/constants"
	moodModel "api/database/models/mood-model"
	"api/handlers/mood/typings"
	"api/helpers/httpHelper"
	"api/helpers/responses"
	"fmt"
	"github.com/gin-gonic/gin"
)

func CreateHandler(c *gin.Context) {
	var body typings.Body
	var header httpHelper.Header
	_, session := configs.StartNeo4jDriver()
	defer session.Close()

	err := c.ShouldBind(&body)

	if err != nil {
		c.AbortWithStatusJSON(responses.BadRequest(constants.FailedCreateMood, []error{err}))
	}

	err = c.ShouldBindHeader(&header)
	if err != nil {
		c.AbortWithStatusJSON(responses.BadRequest(constants.FailedCreateMood, []error{err}))
	}

	decodedUserSerializer, err := header.DecodeAuthorizationBearer()
	if err != nil {
		c.AbortWithStatusJSON(responses.BadRequest(constants.FailedCreateMood, []error{err}))
	}

	fmt.Println(decodedUserSerializer)

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
