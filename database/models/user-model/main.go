package userModel

import (
	"api/constants"
	"api/database/models/typings"
	"api/database/structures"
	"api/helpers"
	"api/helpers/error-helper"
	"api/helpers/httpHelper"
	"encoding/json"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type ModelCypherQuery interface {
	GetByEmail(tx neo4j.Transaction) func(email string) (neo4j.Result, error)
	Create(tx neo4j.Transaction) func(c CreateProps) (neo4j.Result, error)
}

func GetById(props GetByIdProps) (structures.UserRecord, errorHelper.CustomError) {
	var userRecord structures.UserRecord
	cypher := "MATCH (u:User { id: $id }) RETURN u LIMIT 1"
	params := httpHelper.JSON{"id": props.GetId()}

	result, err := props.GetSession().Run(cypher, params)

	if err != nil {
		return userRecord, typings.RecordError{
			Message: constants.FailedCreateRecord,
			Err:     err,
		}
	}

	record, err := helpers.GetSingleRecord(result)

	if err != nil {
		return userRecord, typings.RecordError{
			Message: constants.FailedFetchRecord,
			Err:     err,
		}
	}

	if err = httpHelper.JSONParse(record, &userRecord); err != nil {
		return userRecord, typings.RecordError{
			Message: constants.FailedEncodeRecord,
			Err:     err,
		}
	}

	return userRecord, nil
}

func GetByEmail(props GetByEmailProps) (structures.UserRecord, errorHelper.CustomError) {
	var userRecord structures.UserRecord
	cypher := "MATCH (u:User { email: $email }) RETURN u LIMIT 1"
	params := httpHelper.JSON{"email": props.GetEmail()}

	result, err := props.GetSession().Run(cypher, params)
	if err != nil {
		return userRecord, typings.RecordError{
			Message: constants.GetRecordFailed,
			Err:     err,
		}
	}

	if result == nil {
		return userRecord, typings.RecordError{
			Message: constants.FailedCreateRecord,
		}
	}

	record, err := helpers.GetSingleRecord(result)
	if err != nil {
		return userRecord, typings.RecordError{
			Message: constants.GetRecordFailed,
			Err:     err,
		}
	}

	if record == nil {
		return userRecord, nil
	}

	if err = httpHelper.JSONParse(record, &userRecord); err != nil {
		return userRecord, httpHelper.JSONParseError{
			Message: constants.FailedParseClaim,
			Err:     err,
		}
	}

	return userRecord, nil
}

func Create(props CreateProps) (structures.UserRecord, errorHelper.CustomError) {
	var userRecord structures.UserRecord
	var emptyRecord structures.UserRecord
	tx := props.GetSession()
	cypher := "CREATE (u:User { id: $id, firstName: $firstName, lastName: $lastName, email: $email, password: $password }) RETURN u LIMIT 1"

	userRecord, err := props.GetUserData()

	// if error is nil means the record exist
	if userRecord != emptyRecord && err == nil {
		return userRecord, typings.RecordError{
			Message: constants.DuplicateRecord,
			Err:     err,
		}
	}

	input := props.GetUserInput()
	params := httpHelper.JSON{
		"id":        input.Id,
		"firstName": input.FirstName,
		"lastName":  input.LastName,
		"email":     input.Email,
		"password":  input.Password,
	}

	records, err := tx.Run(cypher, params)

	if err != nil {
		return userRecord, typings.RecordError{
			Message: constants.FailedCreateRecord,
			Err:     err,
		}
	}

	userRecord, err = getUserSingleRecord(records)
	if err != nil {
		return userRecord, typings.RecordError{
			Message: constants.FailedFetchRecord,
			Err:     err,
		}
	}

	return userRecord, nil

}

func getUserSingleRecord(result neo4j.Result) (structures.UserRecord, errorHelper.CustomError) {
	var userRecord structures.UserRecord

	singleRecord, err := helpers.GetSingleRecord(result)
	if err != nil {
		return userRecord, typings.RecordError{
			Message: constants.GetRecordFailed,
		}
	}

	byteArray, err := json.Marshal(singleRecord)
	if err != nil {
		return userRecord, typings.RecordError{
			Message: constants.FailedEncodeRecord,
			Err:     err,
		}
	}

	err = json.Unmarshal(byteArray, &userRecord)
	if err != nil {
		return userRecord, typings.RecordError{
			Message: constants.FailedDecodeRecord,
			Err:     err,
		}
	}
	return userRecord, nil
}
