package signup

import (
	"api/configs"
	"api/constants"
	"api/constants/jwt"
	userModel "api/database/models/user-model"
	"api/helpers/httpHelper"
	jwtHelper "api/helpers/jwt-helper"
	"api/helpers/responses"
	"github.com/gin-gonic/gin"
<<<<<<< HEAD
=======
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
>>>>>>> 8140b66 (Code improvements and update mod files)
	"os"
)

func Handler(c *gin.Context) {
	var body Body
	var jwtClaims httpHelper.JSON

	_, session := configs.StartNeo4jDriver()
<<<<<<< HEAD
	defer session.Close()
=======
	defer func(session neo4j.Session) {
		err := session.Close()
		if err != nil {
			c.AbortWithStatusJSON(responses.BadRequest(err.Error(), []error{err}))
		}
	}(session)
>>>>>>> 8140b66 (Code improvements and update mod files)

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(responses.BadRequest(constants.RequiredMissingFields, []error{err}))
	}

	createUserProps, err := createUserFactory(session, body)
	if err != nil {
		c.AbortWithStatusJSON(responses.BadRequest(err.Error(), []error{err}))
	}

	userRecord, err := userModel.Create(createUserProps)
	if err != nil {
		c.AbortWithStatusJSON(responses.BadRequest(constants.ExistEmail, []error{err}))
	}

	serializedRecord, err := getRecordSerializer(userRecord)
	err = httpHelper.JSONParse(serializedRecord, &jwtClaims)
	if err != nil {
		c.AbortWithStatusJSON(responses.BadRequest(constants.FailedParseClaim, []error{err}))
	}

	claim := jwtHelper.JWTClaim(jwtClaims)

	accessToken, err := claim.SignClaim([]byte(os.Getenv(jwt.JwtSecret)))
	if err != nil {
		c.AbortWithStatusJSON(responses.BadRequest(constants.FailedParseClaim, []error{err}))
	}

	if !c.IsAborted() && err != nil {

		c.AbortWithStatusJSON(responses.BadRequest(constants.ExistEmail, []error{err}))
	}

	if !c.IsAborted() {
		c.JSON(responses.OkRequest(constants.RegisteredSuccess, httpHelper.JSON{
			"accessToken": accessToken,
			"profile":     serializedRecord,
		}))
	}
}
