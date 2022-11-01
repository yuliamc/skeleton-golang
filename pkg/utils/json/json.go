package json

import (
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// Marshal delegate to jsoniter
func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// Unmarshal delegate to jsoniter
func Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

// MarshalIndent delegate to jsoniter
func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	return json.MarshalIndent(v, prefix, indent)
}

// Convert json interface into map string interface
func ToMapStringInterface(source *interface{}) *map[string]interface{} {
	if sourceParsed, err := json.Marshal(source); err == nil {
		resultMap := make(map[string]interface{})
		if err := json.Unmarshal(sourceParsed, &resultMap); err == nil {
			return &resultMap
		}
	}
	return nil
}
