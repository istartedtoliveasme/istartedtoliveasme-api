package mood

import (
	"api/configs"
	"api/constants"
	userModel "api/database/models/user-model"
	"api/handlers/mood/typings"
	"api/helpers/httpHelper"
	"api/helpers/responses"
	"github.com/gin-gonic/gin"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func CreateHandler(c *gin.Context) {
	var body typings.Body
	var header httpHelper.Header
	_, session := configs.StartNeo4jDriver()
	defer session.Close()

	bindError := c.ShouldBind(&body)

	if bindError != nil {
		c.AbortWithStatusJSON(responses.BadRequest(constants.FailedCreateMood, responses.BindError{
			Message: constants.FailedBindParams,
			Err:     bindError,
		}))
	}

	bindError = c.ShouldBindHeader(&header)
	if bindError != nil {
		c.AbortWithStatusJSON(responses.BadRequest(constants.FailedCreateMood, responses.BindError{
			Message: constants.FailedToBindRequestBody,
			Err:     bindError,
		}))
	}

	mood := createMoodRecord(body)

	getByEmailProps := userModel.GetByEmailProps{
		GetEmail: func() string {
			return "istartedtoliveasme@gmail.com"
		},
		GetSession: func() neo4j.Session {
			return session
		},
	}
	userRecord, err := userModel.GetByEmail(getByEmailProps)
	if err != nil {
		c.AbortWithStatusJSON(responses.BadRequest(constants.GetRecordFailed, err))
	}

	moodRecord, err := mood.Create(CreateMoodPropertyFactory(session, userRecord))

	if !c.IsAborted() && err != nil {
		c.AbortWithStatusJSON(responses.BadRequest(constants.FailedCreateMood, err))
	}

	if !c.IsAborted() {
		c.JSON(responses.OkRequest(constants.Success, moodRecord))
	}
}
