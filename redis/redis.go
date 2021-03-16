package redis

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

func NewRedisDB(host, port, password string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password,
		DB:       0,
	})
}
