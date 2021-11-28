package database

import "github.com/go-redis/redis"

func NewRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
			Addr:     "redisdb:6379",
			Password: "",
			DB:       0,
		});
}