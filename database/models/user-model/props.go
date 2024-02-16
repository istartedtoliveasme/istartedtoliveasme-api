package userModel

import (
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
}
