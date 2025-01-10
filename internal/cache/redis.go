package cache

import (
	"context"
	"github.com/redis/go-redis"
)

var ctx = context.Background()

func NewRedisClient(addr string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: addr,
	})
}
