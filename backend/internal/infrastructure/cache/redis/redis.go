package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var (
	Redis *redis.Client
)

const RedisPort = "localhost:6379"

func InitRedis(ctx context.Context) error {
	Redis = redis.NewClient(&redis.Options{
		Addr: RedisPort,
		DB:   0,
	})
	if err := Redis.Ping(ctx).Err(); err != nil {
		return fmt.Errorf("cannot connect to Redis: %w", err)
	}
	return nil
}
