package jsonParser

import (
	"errors"
	"log"

	"github.com/goccy/go-json"
)


func EncodeJson(v interface{}) ([]byte, bool) {
	value, marchalErr := json.Marshal(&v)
	if marchalErr != nil { 
		log.Printf("Can not marshal Value")
		return nil, true
	}
	return value, false
}

func DecodeJson(message []byte, v interface{}) (err error) {
	unmarchalErr := json.Unmarshal(message, &v)
	if unmarchalErr != nil {
		return errors.New("Can not unmarshal JSON")
		}
	return nil
}

