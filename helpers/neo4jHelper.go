package helpers

import (
	"api/constants"
	"api/database/models/typings"
	helperTypes "api/helpers/typings"
	"errors"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type GetSingleRecordResponse map[string]interface{}

func GetSingleRecord(result neo4j.Result) (GetSingleRecordResponse, error) {
	var jsonResponse helperTypes.Json[GetSingleRecordResponse]

	record, err := result.Single()

	if err != nil {
		return nil, err
	}

	singleProps, err := GetSinglePropsByRecord(*record)

	if err != nil {
		return nil, err
	}

	if err = jsonResponse.ParsePayload(singleProps); err != nil {
		return nil, err
	}

	return jsonResponse.Payload, nil
}

func GetAllRecords(result neo4j.Result) ([]GetSingleRecordResponse, helperTypes.CustomError) {
	var allRecords helperTypes.Json[[]GetSingleRecordResponse]

	collections, err := result.Collect()

	if err != nil {
		return nil, typings.RecordError{
			Message: constants.FailedFetchRecord,
			Err:     err,
		}
	}

	for _, record := range collections {
		var singleRecord helperTypes.Json[GetSingleRecordResponse]
		data, err := GetSinglePropsByRecord(*record)
		if parseErr := singleRecord.ParsePayload(data); err != nil || parseErr != nil {
			return allRecords.Payload, typings.RecordError{
				Message: constants.FailedFetchRecord,
				Err:     err,
			}
		}

		allRecords.Payload = append(allRecords.Payload, singleRecord.Payload)
	}

	return allRecords.Payload, nil
}

func GetSinglePropsByRecord(record neo4j.Record) (interface{}, helperTypes.CustomError) {
	if len(record.Values) > 0 {
		for _, recordValue := range record.Values {
			fmt.Println("%T", recordValue)
			switch recordValue.(type) {
			case neo4j.Node:
				return recordValue.(neo4j.Node).Props, nil
			}
		}

	}

	return nil, typings.RecordError{
		Message: constants.GetRecordFailed,
		Err:     errors.New(constants.GetRecordFailed),
	}
}
