package handlers

import (
	"api/helpers/httpHelper"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, httpHelper.JSON{
		"message": "pong",
	})
}
