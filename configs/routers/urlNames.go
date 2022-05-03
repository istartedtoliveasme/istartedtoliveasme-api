package routers

const (
	Version1 = "v1"
	SignIn   = "signin"
	SignUp = "signup"
	Ping = "ping"
	Mood = "moods"
)

func GetURLPath(routeName string) string {
	return "/" + routeName
}
