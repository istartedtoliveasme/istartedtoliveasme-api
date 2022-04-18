package signin

import (
	"api/configs"
	"api/constants"
	"api/database/models"
	"api/helpers"
	"api/helpers/httpHelper"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Handler(context *gin.Context) {
	var body Body
	var code int
	var response = httpHelper.JSON{}
	_, session := configs.Neo4jDriver()

	err := context.ShouldBindJSON(&body)

	if err != nil {
		code, response = helpers.BadRequest([]error{err})
	}

	record, err := userModel.GetByEmail(session)(body.Email)

	if record == nil {
		code = http.StatusBadRequest
		response = httpHelper.JSON{"message": constants.FailedAuthentication}
	}

	context.JSON(code, response)
}

type Body struct {
	Email string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
