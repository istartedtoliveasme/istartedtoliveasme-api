package routers

import "path"

const (
	Version1 = "v1"
	SignIn   = "signin"
	SignUp   = "signup"
	Ping     = "ping"
	Mood     = "moods"
	Profile  = "profile"
)

type UrlPath string

func (urlPath *UrlPath) Slugs(paths ...string) string {
	return path.Join(paths...)
}
