package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
)

var (
	Ctx   = context.Background()
	Redis *redis.Client
)

func InitRedis() {
	Redis = redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"),
		DB:   0,
	})
	if err := Redis.Ping(Ctx).Err(); err != nil {
		log.Fatalf("cannot connect to Redis: %v", err)
	}
}
