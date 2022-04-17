package signup

import (
	userModel "api/database/models"
	"api/database/structures"
	"api/helpers"
	"api/helpers/httpHelper"
	"github.com/gin-gonic/gin"
)

func Handler(context *gin.Context)  {
	var body body
	var code int
	var response httpHelper.JSON

	err := context.ShouldBindJSON(&body)

	if err != nil {
		code, response = helpers.BadRequest([]error{err})
	}

	_, err = userModel.Create(structures.User{
		FirstName: body.firstName,
		LastName: body.lastName,
		Email: body.email,
	})

	if err != nil {
		code, response = helpers.BadRequest([]error{err})
	}

	context.JSON(code, response)
}

type body struct {
	firstName string `form:"firstName" json:"username" binding:"required"`
	lastName string `form:"lastName" json:"lastName" binding:"required"`
	email string `form:"email" json:"email" binding:"required"`
}