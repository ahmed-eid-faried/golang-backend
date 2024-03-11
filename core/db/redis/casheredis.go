package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var CTX context.Context = context.Background()

// RedisClient holds the Redis client instance
var RedisClient *redis.Client

func Init() {
	// Initialize Redis client
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "",               // No password set
		DB:       0,                // Use default DB
	})
}

// InitDatabase initializes the Redis database with some sample data
func InitDatabase() error {
	// Initialize Redis client
	Init()

	// Example initialization - you can customize this according to your needs
	err := RedisClient.Set(CTX, "key1", "value1", 0).Err()
	if err != nil {
		return err
	}
	err = RedisClient.Set(CTX, "key2", "value2", 0).Err()
	if err != nil {
		return err
	}
	return nil
}

// CacheValue caches a value in Redis with a specified expiration time
func CacheValue(key, value string, expiration time.Duration) error {
	return RedisClient.Set(CTX, key, value, expiration).Err()
}

// // GetValue get a value in Redis
// func GetValue(key string, expiration time.Duration) error {
// 	return RedisClient.Get(CTX, key).Err()
// }

// // SearchValue search a value in Redis by key and query
// func SearchValue(ctx context.Context, key string, q *redis.GeoSearchQuery) error {
// 	return RedisClient.GeoSearch(CTX, key, q).Err()
// }

// RemoveCachedValue removes a cached value from Redis
func RemoveCachedValue(key string) error {
	return RedisClient.Del(CTX, key).Err()
}
