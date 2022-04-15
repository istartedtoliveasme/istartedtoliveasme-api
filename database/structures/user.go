package structures

type User struct {
	Id        int
	FirstName string
	LastName  string
	Email     string
	Password  string
}

func (u User) ComparePassword(pass string) bool {
	return pass == u.Password
}
