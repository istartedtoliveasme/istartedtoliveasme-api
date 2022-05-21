package handlers

import (
	"api/helpers/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, types.Json{
		"message": "pong",
	})
}
