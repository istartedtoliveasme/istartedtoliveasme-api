package typings

import (
	"api/constants"
	jsonEncoding "encoding/json"
)

type JsonPayload map[string]interface{}
type Json[T any] struct {
	Payload T
}

type Error struct {
	Message string
	Err     error
}

func (jsonP Error) Error() string {
	return jsonP.Message
}

func (jsonP Error) Unwrap() error {
	return jsonP.Err
}

func (j *Json[T]) ParsePayload(data interface{}) CustomError {
	byteArray, err := jsonEncoding.Marshal(data)

	if err != nil {
		return Error{
			Message: constants.FailedParseClaim,
			Err:     err,
		}
	}

	if err = jsonEncoding.Unmarshal(byteArray, &j.Payload); err != nil {
		return Error{
			Message: constants.FailedDecodeRecord,
			Err:     err,
		}
	}

	return nil
}
