package Ijsonparser

type IJsonParser interface {
	EncodeJson(v interface{}) (*[]byte, error)
	DecodeJson(message *[]byte, v interface{}) error
}
