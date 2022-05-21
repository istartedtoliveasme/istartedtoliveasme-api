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

	urlPath := UrlPath(Version1)
	v1 := router.Group(string(urlPath))
	{
		v1.POST(urlPath.Slugs(SignIn), signin.Handler)
		v1.GET(urlPath.Slugs(Ping), handlers.PingHandler)
		v1.POST(urlPath.Slugs(SignUp), signup.Handler)
		v1.POST(urlPath.Slugs(Mood), mood.CreateHandler)
		v1.GET(urlPath.Slugs(Mood), mood.GetAllHandler)
		v1.GET(urlPath.Slugs(Profile, ":email"), profile.Handler)
	}
}
