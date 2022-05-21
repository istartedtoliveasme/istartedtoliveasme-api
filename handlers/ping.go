package handlers

import (
	helperTypes "api/helpers/typings"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, helperTypes.JsonPayload{
		"message": "pong",
	})
}
