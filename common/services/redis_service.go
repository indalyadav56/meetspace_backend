package services

import (
	"context"
	"encoding/json"
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

// Sets
func (r *RedisService) SAdd(key string, members ...interface{}) *redis.IntCmd {
    return r.client.SAdd(context.Background(), key, members)
}

// get the members of sets
func (r *RedisService) SMembers(key string) *redis.StringSliceCmd {
    return r.client.SMembers(context.Background(), key) 
}

// remove member from sets
func (r *RedisService) SRem(key string, members ...interface{}) *redis.IntCmd {
    return r.client.SRem(context.Background(), key, members)
}

// Lists
func (r *RedisService) LPush(key string, values ...interface{}) *redis.IntCmd {
    return r.client.LPush(context.Background(), key, values...)
} 

func (r *RedisService) LRange(key string, start, stop int64) *redis.StringSliceCmd {
    return r.client.LRange(context.Background(), key, start, stop) 
}

func (r *RedisService) LPop(key string) *redis.StringCmd {
return r.client.LPop(context.Background(), key)
}

// Hashes
func (r *RedisService) HSet(key, field, value string)  {
    r.client.HSet(context.Background(), key, field, value)
}
  
func (r *RedisService) HGetAll(key string) *redis.StringStringMapCmd {
    return r.client.HGetAll(context.Background(), key)
}

// Streams
func (r *RedisService) XAdd(stream string, values map[string]interface{})  {
    r.client.XAdd(context.Background(), &redis.XAddArgs{
       Stream: stream,
       Values: values,
    })
 }
 
func (r *RedisService) XRange(stream, start, stop string) *redis.XMessageSliceCmd {
   return r.client.XRange(context.Background(), stream, start, stop) 
}

// JSON
func (r *RedisService) SetJson(key string, value interface{}) {
    json, _ := json.Marshal(value) 
    r.client.Set(context.Background(), key, json, 0)
}
 
// Publish a message to a channel
func (r *RedisService) Publish(channel string, message interface{}) *redis.IntCmd {
    return r.client.Publish(context.Background(), channel, message)
}

// Subscribe to a channel and receive messages
func (r *RedisService) Subscribe(channel string) *redis.PubSub {
    return r.client.Subscribe(context.Background(), channel)
}

// Remove a channel subscription
func (r *RedisService) Unsubscribe(channel string, pubsub *redis.PubSub) error {
    return pubsub.Unsubscribe(context.Background(), channel)
}
