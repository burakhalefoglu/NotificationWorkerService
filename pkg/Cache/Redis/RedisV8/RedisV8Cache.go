package RedisV8

import (
	"NotificationWorkerService/pkg/helper"
	"context"
	"errors"
	"os"
	"time"

	"github.com/appneuroncompany/light-logger/clogger"
	"github.com/go-redis/redis/v8"
)

type redisCache struct {
	Client *redis.Client
}

func RedisCacheConstructor() *redisCache {
	return &redisCache{Client: createClient()}
}

func createClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     helper.ResolvePath("REDIS_HOST", "REDIS_PORT"),
		Password: os.Getenv("REDIS_PASS"),
		DB:       0,
	})
	func() {
		_, err := client.Ping(context.Background()).Result()
		if err != nil {
			clogger.Error(&map[string]interface{}{
				"redisCache RedisConnection ConnectRedis error: ": err,
			})
		}
	}()

	return client
}

func (r *redisCache) Set(key string, value *[]byte, expirationMinutes int32) (success bool, err error) {

	expMinutes := time.Duration(expirationMinutes) * time.Minute
	var result = r.Client.Set(context.Background(), key, value, expMinutes)
	if result.Err() != nil {
		clogger.Error(&map[string]interface{}{
			"redisCache Set error: ": result.Err(),
		})
		return false, result.Err()
	}
	return true, nil
}

func (r *redisCache) Get(key string) (value string, err error) {

	var result = r.Client.Get(context.Background(), key)
	if result.Err() == redis.Nil {
		return "", errors.New("null data error")
	}
	if result.Err() != nil && result.Err() != redis.Nil {
		clogger.Error(&map[string]interface{}{
			"redisCache Get error: ": result.Err(),
		})
		return "", err
	}
	return result.Val(), nil
}

func (r *redisCache) Delete(key string) (success bool, err error) {

	var result = r.Client.Del(context.Background(), key)
	if result.Err() != nil {
		clogger.Error(&map[string]interface{}{
			"redisCache Delete error: ": result.Err(),
		})
		return false, err
	}
	return true, nil
}

func (r *redisCache) GetHash(key string) (*map[string]string, error) {

	result := r.Client.HGetAll(context.Background(), key)
	if result.Err() == redis.Nil {
		return nil, errors.New("null data error")
	}
	if result.Err() != nil && result.Err() != redis.Nil {
		clogger.Error(&map[string]interface{}{
			"redisCache GetHash error: ": result.Err(),
		})
		return nil, result.Err()
	}
	var resultValue = map[string]string{}
	resultValue = result.Val()
	return &resultValue, nil
}

func (r *redisCache) AddHash(key string, value *map[string]interface{}) (success bool, err error) {
	result := r.Client.HMSet(context.Background(), key, value)
	if result.Err() != nil {
		clogger.Error(&map[string]interface{}{
			"redisCache AddHash error: ": result.Err(),
		})
		return false, result.Err()
	}
	return true, nil
}

func (r *redisCache) DeleteHashElement(key string, fields ...string) (success bool, err error) {
	result := r.Client.HDel(context.Background(), key, fields...)
	if result.Err() != nil {
		clogger.Error(&map[string]interface{}{
			"redisCache DeleteHashElement error: ": result.Err(),
		})
		return false, result.Err()
	}
	return true, nil
}

func (r *redisCache) DeleteHash(key string) (success bool, err error) {
	result := r.Client.Del(context.Background(), key)
	if result.Err() != nil {
		clogger.Error(&map[string]interface{}{
			"redisCache DeleteHash error: ": result.Err(),
		})
		return false, result.Err()
	}
	return true, nil
}
