package RedisV8

import (
	"NotificationWorkerService/pkg/helper"
	"NotificationWorkerService/pkg/logger"
	"context"
	"github.com/go-redis/redis/v8"
	"os"
	"time"
)

type redisCache struct {
	Client *redis.Client
}

func RedisCacheConstructor(log *logger.ILog) *redisCache {
	return &redisCache{Client: getClient(log)}
}

func getClient(log *logger.ILog) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     helper.ResolvePath("REDIS_HOST", "REDIS_PORT"),
		Password: os.Getenv("REDIS_PASS"),
		DB:       0,
	})
	func(log *logger.ILog) {
		_, err := client.Ping(context.Background()).Result()
		if err != nil {
			(*log).SendPanicLog("RedisConnection", "ConnectRedis", err)
		}
	}(log)

	return client
}

func (r *redisCache) Set(key string, value *[]byte, expirationMinutes int32) (success bool, err error) {

	expMinutes := time.Duration(expirationMinutes) * time.Minute
	var result = r.Client.Set(context.Background(), key, value, expMinutes)
	if result.Err() != nil {
		return false, result.Err()
	}
	return true, nil
}

func (r *redisCache) Get(key string) (value string, err error) {

	var result = r.Client.Get(context.Background(), key)
	if result.Err() != nil {
		return "", err
	}
	return result.Val(), nil
}

func (r *redisCache) Delete(key string) (success bool, err error) {

	var result = r.Client.Del(context.Background(), key)
	if result.Err() != nil {
		return false, err
	}
	return true, nil
}

func (r *redisCache) GetHash(key string) (*map[string]string, error) {

	result := r.Client.HGetAll(context.Background(), key)
	if result.Err() != nil {
		return nil, result.Err()
	}
	var resultValue = map[string]string{}
	resultValue = result.Val()
	return &resultValue, nil
}

func (r *redisCache) AddHash(key string, value *map[string]interface{}) (success bool, err error) {
	result := r.Client.HMSet(context.Background(), key, value)
	if result.Err() != nil {
		return false, result.Err()
	}
	return true, nil
}

func (r *redisCache) DeleteHashElement(key string, fields ...string) (success bool, err error) {
	result := r.Client.HDel(context.Background(), key, fields...)
	if result.Err() != nil {
		return false, result.Err()
	}
	return true, nil
}

func (r *redisCache) DeleteHash(key string) (success bool, err error) {
	result := r.Client.Del(context.Background(), key)
	if result.Err() != nil {
		return false, result.Err()
	}
	return true, nil
}
