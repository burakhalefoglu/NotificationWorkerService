package gojson

import (
	"errors"
	"github.com/goccy/go-json"
)

type goJson struct {
}

func GoJsonConstructor() *goJson {
	return &goJson{}
}

func (g *goJson) EncodeJson(v interface{}) (*[]byte, error) {
	value, marshalErr := json.Marshal(&v)
	if marshalErr != nil {
		return nil, errors.New("Can not marshal Value")
	}
	return &value, nil
}

func (g *goJson) DecodeJson(message *[]byte, v interface{}) error {
	unmarshalErr := json.Unmarshal(*message, &v)
	if unmarshalErr != nil {
		return errors.New("Can not unmarshal JSON")
	}
	return nil
}