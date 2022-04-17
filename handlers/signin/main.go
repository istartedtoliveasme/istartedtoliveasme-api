package signin

import (
	"api/constants"
	"api/database/models"
	"api/helpers"
	"api/helpers/httpHelper"
	"errors"
	"github.com/gin-gonic/gin"
)

func Handler(context *gin.Context) {
	var body Body
	var code int
	var response httpHelper.JSON

	err := context.ShouldBindJSON(&body)

	if err != nil {
		code, response = helpers.BadRequest([]error{err})
	}

	userData, err := userModel.GetByUserName(body.Username)

	ok := userData.ComparePassword(body.Password)

	switch true {
	case err != nil:
		code, response = helpers.BadRequest([]error{err})
	case ok == false:
		code, response = helpers.BadRequest([]error{errors.New(constants.MismatchPassword)})
	default:
		code, response = helpers.OkRequest(constants.Authorized, userData)
	}

	context.JSON(code, response)
}

type Body struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
