package redis

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var (
	Ctx   = context.Background()
	Redis *redis.Client
)

const RedisPort = "localhost:6379"

func InitRedis() error {
	Redis = redis.NewClient(&redis.Options{
		Addr: os.Getenv(RedisPort),
		DB:   0,
	})
	if err := Redis.Ping(Ctx).Err(); err != nil {
		log.Fatalf("cannot connect to Redis: %v", err)
	}
	return nil
}
