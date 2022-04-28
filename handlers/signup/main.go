package signup

import (
	"api/configs"
	"api/constants"
	userModel "api/database/models/user-model"
	"api/database/structures"
	"api/helpers"
	"api/helpers/httpHelper"
	"api/helpers/responses"
	"api/serializers"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func Handler(context *gin.Context) {
	var body Body
	var code int
	var response httpHelper.JSON
	_, session := configs.Neo4jDriver()

	if err := context.ShouldBindJSON(&body); err != nil {
		code, response = responses.BadRequest(constants.ExistUserName, []error{err})
		context.JSON(code, response)
	} else {

		result, err := userModel.Create(session)(structures.User{
			FirstName: body.FirstName,
			LastName:  body.LastName,
			Email:     body.Email,
		})

		record, err := result.Single()

		if err != nil {
			code, response = responses.BadRequest(constants.ExistUserName, []error{err})
		}

		serializer := serializers.UserSerializer{}

		singleRecord := helpers.GetSingleNodeProps(*record)

		byteRecord, err := json.Marshal(singleRecord)

		if err != nil {
			code, response = responses.BadRequest(constants.ExistUserName, []error{err})
		}

		err = json.Unmarshal(byteRecord, &serializer)

		if err != nil {
			code, response = responses.BadRequest(constants.ExistUserName, []error{err})
		}

		code, response = responses.OkRequest(constants.RegisteredSuccess, serializer)

		err = session.Close()

		if err != nil {
			code, response = responses.BadRequest(constants.ExistUserName, []error{err})
		}

		context.JSON(code, response)

	}
}

type Body struct {
	FirstName string `form:"firstName" json:"firstName" binding:"required"`
	LastName  string `form:"lastName" json:"lastName" binding:"required"`
	Email     string `form:"email" json:"email" binding:"required"`
}
