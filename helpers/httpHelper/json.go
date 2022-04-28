package httpHelper

import "encoding/json"

type JSON map[string]interface{}

func GetJsonKey(json JSON, key string) interface{} {
	if json[key] != nil {
		return json[key]
	}
	return nil
}

func JSONParse(source interface{}, data any) error {
	byteArray, err := json.Marshal(source)

	if err != nil {
		return err
	}

	if err = json.Unmarshal(byteArray, data); err != nil {
		return err
	}

	return nil
}
