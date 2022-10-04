package main

import (
	"encoding/json"
	"reflect"
	"strings"
)

type customFoo TestType

var useCustomTag = false

func (f TestType) MarshalWithCustomTag() ([]byte, error) {
	useCustomTag = true
	data, err := json.Marshal(f)
	useCustomTag = false
	return data, err
}

func (f TestType) MarshalJSONA() ([]byte, error) {
	custom := customFoo(f)
	t, numFields := reflect.TypeOf(f), reflect.TypeOf(f).NumField()
	_json, err := json.Marshal(custom)
	if useCustomTag {
		if err != nil {
			return _json, err
		}
		var placeholder interface{}
		err = json.Unmarshal(_json, &placeholder)
		if err != nil {
			return _json, err
		}
		original := placeholder.(map[string]interface{})
		result := make(map[string]interface{}, numFields)
		for k, v := range original {
			for i := 0; i < numFields; i++ {
				jsonTag, found := t.Field(i).Tag.Lookup("json")
				if strings.Contains(jsonTag, ",") {
					jsonTag = strings.Split(jsonTag, ",")[0]
				}
				if found && jsonTag == k {
					customTag, found := t.Field(i).Tag.Lookup("custom")
					if found {
						result[customTag] = v
						break
					}
				}
			}
		}
		_json, err = json.Marshal(result)
	}
	return _json, err
}
