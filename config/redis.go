package config

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
)

func InitRedis() *redis.Client {
    host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	db := os.Getenv("REDIS_DB")
	password := os.Getenv("REDIS_PASSWORD")
    
    addr := fmt.Sprintf("%s:%s", host, port)
    fmt.Println("addr:->", addr)
    dbInt, _ :=  strconv.Atoi(db)

    // Connect to Redis
    ctx := context.Background()
    client := redis.NewClient(&redis.Options{
        Addr:     addr,
        Password: password,
        DB:      dbInt,
    })

    // Ping the Redis server to check the connection
    pong, err := client.Ping(ctx).Result()
    if err != nil {
        fmt.Println("Error pinging Redis server:", err)
        return nil
    }
    fmt.Println("Redis successfully connected!:", pong)
    return  client
}