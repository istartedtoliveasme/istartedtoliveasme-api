package main

import (
	"api/routers"
	"github.com/gin-gonic/gin"
)

func getRouterConfig() *gin.Engine  {
	routerConfig := gin.Default()
	routerConfig.TrustedPlatform = gin.PlatformGoogleAppEngine
	routerConfig.SetTrustedProxies([]string{"http://localhost", "0.0.0.0"})
	gin.SetMode(gin.TestMode)

	return routerConfig
}

func main() {
	router := routers.GetV1Routers(getRouterConfig())

	router.Run(":1337") // listen and serve on 0.0.0.0:1337 (for windows "localhost:8080")
}
