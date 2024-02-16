package mood

import (
	"api/configs"
	"api/constants"
	moodModel "api/database/models/mood-model"
	userModel "api/database/models/user-model"
	"api/handlers/mood/typings"
	"api/helpers/httpHelper"
	"api/helpers/responses"
	"github.com/gin-gonic/gin"
<<<<<<< HEAD
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
=======
>>>>>>> 8140b66 (Code improvements and update mod files)
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

	record := createMoodRecord(body)

	getByEmailProps := userModel.GetByEmailProps{
		GetEmail: func() string {
			return "istartedtoliveasme@gmail.com"
		},
<<<<<<< HEAD
		GetSession: func() neo4j.Session {
			return session
		},
=======
>>>>>>> 8140b66 (Code improvements and update mod files)
	}
	userRecord, err := userModel.GetByEmail(getByEmailProps)
	if err != nil {
		c.AbortWithStatusJSON(responses.BadRequest(constants.GetRecordFailed, []error{err}))
	}

	moodRecord, err := moodModel.CreateMood(CreateMoodPropertyFactory(session, record, userRecord))

	if !c.IsAborted() && err != nil {
		c.AbortWithStatusJSON(responses.BadRequest(constants.FailedCreateMood, []error{err}))
	}

	if !c.IsAborted() {
		c.JSON(responses.OkRequest(constants.Success, moodRecord))
	}
}
