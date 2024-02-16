package userModel

import (
<<<<<<< HEAD
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
=======
	"api/database/structures"
	"api/helpers/httpHelper"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type CreateProps struct {
	GetUserData  func() (structures.UserRecord, error)
	GetUserInput func() structures.UserRecord
	Run          func(cypherText string, params httpHelper.JSON) (neo4j.Result, error)
}

type GetByEmailProps struct {
	GetEmail func() string
	Run      func(cypherText string, params httpHelper.JSON) (neo4j.Result, error)
}

type GetByIdProps struct {
	GetId func() string
	Run   func(cypherText string, params httpHelper.JSON) (neo4j.Result, error)
>>>>>>> 8140b66 (Code improvements and update mod files)
}
