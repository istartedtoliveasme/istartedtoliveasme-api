package userModel

import (
	"api/constants"
	"api/database/structures"
	"api/helpers/httpHelper"
	"errors"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"math/rand"
)

type UserModel interface {
	Cypher(transaction neo4j.Transaction) ModelCypherQuery
}

type ModelCypherQuery interface {
	GetByEmail(tx neo4j.Transaction) func(email string) (neo4j.Result, error)
	Create(tx neo4j.Transaction) func(user structures.User) (neo4j.Result, error)
}

func GetByEmail(tx neo4j.Session) func(email string) (neo4j.Result, error) {
	return func(email string) (neo4j.Result, error) {
		return tx.Run("MATCH (u:User { email: $email }) RETURN u LIMIT 1", httpHelper.JSON{"email": email})
	}
}

func Create(tx neo4j.Session) func(user structures.User) (neo4j.Result, error) {
	return func(user structures.User) (neo4j.Result, error) {
		getByEmailRecord, err := GetByEmail(tx)(user.Email)

		if err != nil {
			return nil, err
		}

		if getByEmailRecord.Next() {
			return nil, errors.New(constants.ExistUserName)
		}

		records, err := tx.Run("CREATE (u:User { id: $id, firstName: $firstName, lastName: $lastName, email: $email, password: $password }) RETURN u", httpHelper.JSON{
			"id":        rand.Int(),
			"firstName": user.FirstName,
			"lastName":  user.LastName,
			"email":     user.Email,
			"password":  user.Password,
		})

		if err != nil {
			return nil, err
		}

		return records, nil
	}
}
