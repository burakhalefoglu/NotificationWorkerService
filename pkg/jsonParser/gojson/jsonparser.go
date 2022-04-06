package gojson

import (
	"errors"

	"github.com/appneuroncompany/light-logger/clogger"
	"github.com/goccy/go-json"
)

type goJson struct{}

func GoJsonConstructor() *goJson {
	return &goJson{}
}

func (g *goJson) EncodeJson(v interface{}) (*[]byte, error) {
	value, marshalErr := json.Marshal(&v)
	if marshalErr != nil {
		clogger.Error(&map[string]interface{}{"Can not marshal Value error: ": marshalErr})
		return nil, errors.New("Can not marshal Value")
	}
	return &value, nil
}

func (g *goJson) DecodeJson(message *[]byte, v interface{}) error {
	unmarshalErr := json.Unmarshal(*message, &v)
	if unmarshalErr != nil {
		clogger.Error(&map[string]interface{}{"Can not unmarshal JSON error: ": unmarshalErr})
		return errors.New("Can not unmarshal JSON")
	}
	return nil
}
