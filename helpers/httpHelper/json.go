package httpHelper

type JSON map[string]interface{}

func GetJsonKey(json JSON, key string) interface{} {
	if json[key] != nil {
		return json[key]
	}
	return nil
}
