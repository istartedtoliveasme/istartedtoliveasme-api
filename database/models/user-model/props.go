package userModel

import (
	"api/database/models/typings"
	"api/database/structures"
)

type CreateProps struct {
	typings.GetSession
	GetUserData  func() (structures.UserRecord, error)
	GetUserInput func() structures.UserRecord
}

type GetByEmailProps struct {
	typings.GetSession
	GetEmail func() string
}

type GetByIdProps struct {
	typings.GetSession
	GetId func() string
}
