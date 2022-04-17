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
	var body body
	var code int
	var response httpHelper.JSON

	err := context.ShouldBindJSON(&body)

	if err != nil {
		code, response = helpers.BadRequest([]error{err})
	}

	result, err := userModel.Create(structures.User{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
	})

	code, response = helpers.OkRequest(constants.RegisteredSuccess, result)

	if err != nil {
		code, response = helpers.BadRequest([]error{err})
	}

	context.JSON(code, response)
}

type body struct {
	FirstName string `form:"firstName" json:"firstName" binding:"required"`
	LastName  string `form:"lastName" json:"lastName" binding:"required"`
	Email     string `form:"email" json:"email" binding:"required"`
}
