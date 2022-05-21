package routers

import (
	"api/handlers"
	"api/handlers/mood"
	"api/handlers/profile"
	"api/handlers/signin"
	"api/handlers/signup"
	"github.com/gin-gonic/gin"
)

func GetV1Routers(router *gin.Engine) {

	v1 := router.Group(GetURLPath(Version1))
	{
		v1.POST(GetURLPath(SignIn), signin.Handler)
		v1.GET(GetURLPath(Ping), handlers.PingHandler)
		v1.POST(GetURLPath(SignUp), signup.Handler)
		v1.POST(GetURLPath(Mood), mood.CreateHandler)
		v1.GET(GetURLPath(Mood), mood.GetAllHandler)
		v1.GET(GetURLPath(Profile)+"/:email", profile.Handler)
	}
}
