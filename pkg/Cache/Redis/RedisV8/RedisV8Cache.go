package RedisV8

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type redisCache struct {
	Client *redis.Client
}

var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

var RedisV8 = redisCache{
	Client: rdb,
}

func (r *redisCache) Set(key string, value *[]byte, expirationMinutes int32) (success bool, err error){

	expMinutes := time.Duration(expirationMinutes) * time.Minute
	 var result = rdb.Set(context.Background(),key, value, expMinutes)
	if result.Err() != nil{
		return false, result.Err()
	}
	return true,nil
}

func (r *redisCache) Get(key string) (value string, err error){

	 var result = rdb.Get(context.Background(),key)
	 if result.Err() != nil{
		 return "", err
	 }
	return result.Val(), nil
}

func (r *redisCache) Delete(key string) (success bool, err error){

	 var result = rdb.Del(context.Background(),key)
	 if result.Err() != nil{
		 return false, err
	 }
	return true, nil
}


func (r *redisCache) GetHash(key string) (*map[string]string, error) {

	result := r.Client.HGetAll(context.Background(),key)
	if result.Err() != nil {
		return nil, result.Err()
	}
	var resultValue = map[string]string{}
	resultValue = result.Val()
	return &resultValue, nil
}

func (r *redisCache) AddHash(key string, value *map[string]interface{}) (success bool, err error) {
	result := r.Client.HMSet(context.Background(),key, value)
	if result.Err() != nil {
		return false, result.Err()
	}
	return true, nil
}

func (r *redisCache) DeleteHashElement(key string, fields ...string) (success bool, err error) {
	result := r.Client.HDel(context.Background(),key, fields ...)
	if result.Err() != nil {
		return false, result.Err()
	}
	return true, nil
}

func (r *redisCache) DeleteHash(key string) (success bool, err error) {
	result := r.Client.Del(context.Background(),key)
	if result.Err() != nil {
		return false, result.Err()
	}
	return true, nil
}