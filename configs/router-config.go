package configs

import (
	"github.com/gin-gonic/gin"
)

func GetRouterConfig(router *gin.Engine) error {
	router.TrustedPlatform = gin.PlatformGoogleAppEngine
	err := router.SetTrustedProxies([]string{"0.0.0.0"})

<<<<<<< HEAD
	gin.SetMode(gin.TestMode)

=======
>>>>>>> 8140b66 (Code improvements and update mod files)
	return err
}
