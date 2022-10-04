package main

import (
	"encoding/json"
	"reflect"
)

type esFoo TestType

var useESTags = false

func (f TestType) MarshalWithESTag() ([]byte, error) {
	useESTags = true
	data, err := json.Marshal(f)
	useESTags = false
	return data, err
}

func (f TestType) MarshalJSONB() ([]byte, error) {
	es := esFoo(f)
	_json, err := json.Marshal(es)
	if useESTags {
		if err != nil {
			goto end
		}
		var intf interface{}
		err = json.Unmarshal(_json, &intf)
		if err != nil {
			goto end
		}
		m := intf.(map[string]interface{})
		_m := make(map[string]interface{}, len(m))
		t := reflect.TypeOf(f)
		i := 0
		for _, v := range m {
			tag, found := t.Field(i).Tag.Lookup("es")
			if !found {
				continue
			}
			_m[tag] = v
			i++
		}
		_json, err = json.Marshal(_m)
	}
end:
	return _json, err
}
