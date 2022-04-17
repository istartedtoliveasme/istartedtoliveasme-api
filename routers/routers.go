package routers

import (
	"api/handlers"
	"api/handlers/signin"
	"api/handlers/signup"
	"github.com/gin-gonic/gin"
)

func GetV1Routers(router *gin.Engine) {

	v1 := router.Group(GetURLPath(Version1))
	{
		v1.POST(GetURLPath(SignIn), signin.Handler)
		v1.GET(GetURLPath(Ping), handlers.PingHandler)
		v1.GET(GetURLPath(SignUp), signup.Handler)
	}
}
