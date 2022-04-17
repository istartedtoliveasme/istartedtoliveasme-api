package signup

import (
	"api/constants"
	userModel "api/database/models"
	"api/database/structures"
	"api/helpers"
	"api/helpers/httpHelper"
	"github.com/gin-gonic/gin"
)

func Handler(context *gin.Context) {
	var body Body
	var code int
	var response httpHelper.JSON

	if err := context.ShouldBindJSON(&body); err != nil {
		code, response = helpers.BadRequest([]error{err})
		context.JSON(code, response)
	} else {

		result, err := userModel.Create(structures.User{
			FirstName: body.FirstName,
			LastName: body.LastName,
			Email: body.Email,
		})

		code, response = helpers.OkRequest(constants.RegisteredSuccess, result)

		if err != nil {
			code, response = helpers.BadRequest([]error{err})
		}

		context.JSON(code, response)

	}
}

type Body struct {
	FirstName string `form:"firstName" json:"firstName" binding:"required"`
	LastName  string `form:"lastName" json:"lastName" binding:"required"`
	Email     string `form:"email" json:"email" binding:"required"`
}
