package helpers

import (
	httpHelper "api/helpers/httpHelper"
	"github.com/gin-gonic/gin"
	"net/http"
)

func BindJSONBody(context *gin.Context) func(s interface{}) bool {
	return func(s interface{}) bool {
		err := context.ShouldBindJSON(&s)

		if err != nil {
			context.JSON(http.StatusBadRequest, httpHelper.JSON{"errors": []string{err.Error()}})
			return false
		}
		return true
	}
}
