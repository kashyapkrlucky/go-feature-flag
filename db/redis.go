package db

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

// RedisClient is a global Redis instance
var RedisClient *redis.Client
var ctx = context.Background()

// InitRedis initializes the Redis client
func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: "", // No password for local dev
		DB:       0,  // Default DB
	})

	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	log.Println("Connected to Redis successfully!")
}

// SetCache sets a key-value pair in Redis
func SetCache(key string, value string, expiration time.Duration) error {
	return RedisClient.Set(ctx, key, value, expiration).Err()
}

// GetCache retrieves a value from Redis
func GetCache(key string) (string, error) {
	return RedisClient.Get(ctx, key).Result()
}

// DeleteCache removes a key from Redis
func DeleteCache(key string) error {
	return RedisClient.Del(ctx, key).Err()
}
