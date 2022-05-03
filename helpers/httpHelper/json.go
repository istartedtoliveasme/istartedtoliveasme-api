package httpHelper

import (
	"api/constants"
	errorHelper "api/helpers/error-helper"
	"encoding/json"
)

type JSON map[string]interface{}

type JSONParseError struct {
	Message string
	Err     error
}

func (jsonP JSONParseError) Error() string {
	return jsonP.Message
}

func (jsonP JSONParseError) Unwrap() error {
	return jsonP.Err
}

func GetJsonKey(json JSON, key string) interface{} {
	if json[key] != nil {
		return json[key]
	}
	return nil
}

func JSONParse(source interface{}, data any) errorHelper.CustomError {
	byteArray, err := json.Marshal(source)

	if err != nil {
		return JSONParseError{
			Message: constants.FailedParseClaim,
			Err:     err,
		}
	}

	if err = json.Unmarshal(byteArray, data); err != nil {
		return JSONParseError{
			Message: constants.FailedDecodeRecord,
			Err:     err,
		}
	}

	return nil
}
