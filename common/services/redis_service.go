package services

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)


type RedisService struct {
    client *redis.Client
}

// Create new RedisService
func NewRedisService(client *redis.Client) *RedisService {
    return &RedisService{client: client} 
}

// GET value from key
func (r *RedisService) Get(key string) *redis.StringCmd {
    return r.client.Get(context.Background(), key)
}

// SET key value with expiration
func (r *RedisService) Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
    return r.client.Set(context.Background(), key, value, expiration) 
}

// DEL keys from Redis
func (r *RedisService) Del(keys ...string) *redis.IntCmd {
    return r.client.Del(context.Background(), keys...)  
}

// Set key expiration
func (r *RedisService) Expire(key string, expiration time.Duration) *redis.BoolCmd {
    return r.client.Expire(context.Background(), key, expiration)
}