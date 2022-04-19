package helpers

import (
	"api/helpers/httpHelper"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func GetSingleNodeProps(record neo4j.Record) httpHelper.JSON {
	if len(record.Values) > 0 {
		return record.Values[0].(neo4j.Node).Props
	}
	return nil
}
