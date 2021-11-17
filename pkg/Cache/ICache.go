package cache

type ICache interface {
	Set(key string, value *[]byte, expirationMinutes int32) (success bool, err error)
	Get(key string) (value string, err error)
	Delete(key string) (success bool, err error)
	GetHash(key string) (*map[string]string, error)
	AddHash(key string, value *map[string]interface{}) (success bool, err error)
	DeleteHashElement(key string, fields ...string) (success bool, err error)
	DeleteHash(key string) (success bool, err error)
}
