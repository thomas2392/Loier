package redis_client

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

func CreateRedisClient() {
	var ctx = context.Background()
	rdb = redis.NewClient(&redis.Options{
		Addr:     "redis-server:6379",
		Password: "",
		DB:       0,
	})

	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println("Error when establishing redis connection:", err)
		return
	}
	fmt.Println("Conection with redis sucessfull:", pong)
}

func GetRedisClient() *redis.Client {
	return rdb
}
