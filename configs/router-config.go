package configs

import (
	"github.com/gin-gonic/gin"
)

func GetRouterConfig(router *gin.Engine) error {
	router.TrustedPlatform = gin.PlatformGoogleAppEngine
	err := router.SetTrustedProxies([]string{"0.0.0.0"})

	gin.SetMode(gin.TestMode)

	return err
}
