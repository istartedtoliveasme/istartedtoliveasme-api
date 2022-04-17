package structures

type User struct {
	Id        int
	FirstName string
	LastName  string
	Email     string
	Password  string `json:"-"` // ignoring property
}

func (u User) ComparePassword(pass string) bool {
	return pass == u.Password
}
