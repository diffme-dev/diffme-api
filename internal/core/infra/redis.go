package infra

import (
	"diffme.dev/diffme-api/config"
	Redis "github.com/go-redis/redis/v8"
)

func NewRedisClient() (*Redis.Client, error) {
	rdb := Redis.NewClient(&Redis.Options{
		Addr:     config.GetConfig().RedisUri,
		Password: "",
		DB:       5,
	})

	return rdb, nil
}
