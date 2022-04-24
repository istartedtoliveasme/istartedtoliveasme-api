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

func GetAllNodeProps(result neo4j.Result) ([]httpHelper.JSON, error) {
	var payload []httpHelper.JSON

	collections, err := result.Collect()

	if err !=  nil {
		return nil, err
	}

	for _, record := range collections {
		payload = append(payload, GetSingleNodeProps(*record))
	}

	return payload, nil
}