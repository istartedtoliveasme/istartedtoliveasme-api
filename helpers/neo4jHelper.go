package helpers

import (
	"api/constants"
	"api/database/models/typings"
	"api/helpers/error-helper"
	"api/helpers/httpHelper"
	"errors"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func GetSingleRecord(result neo4j.Result) (httpHelper.JSON, error) {
	var data httpHelper.JSON

	record, err := result.Single()

	if err != nil {
		return nil, err
	}

	singleProps, err := GetSinglePropsByRecord(*record)

	if err != nil {
		return nil, err
	}

	if err = httpHelper.JSONParse(singleProps, &data); err != nil {
		return nil, err
	}

	return data, nil
}

func GetAllRecords(result neo4j.Result) ([]httpHelper.JSON, errorHelper.CustomError) {
	var payload []httpHelper.JSON

	collections, err := result.Collect()

	if err != nil {
		return nil, typings.RecordError{
			Message: constants.FailedFetchRecord,
			Err:     err,
		}
	}

	for _, record := range collections {
		var parsePayload httpHelper.JSON
		data, err := GetSinglePropsByRecord(*record)
		if parseErr := httpHelper.JSONParse(data, &parsePayload); err != nil || parseErr != nil {
			return payload, typings.RecordError{
				Message: constants.FailedFetchRecord,
				Err:     err,
			}
		}

		payload = append(payload, parsePayload)
	}

	return payload, nil
}

func GetSinglePropsByRecord(record neo4j.Record) (interface{}, errorHelper.CustomError) {

	if len(record.Values) > 0 {
		return record.Values[0].(neo4j.Node).Props, nil
	}

	return nil, typings.RecordError{
		Message: constants.GetRecordFailed,
		Err:     errors.New(constants.GetRecordFailed),
	}
}
