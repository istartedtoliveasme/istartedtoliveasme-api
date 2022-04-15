package singup

import (
	"api/database/models"
	"api/helpers"
	"api/helpers/httpHelper"
	"errors"
	"github.com/gin-gonic/gin"
)

func LoginHandler(context *gin.Context) {
	var body SignupBody
	var code int
	var response httpHelper.JSON

	ok := helpers.BindJSONBody(context)(&body)

	// golang ok idiom
	if ok {
		userData, err := userModel.GetByUserName(body.Username)

		ok := userData.ComparePassword(body.Password)

		switch true {
		case err != nil:
			code, response = helpers.BadRequest([]error{err})
		case ok == false:
			code, response = helpers.BadRequest([]error{errors.New("password doesn't match")})
		default:
			code, response = helpers.OkRequest("Successfully authorized!", userData)
		}

		context.JSON(code, response)
	}
}

type SignupBody struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
