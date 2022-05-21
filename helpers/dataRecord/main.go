package dataRecord

import (
	"api/constants"
	errorHelper "api/helpers/error-helper"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type BindFunction func(node neo4j.Node) errorHelper.CustomError

type BindRecordValuesError struct {
	Message string
	Err     error
}

func (b BindRecordValuesError) Error() string {
	return b.Message
}

func (b BindRecordValuesError) Unwrap() error {
	return b.Err
}

type DataRecord neo4j.Record

func (record *DataRecord) BindNode(bindFunction BindFunction) errorHelper.CustomError {
	for _, recordValue := range record.Values {
		switch recordValue.(type) {
		case neo4j.Node:
			err := bindFunction(recordValue.(neo4j.Node))
			if err != nil {
				return BindRecordValuesError{
					Message: constants.FailedParseMood,
					Err:     err.Unwrap(),
				}
			}
		}

	}

	return nil
}

type DataRecordCollections []*neo4j.Record

func (collections *DataRecordCollections) Bind(bindFunction BindFunction) ([]interface{}, errorHelper.CustomError) {
	var items []interface{}

	for _, record := range *collections {
		var item interface{}
		databaseRecord := DataRecord(*record)
		err := databaseRecord.BindNode(bindFunction)
		if err != nil {
			return items, err
		}
		items = append(items, item)
	}

	return items, nil
}
