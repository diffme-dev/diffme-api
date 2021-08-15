package infra

import (
	"context"
	Redis "github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func NewRedisClient() (*Redis.Client, error) {
	rdb := Redis.NewClient(&Redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return rdb, nil
}
