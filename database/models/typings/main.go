package typings

import "github.com/neo4j/neo4j-go-driver/v4/neo4j"

type GetSession func() neo4j.Session

type RecordError struct {
	Message string
	Err     error
}

func (e RecordError) Error() string {
	return e.Message
}

func (e RecordError) Unwrap() error {
	return e.Err
}