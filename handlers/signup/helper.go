package signup

import (
	userModel "api/database/models/user-model"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type Body struct {
	FirstName string `form:"firstName" json:"firstName" binding:"required"`
	LastName  string `form:"lastName" json:"lastName" binding:"required"`
	Email     string `form:"email" json:"email" binding:"required"`
	Password  string `form:"password" json:"password" binding:"required"`
}

func (body *Body) CreateUserFactory(s neo4j.Session) userModel.CreateProps {
	getSession := func() neo4j.Session {
		return s
	}
	return userModel.CreateProps{
		GetSession: getSession,
	}
}
