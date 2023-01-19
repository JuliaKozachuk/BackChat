package redisconnect

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func Setup() {
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     "redis:6380",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := RedisClient.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := RedisClient.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)
}
