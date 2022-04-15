package configs

import "github.com/gin-gonic/gin"

func GetRouterConfig() *gin.Engine  {
	routerConfig := gin.Default()
	routerConfig.TrustedPlatform = gin.PlatformGoogleAppEngine
	routerConfig.SetTrustedProxies([]string{"http://localhost", "0.0.0.0"})
	gin.SetMode(gin.TestMode)

	return routerConfig
}
