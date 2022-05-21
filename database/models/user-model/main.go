package userModel

import (
	"api/constants"
	"api/database/models/typings"
	"api/helpers"
	helperTypes "api/helpers/typings"
	"encoding/json"
	"errors"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"math/rand"
	"strconv"
)

type User struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"_"` // Ignore password from Json parse
}

type UserError struct {
	Message string
	Err     error
}

func (e UserError) Error() string {
	return e.Message
}

func (e UserError) Unwrap() error {
	return e.Err
}

func (user User) ComparePassword(pass string) bool {
	return pass == user.Password
}

func (user *User) SetFirstName(firstName string) error {
	if len(firstName) <= 0 {
		return errors.New(constants.RequiredFirstName)
	}

	user.FirstName = firstName
	return nil
}

func (user *User) SetLastName(lastName string) error {
	if len(lastName) <= 0 {
		return errors.New(constants.RequiredLastName)
	}

	user.LastName = lastName
	return nil
}

func (user *User) SetEmail(email string) error {
	if len(email) <= 0 {
		return errors.New(constants.RequiredEmail)
	}

	user.Email = email
	return nil
}

func (user User) SetPassword(password string) error {
	if len(password) <= 0 {
		return errors.New(constants.RequiredPassword)
	}

	user.Password = password
	return nil
}

type GetByIdCypher map[string]interface{}
type GetByIdProps struct {
	typings.GetSession
	Id string
}

func (user *User) GetById(props GetByIdProps) helperTypes.CustomError {
	cypher := "MATCH (u:User { id: $id }) RETURN u LIMIT 1"
	params := GetByIdCypher{"id": props.Id}

	result, err := props.GetSession().Run(cypher, params)

	if err != nil {
		return UserError{
			Message: constants.FailedCreateRecord,
			Err:     err,
		}
	}

	record, err := helpers.GetSingleRecord(result)

	if err != nil {
		return UserError{
			Message: constants.FailedFetchRecord,
			Err:     err,
		}
	}

	userJson := helperTypes.Json[User]{
		Payload: *user,
	}
	if err = userJson.ParsePayload(record); err != nil {
		return UserError{
			Message: constants.FailedEncodeRecord,
			Err:     err,
		}
	}

	return nil
}

type GetByEmailCypher map[string]interface{}
type GetByEmailProps struct {
	typings.GetSession
	Email string
}

func (user *User) GetByEmail(props GetByEmailProps) helperTypes.CustomError {
	cypher := "MATCH (u:User { email: $email }) RETURN u LIMIT 1"
	params := GetByEmailCypher{"email": props.Email}

	result, err := props.GetSession().Run(cypher, params)
	if err != nil {
		return UserError{
			Message: constants.GetRecordFailed,
			Err:     err,
		}
	}

	if result == nil {
		return UserError{
			Message: constants.NotFoundRecord,
		}
	}

	record, err := helpers.GetSingleRecord(result)
	if err != nil {
		return UserError{
			Message: constants.GetRecordFailed,
			Err:     err,
		}
	}

	userJson := helperTypes.Json[User]{
		Payload: *user,
	}
	if err = userJson.ParsePayload(record); err != nil {
		return UserError{
			Message: constants.FailedParseClaim,
			Err:     err,
		}
	}

	return nil
}

type UserCreateCypher map[string]interface{}
type CreateProps struct {
	typings.GetSession
}

func (user *User) Create(props CreateProps) helperTypes.CustomError {
	tx := props.GetSession()
	cypher := "MATCH (u:User {email:$email}) WHERE u IS NULL " +
		"CREATE (user:User { id: $id, firstName: $firstName, lastName: $lastName, email: $email, password: $password }) " +
		"RETURN user LIMIT 1"

	params := UserCreateCypher{
		"id":        strconv.Itoa(rand.Int()),
		"firstName": user.FirstName,
		"lastName":  user.LastName,
		"email":     user.Email,
		"password":  user.Password,
	}

	records, err := tx.Run(cypher, params)

	// if error is nil means the record exist
	if records == nil {
		return UserError{
			Message: constants.DuplicateRecord,
			Err:     err,
		}
	}

	if err != nil {
		return UserError{
			Message: constants.FailedCreateRecord,
			Err:     err,
		}
	}

	err = user.BindRecord(records)
	if err != nil {
		return UserError{
			Message: constants.FailedFetchRecord,
			Err:     err,
		}
	}

	return nil

}

func (user *User) BindRecord(result neo4j.Result) helperTypes.CustomError {
	singleRecord, err := helpers.GetSingleRecord(result)
	if err != nil {
		return UserError{
			Message: constants.GetRecordFailed,
		}
	}

	byteArray, err := json.Marshal(singleRecord)
	if err != nil {
		return UserError{
			Message: constants.FailedEncodeRecord,
			Err:     err,
		}
	}

	err = json.Unmarshal(byteArray, &user)
	if err != nil {
		return UserError{
			Message: constants.FailedDecodeRecord,
			Err:     err,
		}
	}
	return nil
}
