package routers

import (
	"api/handlers"
	"api/handlers/singup"
	"github.com/gin-gonic/gin"
)

func GetV1Routers(router *gin.Engine) *gin.Engine {

	v1 := router.Group("/v1")
	{
		v1.POST("/login", singup.LoginHandler)
		v1.GET("/ping", handlers.PingHandler)
	}

	return router
}
