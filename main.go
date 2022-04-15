package main

import (
	"api/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.TrustedPlatform = gin.PlatformGoogleAppEngine
	router.SetTrustedProxies([]string{"http://localhost", "0.0.0.0"})
	gin.SetMode(gin.TestMode)

	v1 := router.Group("/v1")
	{
		v1.POST("/login", handlers.LoginHandler)
		v1.GET("/ping", handlers.PingHandler)
	}

	router.Run(":1337") // listen and serve on 0.0.0.0:1337 (for windows "localhost:8080")
}
